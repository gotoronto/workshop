package display

import (
	"fmt"

	"github.com/fatih/color"
)

var Colors = []color.Attribute{
	color.BgBlack, color.BgRed, color.BgGreen, color.BgYellow,
	color.BgBlue, color.BgMagenta, color.BgCyan, color.BgWhite,
	color.BgHiBlack, color.BgHiRed, color.BgHiGreen, color.BgHiYellow,
	color.BgHiBlue, color.BgHiMagenta, color.BgHiCyan, color.BgHiWhite,
}

func Show(pic [][]int) {
	for _, row := range pic {
		for _, y := range row {
			currentColor := color.New(Colors[y%len(Colors)])
			currentColor.Print(" ")
		}
		fmt.Println(" ")
	}
}
