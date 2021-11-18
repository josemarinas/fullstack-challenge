package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"server/conf"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var db *sql.DB

func Setup() {
	var err error
	confObject := conf.Get()
	wait := 1
	log.Info().Msg("starting database connection")
	var connString string
	switch confObject.Database.Driver {
	case "postgres":
		connString = fmt.Sprintf("host=%s port=%d password=%s dbname=%s user=%s sslmode=%s", confObject.Database.Host, confObject.Database.Port, confObject.Database.Pass, confObject.Database.DB, confObject.Database.User, confObject.Database.SSL)
	default:
		log.Panic().Msg("database not suported")
	}
	for {
		err = connectAndPing(confObject, connString)
		if err == nil {
			log.Info().Msg("connection to database established successfully")
			log.Info().Msg("starting database initialization")
			err = initialize()
			if err == nil {
				log.Info().Msg("database initialized successfully")
				return
			}
		}
		if wait == 32 {
			wait = 1
		} else {
			wait = wait * 2
		}
		log.Error().Err(err).Msg(fmt.Sprintf("error in databse connection or initialization, retrying in %d seconds", wait))
		time.Sleep(time.Duration(wait) * time.Second)
	}
}

func connectAndPing(confObject *conf.Conf, connString string) error {
	var err error
	db, err = sql.Open(confObject.Database.Driver, connString)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func initialize() error {
	sql, err := ioutil.ReadFile("./db/schema.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sql))
	if err != nil {
		return err
	}
	return nil
}
