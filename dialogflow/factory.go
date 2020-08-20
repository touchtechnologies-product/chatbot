package dialogflow

import (
	"chatbot/common"
	"context"
)

type Factory struct {
}

func (fac Factory) MakeChatBot(ctx context.Context) (bot common.Bot, err error) {
	dfBot := &Bot{
		project: "tag-health-project",
		trainingPhrases: []string{
			"What is your name?",
			"Do you have a name?",
			"Tell me your name",
		},
		Responses: []string{
			"My name is Dialogflow!",
			"I have a name Dialogflow!",
			"Dialogflow is my name",
		},
	}

	intentDisplayName := "TAGHealthLabIntent"
	_, err = dfBot.getIntentByDisplayName(ctx, intentDisplayName)
	if err != nil {
		_, err = dfBot.createIntentFromTrainingPhrases(ctx, intentDisplayName)
	}

	return dfBot, err
}
