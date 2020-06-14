package core

import (
	"github.com/google/uuid"
)

func init() {
	/* Init application id */
	createApplicationId()
	/* Configure Log */
	configureLog()
	/* Type Converter Service */
	Register(NewDefaultTypeConverterService)
}

var (
	applicationId uuid.UUID
)

func GetApplicationId() uuid.UUID {
	return applicationId
}

func createApplicationId() {
	var err error
	applicationId, err = uuid.NewUUID()
	if err != nil {
		panic("Could not application id")
	}
}
