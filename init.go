package core

import (
	"github.com/google/uuid"
	"strings"
)

func init() {
	/* Init application id */
	createApplicationId()
	/* Configure logger formatter */
	configureLoggerFormatter()
	/* Type Converter Service */
	Register(NewDefaultTypeConverterService)
}

var (
	applicationId uuid.UUID
)

func createApplicationId() {
	var err error
	applicationId, err = uuid.NewUUID()
	if err != nil {
		Logger.Fatal("Could not application id")
	}
}

func configureLoggerFormatter() {
	formatter := Logger.Formatter.(*ProcyonLoggerFormatter)
	strAppId := applicationId.String()
	separatorIndex := strings.Index(strAppId, "-")
	formatter.applicationId = strAppId[:separatorIndex]
}

func GetApplicationId() uuid.UUID {
	return applicationId
}
