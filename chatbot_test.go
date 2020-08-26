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
	ctx := context.Background()
	fac, err := GetFactory("dialogflow")
	suite.NoError(err)
	bot, err := fac.LoadChatBot(ctx, "tag-health-project")
	suite.NoError(err)
	time.Sleep(15 * time.Second)
	ans := bot.AnswerQuestionByLangCode(context.Background(),"6ba7b810-9dad-11d1-80b4-00c04fd430c8", "age", "en")

	contains := []string{
		"what is your age?",
		"how ur age",
	}
	suite.Contains(contains, ans)
}
