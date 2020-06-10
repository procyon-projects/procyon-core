package core

import "github.com/google/uuid"

func init() {
	/* Init application id */
	createApplicationId()
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

func GetApplicationId() uuid.UUID {
	return applicationId
}
