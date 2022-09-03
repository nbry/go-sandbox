package tourofgo

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/png"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, 0, dy)

	for y := 0; y < dy; y++ {
		xRow := make([]uint8, 0, dx)

		for x := 0; x < dx; x++ {
			value := uint8(x * y)
			xRow = append(xRow, value)
		}

		pic = append(pic, xRow)
	}

	return pic
}

func Show() {
	pic.Show(Pic)
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}
