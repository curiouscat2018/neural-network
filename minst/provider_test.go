package minst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleImage() {
	p := newProvider(
		"train-images-idx3-ubyte",
		"train-labels-idx1-ubyte")

	fmt.Println("start")
	for i := 0; i < 10; i++ {
		image := p.GetNextImage()
		image.PrettyPrint()
	}
}

func Test_provider(t *testing.T) {
	p := newProvider(
		"train-images-idx3-ubyte",
		"train-labels-idx1-ubyte")

	count := 0
	for p.HasNextImage() {
		_ = p.GetNextImage()
		count++
	}

	assert.Equal(t, 60000, count)
}
