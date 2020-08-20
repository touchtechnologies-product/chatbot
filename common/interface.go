package common

import "context"

type Factory interface {
	MakeChatBot(ctx context.Context) (bot Bot, err error)
}

type Bot interface {
	AnswerQuestionByLangCode(ctx context.Context, question string, langCode string) (answer string)
}