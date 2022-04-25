package ascart

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"

	_ "image/jpeg"
	_ "image/png"
)

const (
	OptionAsciis = " .:-=+*#%@"
)

type size struct {
	width, height int
}

func AtoiDefault(s string, d int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return d
	}
	return i
}

func getTerminalSize() (size, error) {
	var sz size
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return sz, err
	}
	s := strings.Fields(string(out))
	if len(s) != 2 {
		return sz, fmt.Errorf("invalid output: %s", string(out))
	}
	return size{AtoiDefault(s[1], 40), AtoiDefault(s[0], 41) - 1}, nil
}

func readImage(p string) (image.Image, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	im, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return im, nil
}

func toGrayScale(im image.Image) *image.RGBA {
	bounds := im.Bounds()
	gray := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			pixel := color.GrayModel.Convert(im.At(x, y))
			gray.Set(x, y, pixel)
		}
	}
	return gray
}

func getHeightByWidth(imw, imh, w int) int {
	return int(float64(imh) / float64(imw) * float64(w) * ratio)
}

func getWidthByHeight(imw, imh, h int) int {
	return int(float64(h) * float64(imw) / float64(imh) / ratio)
}

func grayToAscii(gray *image.RGBA) ([]string, error) {
	termSize, err := getTerminalSize()
	if err != nil {
		return nil, err
	}
	imSize := size{gray.Bounds().Max.X, gray.Bounds().Max.Y}

	if ascWidth == 0 && ascHeight == 0 {
		ascWidth = imSize.width
		ascHeight = int(float64(imSize.height) * ratio)
	} else if ascWidth == 0 {
		ascWidth = getWidthByHeight(imSize.width, imSize.height, ascHeight)
	} else if ascHeight == 0 {
		ascHeight = getHeightByWidth(imSize.width, imSize.height, ascWidth)
	}
	ascSize := size{ascWidth, ascHeight}

	if ascSize.width > termSize.width || ascSize.height > termSize.height {
		ascSize.width = termSize.width
		ascSize.height = getHeightByWidth(imSize.width, imSize.height, ascSize.width)
		if ascSize.height > termSize.height {
			ascSize.height = termSize.height
			ascSize.width = getWidthByHeight(imSize.width, imSize.height, ascSize.height)
		}
	}

	xStep := float64(imSize.width) / float64(ascSize.width)
	yStep := float64(imSize.height) / float64(ascSize.height)

	endX := 0
	endY := 0
	var out []string
	for j := 0; j < ascSize.height; j++ {
		startY := endY
		endY = int(math.Round(float64(j+1) * yStep))
		var sb strings.Builder
		for i := 0; i < ascSize.width; i++ {
			startX := endX
			endX = int(math.Round(float64(i+1) * xStep))
			c := squeeze(gray, startX, endX, startY, endY)
			sb.WriteByte(c)
		}
		endX = 0
		out = append(out, sb.String())
	}
	return out, nil
}

func squeeze(im *image.RGBA, startX, endX, startY, endY int) uint8 {
	count := float64((endX - startX) * (endY - startY))
	var totalGray int
	for j := startY; j < endY; j++ {
		for i := startX; i < endX; i++ {
			rgba := im.RGBAAt(i, j)
			totalGray += int(rgba.R)
		}
	}
	avg := int(math.Round(float64((len(OptionAsciis)-1)*totalGray) / count / 256))
	return OptionAsciis[avg]
}

func convert(im image.Image) ([]string, error) {
	gray := toGrayScale(im)
	return grayToAscii(gray)
}

func draw(asc []string) {
	for _, line := range asc {
		fmt.Println(line)
	}
}
