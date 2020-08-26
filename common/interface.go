package common

import "context"

type Factory interface {
	LoadChatBot(ctx context.Context, ID string) (bot Bot, err error)
}

type Bot interface {
	AnswerQuestionByLangCode(ctx context.Context,sessionID, question string, langCode string) (answer string)
}