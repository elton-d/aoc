package main

import (
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Image struct {
	image [][]string
}

type ImageEnhancer struct {
	algo string
}

func (ie *ImageEnhancer) getSpacerPixel(currValue string) string {
	if currValue == "." {
		return string(ie.algo[0])
	}
	return string(ie.algo[len(ie.algo)-1])
}

func (ie *ImageEnhancer) Enhance(i *Image, spacer string) *Image {
	padded := i.Pad(spacer)
	output := make([][]string, len(padded.image))

	for y := 0; y < len(padded.image); y++ {
		for x := 0; x < len(padded.image[0]); x++ {
			pixel := string(ie.algo[padded.getPixelCode(x, y, spacer)])
			output[y] = append(output[y], pixel)
		}
	}

	return &Image{
		image: output,
	}
}

func (ie *ImageEnhancer) EnhanceN(img *Image, n int) *Image {
	res := img
	currSpacer := "."

	for i := 0; i < n; i++ {
		res = ie.Enhance(res, currSpacer)
		currSpacer = ie.getSpacerPixel(currSpacer)
	}

	return res
}

func (img *Image) getSurroundingPixels(x, y int, spacer string) []string {
	pixels := []string{}
	for i := 0; i < 9; i++ {
		pixels = append(pixels, spacer)
	}
	if x > 0 && y > 0 {
		pixels[0] = img.image[y-1][x-1]
	}
	if y > 0 {
		pixels[1] = img.image[y-1][x]
	}
	if x < len(img.image[y])-1 && y > 0 {
		pixels[2] = img.image[y-1][x+1]
	}
	if x > 0 {
		pixels[3] = img.image[y][x-1]
	}
	pixels[4] = img.image[y][x]
	if x < len(img.image[y])-1 {
		pixels[5] = img.image[y][x+1]
	}
	if x > 0 && y < len(img.image)-1 {
		pixels[6] = img.image[y+1][x-1]
	}
	if y < len(img.image)-1 {
		pixels[7] = img.image[y+1][x]
	}
	if y < len(img.image)-1 && x < len(img.image[y])-1 {
		pixels[8] = img.image[y+1][x+1]
	}
	return pixels
}

func (img *Image) getPixelCode(x, y int, spacer string) int {
	pixels := img.getSurroundingPixels(x, y, spacer)
	sb := strings.Builder{}

	for _, p := range pixels {
		if p == "." {
			sb.WriteString("0")
		} else {
			sb.WriteString("1")
		}
	}
	return util.BinToDec(sb.String())
}

func (img *Image) String() string {
	sb := strings.Builder{}
	for y := 0; y < len(img.image); y++ {
		for x := 0; x < len(img.image[y]); x++ {
			sb.WriteString(img.image[y][x])
		}
		if y < len(img.image)-1 {
			sb.WriteString("\n")
		}

	}
	return sb.String()
}

func (img *Image) GetLitPixelsCount() int {
	count := 0
	for y := 0; y < len(img.image); y++ {
		for x := 0; x < len(img.image[y]); x++ {
			if img.image[y][x] == "#" {
				count += 1
			}
		}
	}
	return count
}

func (img *Image) Pad(spacer string) *Image {
	newImg := make([][]string, len(img.image)+2)

	for i := 0; i < len(img.image[0])+2; i++ {
		newImg[0] = append(newImg[0], spacer)
	}

	for j := 0; j < len(img.image); j++ {
		newImg[j+1] = append(newImg[j+1], spacer)
		newImg[j+1] = append(newImg[j+1], img.image[j]...)
		newImg[j+1] = append(newImg[j+1], spacer)
	}

	for i := 0; i < len(img.image)+2; i++ {
		newImg[len(img.image)+1] = append(newImg[len(img.image)+1], spacer)
	}

	return &Image{
		image: newImg,
	}
}

func ImageFromStr(s string) *Image {
	i := [][]string{}
	for y, line := range strings.Split(s, "\n") {
		i = append(i, []string{})
		for _, p := range line {
			i[y] = append(i[y], string(p))
		}
	}
	return &Image{
		image: i,
	}
}

func Part1(algo, imgStr string) int {
	img := ImageFromStr(imgStr)
	en := &ImageEnhancer{
		algo: algo,
	}

	res := en.EnhanceN(img, 2)

	return res.GetLitPixelsCount()
}

func Part2(algo, imgStr string) int {
	img := ImageFromStr(imgStr)
	en := &ImageEnhancer{
		algo: algo,
	}

	res := en.EnhanceN(img, 50)

	return res.GetLitPixelsCount()
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/20/input")
	splits := strings.Split(input, "\n\n")
	algo, imgStr := splits[0], splits[1]
	fmt.Println(Part1(algo, imgStr))
	fmt.Println(Part2(algo, imgStr))
}
