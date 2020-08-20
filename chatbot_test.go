package chatbot

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupTest() {
	_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "tag-health-project-768968c5dd0b.json")
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestAnswerQuestion() {
	fac, err := GetFactory("dialogflow")
	suite.NoError(err)
	bot, err := fac.MakeChatBot(context.Background())
	suite.NoError(err)
	time.Sleep(15*time.Second)
	ans := bot.AnswerQuestionByLangCode(context.Background(), "Your name is?", "en")

	contains := []string{
		"My name is Dialogflow!",
		"I have a name Dialogflow!",
		"Dialogflow is my name",
	}
	suite.Contains(contains, ans)
}
