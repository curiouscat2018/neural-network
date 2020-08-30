package main

import (
	"fmt"
	"neural-network/minst"
)

func main() {
	p := minst.NewTrainingDataProvider()
	for i := 0; i < 10; i++ {
		image := p.GetNextImage()
		image.PrettyPrint()
		fmt.Println()
	}
}
