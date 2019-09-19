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

func AnalyzeSentiment(gcsURI string) error {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}
	op, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_GcsContentUri{
				GcsContentUri: gcsURI,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Document Sentiment:\n")
	fmt.Printf("\tScore: %.2f\n", op.DocumentSentiment.Score)
	fmt.Printf("\tMagnitude: %.2f\n", op.DocumentSentiment.Magnitude)
	fmt.Printf("Language: %q\n\n", op.Language)
	for _, annotation := range op.Sentences {
		text := annotation.GetText()
		fmt.Printf("Text:\n")
		fmt.Printf("\tContent: %s\n", text.Content)
		fmt.Printf("\tSentiment: %q\n", annotation.Sentiment)
	}

	return nil
}

func AnalyzeSyntax(gcsURI string) error {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}

	op, err := client.AnnotateText(ctx, &languagepb.AnnotateTextRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_GcsContentUri{
				GcsContentUri: gcsURI,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		Features: &languagepb.AnnotateTextRequest_Features{
			ExtractSyntax: true,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", op)
	return nil
}
