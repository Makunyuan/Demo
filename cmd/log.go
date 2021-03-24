package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {

	writer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	sugarLogger.Debug()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./test.log")
	umberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    20,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(umberJackLogger)
}

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	for {
		fmt.Println("++")
		simpleHttpGet("www.google.com")
		simpleHttpGet("http://www.google.com")
	}

}
