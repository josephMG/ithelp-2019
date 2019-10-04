package translation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	auth "../auth"
	"github.com/imroc/req"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func TranslateText(text []string) error {
	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the target language.
	target, err := language.Parse("zh-TW")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, text, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translations: %+v\n", translations)
	fmt.Printf("Translation[0]: %v\n", translations[0].Text)

	return nil
}

func TranslateTextV3Beta1(text []string) error {
	token, _ := auth.ServiceAccount("./authentication.json", "https://www.googleapis.com/auth/cloud-translation")

	header := req.Header{
		"Accept":        "application/json",
		"Content-Type":  "application/json; charset=utf-8",
		"Authorization": "Bearer " + token.AccessToken,
	}

	body := struct {
		SourceLanguageCode string   `json:"sourceLanguageCode"`
		TargetLanguageCode string   `json:"targetLanguageCode"`
		Contents           []string `json:"contents"`
	}{
		SourceLanguageCode: "en",
		TargetLanguageCode: "zh-TW",
		Contents:           text,
	}
	json_string, _ := json.Marshal(body)

	param := req.BodyJSON(json_string)
	// only url is required, others are optional.
	r, err := req.Post(
		fmt.Sprintf("https://translation.googleapis.com/v3beta1/projects/%s:translateText", os.Getenv("PROJECT_ID")),
		header,
		param,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", r) // print info (try it, you may surprise)

	return nil
}
