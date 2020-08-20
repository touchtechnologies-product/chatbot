package dialogflow

import (
	sdk "cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

type Bot struct {
	project          string
	trainingPhrases []string
	Responses       []string
}

func (bot *Bot) AnswerQuestionByLangCode(ctx context.Context, question string, langCode string) (answer string) {
	sessionClient, err := sdk.NewSessionsClient(ctx)
	if err != nil {
		return err.Error()
	}
	defer func() {_ = sessionClient.Close()}()

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", bot.project, "123456789")
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

func (bot *Bot) createIntentFromTrainingPhrases(ctx context.Context, displayName string) (intent *pb.Intent, err error) {
	cli, err := sdk.NewIntentsClient(ctx)
	if err != nil {
		return nil, err
	}

	intent = &pb.Intent{
		DisplayName: displayName,
		TrainingPhrases: bot.makeTrainingPhrasesFromStrings(),
		Messages: bot.makeResponsesFromStrings(),
	}
	req := &pb.CreateIntentRequest{
		Parent: fmt.Sprintf("projects/%s/agent", bot.project),
		Intent: intent,
	}

	return cli.CreateIntent(ctx, req)
}

func (bot *Bot) makeTrainingPhrasesFromStrings() (phrases []*pb.Intent_TrainingPhrase) {
	phrases = make([]*pb.Intent_TrainingPhrase, len(bot.trainingPhrases))
	for i, v := range bot.trainingPhrases {
		phrases[i] = &pb.Intent_TrainingPhrase{
			Type:  pb.Intent_TrainingPhrase_EXAMPLE,
			Parts: []*pb.Intent_TrainingPhrase_Part{{Text: v}},
		}
	}
	return phrases
}

func (bot *Bot) makeResponsesFromStrings() (messages []*pb.Intent_Message) {
	return []*pb.Intent_Message{
		{
			Message: &pb.Intent_Message_Text_{
				Text: &pb.Intent_Message_Text{Text: bot.Responses},
			},
		},
	}
}

func (bot *Bot) getIntentByDisplayName(ctx context.Context, displayName string) (intent *pb.Intent, err error) {
	cli, err := sdk.NewIntentsClient(ctx)
	if err != nil {
		return nil, err
	}

	intent = &pb.Intent{
		DisplayName: displayName,
		TrainingPhrases: bot.makeTrainingPhrasesFromStrings(),
		Messages: bot.makeResponsesFromStrings(),
	}
	req := &pb.ListIntentsRequest{
		Parent: fmt.Sprintf("projects/%s/agent", bot.project),
	}

	cur := cli.ListIntents(ctx, req)
	for  {
		resp, err := cur.Next()
		if err == iterator.Done {
			return nil, errors.New("not found")
		}

		if err != nil {
			return nil, err
		}

		if resp.DisplayName == displayName {
			return resp, nil
		}
	}
}