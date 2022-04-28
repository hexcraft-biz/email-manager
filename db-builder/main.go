package main

import (
	"github.com/cure1218/dbb"
	"os"
)

func main() {
	vhs := dbb.NewVerHandlers()
	vhs.AddVerHandler("0.0", "1.0", V1_0)

	env := Load()
	dbi, err := dbb.AdminConn("mysql", env.Host, env.Port, env.AdminUser, env.AdminPassword)
	defer dbi.CloseDB()
	MustNot(err)
	MustNot(dbi.Build(env.AppDBName, env.AppDBUser, env.AppDBUserHost, env.AppDBUserPassword, vhs))
}

//================================================================
//
//================================================================
type Env struct {
	AdminUser         string
	AdminPassword     string
	Host              string
	Port              string
	AppDBName         string
	AppDBUser         string
	AppDBUserHost     string
	AppDBUserPassword string
}

func Load() *Env {
	return &Env{
		AdminUser:         os.Getenv("ADMIN_USER"),
		AdminPassword:     os.Getenv("ADMIN_PASSWORD"),
		Host:              os.Getenv("HOST"),
		Port:              os.Getenv("PORT"),
		AppDBName:         os.Getenv("APP_DB_NAME"),
		AppDBUser:         os.Getenv("APP_DB_USER"),
		AppDBUserHost:     os.Getenv("APP_DB_USER_HOST"),
		AppDBUserPassword: os.Getenv("APP_DB_USER_PASSWORD"),
	}
}

//================================================================
//
//================================================================
func MustNot(err error) {
	if err != nil {
		panic(err)
	}
}
