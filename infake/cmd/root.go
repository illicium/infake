package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"revision.aeip.apigee.net/dia/infake"
)

var cfgFile string
var cfg infake.Config

var RootCmd = &cobra.Command{
	Use:   "infake",
	Short: "Fake data generator for InfluxDB",
	Long:  `Fake data generator for InfluxDB`,
	Run: func(cmd *cobra.Command, args []string) {
		gen := infake.NewGen(cfg)
		points, err := gen.Generate()

		if err != nil {
			log.Fatal(err)
		}

		for p := range points {
			fmt.Printf("%s\n", p)
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.infake.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".infake")
	viper.AddConfigPath("$HOME")
	viper.SetEnvPrefix("infake")

	viper.SetDefault("seed", 0)
	viper.SetDefault("time", time.Now())

	viper.BindEnv("seed")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Using config file: %q", viper.ConfigFileUsed())

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}
}
