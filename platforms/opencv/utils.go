package opencv

import (
	"image"
	"image/color"

	"github.com/hybridgroup/gocv"
)

var classifier *gocv.CascadeClassifier

// loadHaarClassifierCascade returns open cv HaarCascade loaded
func loadCascadeClassifier(haar string) *gocv.CascadeClassifier {
	if classifier != nil {
		return classifier
	}

	c := gocv.NewCascadeClassifier()
	c.Load(haar)
	classifier = &c
	return classifier
}

// DetectFaces loads Haar cascade to detect face objects in image
func DetectFaces(haar string, img gocv.Mat) []image.Rectangle {
	return loadCascadeClassifier(haar).DetectMultiScale(img)
}

// DrawRectangles uses Rect array values to return image with rectangles drawn.
func DrawRectangles(img gocv.Mat, rects []image.Rectangle, r int, g int, b int, thickness int) {
	for _, rect := range rects {
		gocv.Rectangle(img, rect, color.RGBA{uint8(r), uint8(g), uint8(b), 0}, thickness)
	}
	return
}
