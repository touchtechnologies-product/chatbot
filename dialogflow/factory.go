package dialogflow

import (
	"context"

	"github.com/touchtechnologies-product/chatbot/common"
)

type Factory struct{}

func (fac Factory) LoadChatBot(ctx context.Context, ID string) (bot common.Bot, err error) {
	dfBot := &Bot{
		project: ID,
	}

	_, err = dfBot.getAgent(ctx)

	return dfBot, err
}
