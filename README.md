## Feature:

1. Dialogflow 
    1. Get agent by name
    2. Create intent
    3. Detect intent from sentence

<strong>Note: Dialog API Currently does not support agent and fulfilment creation</strong>

## Example

```
_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", [path to credentials])

ctx := context.Background()
fac, err := GetFactory("dialogflow")
if err != nil {
 //TO DO handle error
}
bot, err := fac.LoadChatBot(ctx, "project-name")
if err != nil {
 //TO DO handle error
}

ans := bot.AnswerQuestionByLangCode(context.Background(),"sessionID", "who are u..?", "en")
```