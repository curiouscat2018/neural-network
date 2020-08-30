package minst

import (
	"fmt"

	"github.com/gookit/color"
)

type Image struct {
	Pixels [][]uint8
	Label  int
}

func (im *Image) PrettyPrint() {
	for i := 0; i < len(im.Pixels); i++ {
		for j := 0; j < len(im.Pixels[0]); j++ {
			g := uint8(im.Pixels[i][j])
			color.RGB(g, g, g).Print("█")
		}
		fmt.Println()
	}
	fmt.Printf("label: %v\n", im.Label)
}

func (im *Image) ToSlice() []uint8 {
	res := make([]uint8, 0)
	for _, row := range im.Pixels {
		res = append(res, row...)
	}

	return res
}
