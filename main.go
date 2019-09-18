// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"os"

	language "./modules/natural_language"
	video "./modules/video"
	vision "./modules/vision"
)

// Usage: `docker run -it golang ./app [DayXX]`
func main() {
	arg := os.Args[1]

	if arg == "Day3" {
		vision.DetectLabel(os.Stdout, "./testdata/furniture.jpg")
	} else if arg == "Day4" {
		vision.DetectText(os.Stdout, "./testdata/las-vegas.jpg")
		vision.DetectFaces(os.Stdout, "./testdata/face.jpg")
	} else if arg == "Day7" {
		video.DemoCode(os.Stdout, "gs://cloud-samples-data/video/cat.mp4")
	} else if arg == "Day8" {
		video.ShotChangeURI(os.Stdout, "gs://cloud-samples-data/video/gbikes_dinosaur.mp4")
		video.TextDetectionGCS(os.Stdout, "gs://python-docs-samples-tests/video/googlework_short.mp4")
	} else if arg == "Day11" {
		language.DemoCode(os.Stdout, "Hello World")
	}
}
