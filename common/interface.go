package common

import (
	"context"

	pb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

type Factory interface {
	LoadChatBot(ctx context.Context, ID string) (bot Bot, err error)
}

type Bot interface {
	AnswerQuestionByLangCode(ctx context.Context, sessionID, question string, langCode string) (fulfillmentMessages []*pb.Intent_Message, err error)
}