package core

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func init() {
	/* Init application id */
	createApplicationId()
	/* Init logger */
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

func initLogger() {
	Logger = log.Logger{
		Formatter: NewProcyonLoggerFormatter(),
	}
}

func GetApplicationId() uuid.UUID {
	return applicationId
}
