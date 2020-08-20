package chatbot

import (
	"chatbot/common"
	"chatbot/dialogflow"
	"errors"
)

func GetFactory(botType string) (fac common.Factory, err error) {
	switch botType {
	case "dialogflow":
		return dialogflow.Factory{}, nil
	default:
		return nil, errors.New("invalid type")
	}
}