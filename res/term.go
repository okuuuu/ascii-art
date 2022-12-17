package res

import (
	"strings"
	"syscall"
	"unsafe"
)

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

const (
	SYS_IOCTL  = 16
	TIOCGWINSZ = 0x5413
)

func ioctl(fd int, req uint, arg uintptr) (err error) {
	_, _, e1 := syscall.Syscall(SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		Check(e1)
	}
	return
}

// IoctlGetWinsize function copy from the unix package

func ioctlGetWinsize(fd int, req uint) (*Winsize, error) {
	var value Winsize
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&value)))
	return &value, err
}

func GetCols(fd uintptr) int {
	size, err := ioctlGetWinsize(int(fd), TIOCGWINSZ)
	if err != nil {
		Check(err)
	}
	return int(size.Col)
}

func Align(col int, word []string, align string) []string {
	paddedLines := make([]string, len(word))

	for i, line := range word {
		maxLength := len(DeleteColorTags(word)[i])
		var padded string
		if maxLength >= col {
			padded = line
			paddedLines[i] = padded
			continue
		}

		if maxLength == 0 {
			padded = line
			paddedLines[i] = padded
			continue
		}

		var leftPad int
		if align == "center" {
			leftPad = (col - maxLength) / 2
		} else if align == "right" {
			leftPad = col - maxLength
		} else {
			leftPad = 0
		}
		if leftPad < 0 {
			leftPad = 0
		}

		padded = strings.Repeat(" ", leftPad) + line
		paddedLines[i] = padded
	}
	return paddedLines
}
