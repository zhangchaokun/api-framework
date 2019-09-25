package logging

import (
	"api-framework/core/config"
	"fmt"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", config.AppConfig.RuntimeRootPath, config.AppConfig.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		config.AppConfig.LogSaveName,
		time.Now().Format(config.AppConfig.TimeFormat),
		config.AppConfig.LogFileExt,
	)
}
