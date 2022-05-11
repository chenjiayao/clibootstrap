package cmd

import (
	"clibootstrap/lib/logger"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfgFile string

var test string

var rootCmd = &cobra.Command{
	Use:   "scr95",
	Short: "cli 项目快速开发手脚架",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		Start()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file ")
	rootCmd.PersistentFlags().StringVar(&test, "test", "", "test")

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.ReadInConfig()
	} else {
		log.Fatal("no config file")
	}
	logger.InitLogger()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Start() {
	logger.Debug("test", zap.String("typ", test))
}
