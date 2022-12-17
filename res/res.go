package res

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var out io.Writer = os.Stdout

// Arbitrary functions.

func DeleteColorTags(str []string) []string {

	otp := make([]string, len(str))

	for i := range str {
		r, _ := regexp.Compile("\x1b(.*?)m")
		otp[i] = r.ReplaceAllString(str[i], "")
	}

	return otp
}

func ContainsArr(a [95][8]string, b [8]string) int {
	for i, v := range a {
		if v == b {
			return i + 32
		}
	}
	return -1
}

func LetterInRange(lit int, mapColor map[[2]int]string) (bool, int) {
	for i, v := range mapColor {
		if lit >= i[0] && lit < i[1] {
			return true, GetColor(v)
		}
	}
	return false, 15
}

func EveryIndex(s, substr string, i []int) []int {
	if strings.Contains(s, substr) {
		if len(i) == 0 {
			i = EveryIndex(s[strings.Index(s, substr)+len(substr):], substr, append(i, strings.Index(s, substr)))
		} else {
			i = EveryIndex(s[strings.Index(s, substr)+len(substr):], substr, append(i, strings.Index(s, substr)+i[len(i)-1]+len(substr)))
		}
		return i
	} else {
		return i
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func FlagValues(arg string) string {
	values := strings.Split(arg, "=")

	return values[len(values)-1]
}

func HelpMsg() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --reverse=<fileName.txt> something standard")
	os.Exit(0)
}

// Get color for every instance of substring v at the index range (val val+len(v)). Empty color map
// is filled with white as default.

func GetMapCol(m map[string]string, w string) map[[2]int]string {

	res := make(map[[2]int]string)

	if len(m) == 0 {
		res[[2]int{0, len(w)}] = "white"
	}

	for k, v := range m {
		if v == "" {
			res[[2]int{0, len(w)}] = k
			continue
		}

		for _, val := range EveryIndex(w, v, []int{}) {
			res[[2]int{val, val + len(v)}] = k
		}
	}
	return res
}

// Open a file containing ASCII art characters. Scan every line of the file until the row with
// a corresponding character has been found. Then scan next 8 lines, add them correspondingly
// to the subarray of specific character and then close the file.
// TODO: change 2D array to an actual map.

func ReadAscii(banner string) [95][8]string {

	file, err := os.Open(banner)
	Check(err)
	defer file.Close()

	var mp [95][8]string

	sc := bufio.NewScanner(file)
	for c := 0; c < 95; c++ {
		sc.Scan()
		for i := 0; i < 8; i++ {
			sc.Scan()
			mp[c][i] = sc.Text()
		}
	}

	return mp
}

// Initialize default color as the color of wordstr substring, if such exists.
// Add to the output slice formatted string with specified color letter by letter.

func WriteAscii(input []string, mp [95][8]string, mapColor map[[2]int]string, align string) []string {

	var output []string
	lit := 0

	colm := GetCols(os.Stdout.Fd())
	spaceMap := make(map[int][]int)

	coldef := 15
	if val, b := mapColor[[2]int{0, len(strings.Join(input, "\\n"))}]; b {
		coldef = GetColor(val)
		delete(mapColor, [2]int{0, len(strings.Join(input, "\\n"))})
	}

	for _, i := range input {
		if i == "" {
			output = append(output, "")
			continue
		}
		splitString := strings.SplitAfter(i, " ")
		for v, word := range splitString {
			rword := []rune(word)
			for n, k := range rword {
				for l := 0; l < 8; l++ {
					if v == 0 && n == 0 {
						b, col := LetterInRange(lit, mapColor)
						if b {
							output = append(output, fmt.Sprintf("\x1B[38;5;%dm%s\x1B[0m", col, mp[k-32][l]))
						} else {
							output = append(output, fmt.Sprintf("\x1B[38;5;%dm%s\x1B[0m", coldef, mp[k-32][l]))
						}
					} else {
						b, col := LetterInRange(lit, mapColor)
						if b {
							output[len(output)-8+l] += fmt.Sprintf("\x1B[38;5;%dm%s\x1B[0m", col, mp[k-32][l])
						} else {
							output[len(output)-8+l] += fmt.Sprintf("\x1B[38;5;%dm%s\x1B[0m", coldef, mp[k-32][l])
						}
					}
				}
				lit++
			}
			if v != len(splitString)-1 {
				spaceMap[len(output)-8] = append(spaceMap[len(output)-8], len(output[len(output)-8]))
			}
		}
		lit += 2
	}
	if align == "justify" {
		for i := range spaceMap {
			justPad := (colm - len(DeleteColorTags(output)[i])) / len(spaceMap[i])
			if justPad < 0 {
				justPad = 0
			}
			for j, k := range spaceMap[i] {
				for l := 0; l < 8; l++ {
					output[i+l] = output[i+l][:k+j*justPad] + strings.Repeat(" ", justPad) + output[i+l][k+j*justPad:]
				}
			}
		}
	}
	output = Align(colm, output, align)
	return output
}

func PrintAscii(str2 []string, wr io.Writer) {
	output := strings.Join(str2, "\n")
	if len(strings.Join(str2, "")) != 0 {
		output += "\n"
	}
	_, err := fmt.Fprintf(wr, "%s", output)
	Check(err)
}

func FileOrStdout(str1 []string, str2 []string, b bool, f string) {
	if b {
		f, err := os.Create(f)
		Check(err)
		defer f.Close()

		w := bufio.NewWriter(f)

		str2 = DeleteColorTags(str2)

		PrintAscii(str2, w)

		w.Flush()
	} else {
		PrintAscii(str2, out)
	}
}

func AsciiReverse(f string, mp [95][8]string) string {

	var res []rune

	file, err := os.Open(f)
	Check(err)
	defer file.Close()
	var buf [8]string
	i_prev := 0

	sc := bufio.NewScanner(file)

	for c := 0; ; c++ {
		if !sc.Scan() {
			return string(res)
		}
		if c != 0 {
			res = append(res, '\\', 'n')
		}
		if sc.Text() == "" {
			continue
		}
		for i := 0; i < 8; i++ {
			buf[i] = sc.Text()
			if i != 7 {
				sc.Scan()
			}
		}
		for i := range buf[0] {
			char := [8]string{}
			if buf[0][i] == ' ' && buf[1][i] == ' ' && buf[2][i] == ' ' && buf[3][i] == ' ' && buf[4][i] == ' ' && buf[5][i] == ' ' && buf[6][i] == ' ' && buf[7][i] == ' ' {
				for j := 0; j < 8; j++ {
					char[j] = buf[j][i_prev : i+1]
				}
				rn := ContainsArr(mp, char)
				if rn > 0 {
					res = append(res, rune(rn))
					i_prev = i + 1
				}
			}
		}
		i_prev = 0
		buf = [8]string{}
	}
}
