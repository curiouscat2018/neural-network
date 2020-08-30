package minst

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Side length of image.
const N = 28

type Provider interface {
	HasNextImage() bool
	GetNextImage() *Image
}

type provider struct {
	curIndex     int
	size         int
	imagesReader *bufio.Reader
	labelsReader *bufio.Reader
}

func NewTrainingDataProvider() Provider {
	p := newProvider(
		"minst/train-images-idx3-ubyte",
		"minst/train-labels-idx1-ubyte")

	fmt.Printf("new training data provider with %v images\n", p.size)
	return p
}

func newProvider(imagesFilename, labelsFilename string) *provider {
	p := provider{}

	imagesFile, err := os.Open(imagesFilename)
	if err != nil {
		panic(err)
	}
	p.imagesReader = bufio.NewReader(imagesFile)

	labelsFile, err := os.Open(labelsFilename)
	if err != nil {
		panic(err)
	}
	p.labelsReader = bufio.NewReader(labelsFile)

	p.prepare()
	return &p
}

func (p *provider) HasNextImage() bool {
	return p.curIndex < p.size
}

func (p *provider) GetNextImage() *Image {
	if p.curIndex >= p.size {
		panic("no more pixels")
	}

	label := p.read1ByteInt(p.labelsReader)
	pixels := make([][]uint8, N)
	for i := 0; i < N; i++ {
		pixels[i] = p.readNBytesSlice(p.imagesReader)
	}

	p.curIndex++
	return &Image{
		Pixels: pixels,
		Label:  label,
	}
}

func (p *provider) prepare() {
	magic := p.read4BytesInt(p.imagesReader)
	if magic != 2051 {
		panic("invalid magic number")
	}

	size := p.read4BytesInt(p.imagesReader)
	if size <= 0 {
		panic("invalid size")
	}
	p.size = size

	if rowSize := p.read4BytesInt(p.imagesReader); rowSize != N {
		panic("invalid row size")
	}

	if colSize := p.read4BytesInt(p.imagesReader); colSize != N {
		panic("invalid col size")
	}

	magic = p.read4BytesInt(p.labelsReader)
	if magic != 2049 {
		panic("invalid magic number")
	}

	if size := p.read4BytesInt(p.labelsReader); size != p.size {
		panic("mismatched size")
	}
}

func (p *provider) read1ByteInt(reader *bufio.Reader) int {
	b, err := reader.ReadByte()
	if err != nil {
		panic(err)
	}

	return int(b)
}

func (p *provider) read4BytesInt(reader *bufio.Reader) int {
	bytes := make([]byte, 4)
	_, err := io.ReadFull(reader, bytes)
	if err != nil {
		panic(err)
	}

	res := binary.BigEndian.Uint32(bytes)
	return int(res)
}

func (p *provider) readNBytesSlice(reader *bufio.Reader) []uint8 {
	bytes := make([]byte, N)
	_, err := io.ReadFull(reader, bytes)

	if err != nil {
		panic(err)
	}

	res := make([]uint8, N)
	for i := 0; i < N; i++ {
		res[i] = uint8(bytes[i])
	}

	return res
}
