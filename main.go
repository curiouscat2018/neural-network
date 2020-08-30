package main

import (
	"fmt"
	"neural-network/minst"
)

func main() {
	p := minst.NewTrainingDataProvider()
	image := p.GetNextImage()
	fmt.Println(image.Label)
}
