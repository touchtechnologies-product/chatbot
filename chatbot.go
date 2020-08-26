package chatbot

import (
	"errors"

	"github.com/touchtechnologies-product/chatbot/common"
	"github.com/touchtechnologies-product/chatbot/dialogflow"
)

func GetFactory(botType string) (fac common.Factory, err error) {
	switch botType {
	case "dialogflow":
		return dialogflow.Factory{}, nil
	default:
		return nil, errors.New("invalid type")
	}
}
