package natural_language

import (
	"context"
	"fmt"
	"io"
	"log"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func DemoCode(w io.Writer, text string) error {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
		return err
	}

	fmt.Fprintf(w, "Text: %v\n", text)
	if sentiment.DocumentSentiment.Score >= 0 {
		fmt.Fprintln(w, "Sentiment: positive")
	} else {
		fmt.Fprintln(w, "Sentiment: negative")
	}
	return nil
}
