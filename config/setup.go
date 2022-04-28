package config

import (
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
)

//================================================================
//
//================================================================
type Config struct {
	*Env
	DB *sqlx.DB
}

func Load() (*Config, error) {
	env, err := GetEnv()
	if err != nil {
		return nil, err
	}

	return &Config{Env: env}, nil
}

func (conf *Config) GetDB() (*sqlx.DB, error) {
	var err error

	if conf.DB == nil {
		if conf.DB, err = conf.Open(); err != nil {
			return nil, err
		}
	}

	return conf.DB, nil
}

func (conf *Config) CloseDB() {
	if conf.DB != nil {
		conf.DB.Close()
	}
}

//================================================================
//
//================================================================
type Env struct {
	*DBI
	GinMode          string
	AppPort          string
	SmtpHost         string
	SmtpPort         string
	XgcAuthorization string
}

func GetEnv() (*Env, error) {
	var err error

	dbi := &DBI{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	if dbi.MaxOpen, err = strconv.Atoi(os.Getenv("DB_MAX_OPEN")); err != nil {
		return nil, err
	}

	if dbi.MaxIdle, err = strconv.Atoi(os.Getenv("DB_MAX_IDLE")); err != nil {
		return nil, err
	}

	if dbi.LifeTime, err = strconv.Atoi(os.Getenv("DB_LIFE_TIME")); err != nil {
		return nil, err
	}

	if dbi.IdleTime, err = strconv.Atoi(os.Getenv("DB_IDLE_TIME")); err != nil {
		return nil, err
	}

	return &Env{
		DBI:              dbi,
		GinMode:          os.Getenv("GIN_MODE"),
		AppPort:          os.Getenv("APP_PORT"),
		SmtpHost:         os.Getenv("SMTP_HOST"),
		SmtpPort:         os.Getenv("SMTP_PORT"),
		XgcAuthorization: os.Getenv("XGC_AUTHORIZATION"),
	}, nil
}
