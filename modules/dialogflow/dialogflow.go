package dialogflow

import (
	"context"
	"errors"
	"fmt"
	"os"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func DetectIntentText(sessionID, text, languageCode string) (string, error) {
	projectID := os.Getenv("PROJECT_ID")
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return "", err
	}
	defer sessionClient.Close()

	if projectID == "" || sessionID == "" {
		return "", errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", projectID, sessionID))
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	fmt.Printf("fulfillmentText: \"%v\"\n", fulfillmentText)
	return fulfillmentText, nil
}
