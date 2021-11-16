package conf

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Conf struct {
	Port     int
	Ip       string
	Location string
	Database DatabaseConf
}

type DatabaseConf struct {
	Driver string
	Host   string
	Port   int
	User   string
	Pass   string
	DB     string
	SSL    string
}

var conf Conf
var err error

func Setup() {
	log.Info().Msg("starting service configuration")
	wait := 1
	confPath := viper.GetString("conf")
	if confPath == "" {
		log.Info().Msg("conf file not specified, using default (./conf.toml or ./conf/conf.toml)")
		viper.SetConfigName("conf")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./conf")
	} else {
		viper.SetConfigFile(confPath)
	}
	viper.AutomaticEnv()
	// serveConf = *conf
	for {
		err = readAndUnmarshal()
		if err == nil {
			log.Info().Msg("conf file loaded successfuly")
			return
		}
		if wait == 32 {
			wait = 1
		} else {
			wait = wait * 2
		}
		log.Error().Err(err).Msg(fmt.Sprintf("cannot read conf file, retrying in %d seconds", wait))
		time.Sleep(time.Duration(wait) * time.Second)
	}
}

func readAndUnmarshal() error {
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return err
	}
	return nil
}

func Get() *Conf {
	return &conf
}
