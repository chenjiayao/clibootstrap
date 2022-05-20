package cmd

import (
	"clibootstrap/globals"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cfgFile string

var test string

var rootCmd = &cobra.Command{
	Use:   "clibootstrap",
	Short: "cli 项目快速开发手脚架",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		Start()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config/development.ini", "config file")
	rootCmd.PersistentFlags().StringVar(&test, "test", "", "test")
	cobra.OnInitialize(bootstrap)
}

func bootstrap() {
	initConfig()
	initLogger()
}

func initLogger() {
	atomicLevel := zap.NewAtomicLevel()
	level := viper.GetString("log.level")
	switch level {
	case "DEBUG":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "INFO":
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case "WARN":
		atomicLevel.SetLevel(zapcore.WarnLevel)
	case "ERROR":
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	case "PANIC":
		atomicLevel.SetLevel(zapcore.PanicLevel)
	case "FATAL":
		atomicLevel.SetLevel(zapcore.FatalLevel)
	default:
		log.Fatalf("error: log level %s is not supported", level)
	}
	zap.NewProduction()
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), //日志的编码方式，
		zapcore.AddSync(os.Stdout),            //日志的输出位置
		atomicLevel,                           // 日志等级
	)
	globals.Logger = zap.New(zapCore, zap.AddCaller(), zap.Fields(zap.String("appname", viper.GetString("log.appname"))))
	globals.SugaredLogger = globals.Logger.Sugar()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.ReadInConfig()
	} else {
		log.Fatal("no config file")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Start() {
	globals.Logger.Debug("test", zap.String("typ", test))
}
