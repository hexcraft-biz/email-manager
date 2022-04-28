package model

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/hexcraft-biz/email-manager/aes"
	"github.com/hexcraft-biz/email-manager/config"
	"github.com/jmoiron/sqlx"
	"net/smtp"
	"strings"
	"time"
)

//================================================================
// https://support.google.com/a/answer/166852
//================================================================

var (
	ErrEmailExists            = errors.New("Email exists.")
	ErrEmailPasswordLen       = errors.New("Password too long.")
	ErrEmailReachedQuotaLimit = errors.New("Service reached the daily quota limit.")
)

const (
	DefGmailDailyLimit           int = 1950 // Gmail says 2000
	DefGmailReceiverPerMailLimit int = 5
	DefLenPassword               int = 16
)

//================================================================
//
//================================================================
type EEmail struct {
	*Engine
}

func NewEEmail(db *sqlx.DB) *EEmail {
	return &EEmail{
		Engine: NewEngine(db),
	}
}

func (e *EEmail) Insert(addr, password string) (*Email, error) {
	email, err := NewEmail(addr, password)
	if err != nil {
		return nil, err
	}

	if _, err := e.NamedExec(`INSERT INTO email (id, ctime, mtime, enabled, addr, password, salt) VALUES (:id, :ctime, :mtime, :enabled, :addr, :password, :salt);`, email); err != nil {
		if myErr, ok := err.(*mysql.MySQLError); !ok {
		} else if myErr.Number == 1062 {
			return nil, ErrEmailExists
		} else {
			return nil, err
		}
	}

	return email, nil
}

func (e *EEmail) GetEmailWithLowestCount() (*Email, error) {
	email := new(Email)
	return email, e.Get(email, `SELECT * FROM email WHERE daily_count = (SELECT MIN(daily_count) FROM email WHERE daily_count < ?) LIMIT 1;`, DefGmailDailyLimit)
}

func (e *EEmail) HitDailyCount(emailIDBin []byte) error {
	_, err := e.Exec(`UPDATE email SET daily_count = daily_count + 1 WHERE id = ?;`, emailIDBin)
	return err
}

func (e *EEmail) ResetDailyCount() error {
	_, err := e.Exec(`UPDATE email SET daily_count = 0;`)
	return err
}

func (e *EEmail) GetAllEnabled() ([]*Email, error) {
	ems := []*Email{}
	if err := e.Select(&ems, `SELECT * FROM email WHERE enabled = 1;`); err != nil {
		return nil, err
	} else {
		for _, em := range ems {
			if id, err := uuid.FromBytes(em.IDBin); err != nil {
				return nil, err
			} else {
				em.EmailID = id.String()
			}
		}
	}

	return ems, nil
}

//================================================================
//
//================================================================
type Email struct {
	EmailID    string `db:"-" json:"emailID"`
	ID         string `db:"-" json:"-"`
	IDBin      []byte `db:"id" json:"-"`
	Ctime      string `db:"ctime" json:"-"`
	Mtime      string `db:"mtime" json:"-"`
	Enabled    bool   `db:"enabled" json:"enabled"`
	Addr       string `db:"addr" json:"addr"`
	Password   []byte `db:"password" json:"-"`
	Salt       []byte `db:"salt" json:"-"`
	DailyCount int    `db:"daily_count" json:"dailyCount"`
}

func NewEmail(addr, password string) (*Email, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	idBin, _ := id.MarshalBinary()

	ts := time.Now().UTC().Format("2006-01-02 15:04:05")
	if salt, err := aes.GenSalt(); err != nil {
		return nil, err
	} else if encPassword, err := aes.Encrypt(salt, password); err != nil {
		return nil, err
	} else {
		return &Email{
			ID:       id.String(),
			IDBin:    idBin,
			Ctime:    ts,
			Mtime:    ts,
			Enabled:  true,
			Addr:     addr,
			Password: encPassword,
			Salt:     salt,
		}, nil
	}
}

func (e *Email) Send(conf *config.Config, to []string, subject, body string) error {
	server := conf.SmtpHost + ":" + conf.SmtpPort
	msg := []byte("From: " + e.Addr + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		body)

	if password, err := aes.Decrypt(e.Salt, e.Password); err != nil {
		return err
	} else {
		return smtp.SendMail(server, smtp.PlainAuth("", e.Addr, password, conf.SmtpHost), e.Addr, to, msg)
	}
}
