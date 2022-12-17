package main

import (
	"fmt"
	"os"
	"strings"

	"main/res"
)

// All this running around
// I can't fight it much longer
// Something's tryin' to get out
// And it's never been closer

func main() {

	size := len(os.Args)
	var wordstr, banner, output_file, reverse_file, align string
	colorval := make(map[string]string)
	reverse, output, color := false, false, false

	// Read command line arguments.

	if size == 1 {
		res.HelpMsg()
	} else {
		isFlag := true
		for i, v := range os.Args[1:] {
			if strings.HasPrefix(v, "--reverse=") && isFlag {
				reverse_file = res.FlagValues(v)
				reverse = true
				color = false
			} else if strings.HasPrefix(v, "--color=") && isFlag {
				colorval[res.FlagValues(v)] = ""
				color = true
			} else if strings.HasPrefix(v, "--output=") && isFlag {
				output = true
				color = false
				output_file = res.FlagValues(v)
			} else if strings.HasPrefix(v, "--align=") && isFlag {
				align = res.FlagValues(v)
				color = false
			} else if strings.Contains(v, "--") {
				res.HelpMsg()
			} else if color && i != size-2 {
				colorval[res.FlagValues(os.Args[i])] = v
				color = false
			} else if wordstr == "" {
				wordstr = v
				isFlag = false
			} else if banner == "" && i == size-2 {
				banner = "res/" + v + ".txt"
			} else {
				res.HelpMsg()
			}
		}

		if banner == "" {
			banner = "res/standard.txt"
		}

		// Generate rune-[8]string map from ascii art file.

		mapChar := res.ReadAscii(banner)

		if reverse {

			fmt.Println(res.AsciiReverse(reverse_file, mapChar))

		} else {

			// Generate index-color map.
			mapColor := res.GetMapCol(colorval, wordstr)

			word := strings.Split(strings.ReplaceAll(strings.ReplaceAll(wordstr, "\\n", "\n"), "\\!", "!"), "\n")
			wordAscii := res.WriteAscii(word, mapChar, mapColor, align)

			// Print into different outputs.
			res.FileOrStdout(word, wordAscii, output, output_file)
		}
	}
}
