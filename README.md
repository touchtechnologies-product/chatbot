## Feature:

1. Dialogflow 
    1. Get agent by name
    2. List intents
    3. Get intent by name
    4. Create intent
    5. Detect intent from sentence

<strong>Note: Dialog API Currently does not support agent and fulfilment creation</strong>

## Example

```
_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", [path to credentials])

fac, err := GetFactory("dialogflow")
if err != nil {
    // TODO handle error
}

ctx := context.Background()
bot, err := fac.MakeChatBot(ctx)
if err != nil {
    // TODO handle error
}

ans := bot.AnswerQuestionByLangCode(ctx, "What is your name?", "en")
```