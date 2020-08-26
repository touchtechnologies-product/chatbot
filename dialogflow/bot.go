package dialogflow

import (
	"context"
	"fmt"

	sdk "cloud.google.com/go/dialogflow/apiv2"
	pb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

type Bot struct {
	project         string
	trainingPhrases []string
	Responses       []string
}

func (bot *Bot) AnswerQuestionByLangCode(ctx context.Context, sessionID, question string, langCode string) (answer string) {
	sessionClient, err := sdk.NewSessionsClient(ctx)
	if err != nil {
		return err.Error()
	}
	defer func() { _ = sessionClient.Close() }()

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", bot.project, sessionID)
	textInput := pb.TextInput{Text: question, LanguageCode: langCode}
	queryTextInput := pb.QueryInput_Text{Text: &textInput}
	queryInput := pb.QueryInput{Input: &queryTextInput}
	request := pb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return err.Error()
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText
}

func (bot *Bot) getAgent(ctx context.Context) (agent *pb.Agent, err error) {
	cli, err := sdk.NewAgentsClient(ctx)
	if err != nil {
		return nil, err
	}
	return cli.GetAgent(ctx, &pb.GetAgentRequest{Parent: fmt.Sprintf("projects/%s", bot.project)})
}