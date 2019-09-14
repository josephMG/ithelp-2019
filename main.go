// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"os"

	// vision "./modules/vision"
	video "./modules/video"
)

func main() {
	// vision.DetectLabel(os.Stdout, "./testdata/furniture.jpg")
	//vision.DetectText(os.Stdout, "./testdata/las-vegas.jpg")
	// vision.DetectFaces(os.Stdout, "./testdata/face.jpg")
	video.DemoCode(os.Stdout, "gs://cloud-samples-data/video/cat.mp4")
}
