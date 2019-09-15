package video

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/golang/protobuf/ptypes"

	video "cloud.google.com/go/videointelligence/apiv1"
	videopb "google.golang.org/genproto/googleapis/cloud/videointelligence/v1"
)

func DemoCode(w io.Writer, file string) error {
	ctx := context.Background()

	// Creates a client.
	client, err := video.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil
	}

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputUri: file,
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
	})
	if err != nil {
		log.Fatalf("Failed to start annotation job: %v", err)
		return nil
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		log.Fatalf("Failed to annotate: %v", err)
		return nil
	}

	// Only one video was processed, so get the first result.
	result := resp.GetAnnotationResults()[0]

	for _, annotation := range result.SegmentLabelAnnotations {
		fmt.Fprintf(w, "Description: %s\n", annotation.Entity.Description)

		for _, category := range annotation.CategoryEntities {
			fmt.Fprintf(w, "\tCategory: %s\n", category.Description)
		}

		for _, segment := range annotation.Segments {
			start, _ := ptypes.Duration(segment.Segment.StartTimeOffset)
			end, _ := ptypes.Duration(segment.Segment.EndTimeOffset)
			fmt.Fprintf(w, "\tSegment: %s to %s\n", start, end)
			fmt.Fprintf(w, "\tConfidence: %v\n", segment.Confidence)
		}
	}
	return nil
}

func ShotChangeURI(w io.Writer, file string) error {
	ctx := context.Background()
	client, err := video.NewClient(ctx)
	if err != nil {
		return err
	}

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		Features: []videopb.Feature{
			videopb.Feature_SHOT_CHANGE_DETECTION,
		},
		InputUri: file,
	})
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Running....\n")
	resp, err := op.Wait(ctx)
	if err != nil {
		return err
	}

	// A single video was processed. Get the first result.
	result := resp.AnnotationResults[0].ShotAnnotations

	for _, shot := range result {
		start, _ := ptypes.Duration(shot.StartTimeOffset)
		end, _ := ptypes.Duration(shot.EndTimeOffset)

		fmt.Fprintf(w, "Shot: %s to %s\n", start, end)
	}

	return nil
}

func TextDetectionGCS(w io.Writer, gcsURI string) error {
	ctx := context.Background()

	// Creates a client.
	client, err := video.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputUri: gcsURI,
		Features: []videopb.Feature{
			videopb.Feature_TEXT_DETECTION,
		},
	})
	if err != nil {
		log.Fatalf("Failed to start annotation job: %v", err)
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		log.Fatalf("Failed to annotate: %v", err)
	}

	// Only one video was processed, so get the first result.
	result := resp.GetAnnotationResults()[0]

	for _, annotation := range result.TextAnnotations {
		fmt.Fprintf(w, "Text: %q\n", annotation.GetText())

		// Get the first text segment.
		segment := annotation.GetSegments()[0]
		start, _ := ptypes.Duration(segment.GetSegment().GetStartTimeOffset())
		end, _ := ptypes.Duration(segment.GetSegment().GetEndTimeOffset())
		fmt.Fprintf(w, "\tSegment: %v to %v\n", start, end)

		fmt.Fprintf(w, "\tConfidence: %f\n", segment.GetConfidence())

		// Show the result for the first frame in this segment.
		frame := segment.GetFrames()[0]
		seconds := float32(frame.GetTimeOffset().GetSeconds())
		nanos := float32(frame.GetTimeOffset().GetNanos())
		fmt.Fprintf(w, "\tTime offset of the first frame: %fs\n", seconds+nanos/1e9)

		fmt.Fprintf(w, "\tRotated bounding box vertices:\n")
		for _, vertex := range frame.GetRotatedBoundingBox().GetVertices() {
			fmt.Fprintf(w, "\t\tVertex x=%f, y=%f\n", vertex.GetX(), vertex.GetY())
		}
	}

	return nil
}
