package cmd

import (
	"github.com/spf13/cobra"

	"server/api"
	"server/conf"
	"server/db"

	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "REST API server",
	Long:  `This is the root command of the REST API server, it will serve an REST APi in the provided port and host`,
	Run: func(cmd *cobra.Command, args []string) {
		conf.Setup()
		db.Setup()
		api.Start()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("conf", "c", "", "setup the config file")
	viper.BindPFlag("conf", rootCmd.PersistentFlags().Lookup("conf"))
	rootCmd.Flags().IntP("port", "p", 1337, "application serve port")
	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	rootCmd.Flags().StringP("ip", "i", "0.0.0.0", "ip of the server that will host the app")
	viper.BindPFlag("ip", rootCmd.Flags().Lookup("ip"))
	rootCmd.Flags().StringP("location", "l", "/api", "Path where the api will be listening")
	viper.BindPFlag("location", rootCmd.Flags().Lookup("location"))
	// DEFAULT DB CONF
	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "user")
	viper.SetDefault("database.pass", "pass")
	viper.SetDefault("database.db", "database")
	viper.SetDefault("database.ssl", "disable")
}
