// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"os"

	vision "./modules/vision"
)

func main() {
	// vision.DetectLabel(os.Stdout, "./testdata/furniture.jpg")
	//vision.DetectText(os.Stdout, "./testdata/las-vegas.jpg")
	vision.DetectFaces(os.Stdout, "./testdata/face.jpg")
}
