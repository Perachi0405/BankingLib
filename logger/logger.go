package logger

import "go.uber.org/zap"

var Log *zap.Logger

func init() { //used for initialization the logger
	var err error
	//Define the default type.
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	config.EncoderConfig = encoderConfig

	//logger with own configuration
	Log, err = config.Build()

	//logger with normal configuration
	// Log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

//commit check
//helper function to customise
func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...) // passing the message and fields inside the logger package.
}
