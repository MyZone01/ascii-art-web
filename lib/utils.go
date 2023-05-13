package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

func GetArgs() (string, string, string, string, string, string, bool) {
	input := ""
	typeAscii := "standard"
	outputFile := ""
	isReverse := false
	align := "left"
	colorize := ""
	colorCode := ""
	for _, arg := range os.Args {
		if len(arg) > 9 && arg[:9] == "--output=" {
			outputFile = strings.Split(arg, `=`)[1]
		} else if len(arg) > 8 && arg[:8] == "--align=" {
			align = strings.Split(arg, `=`)[1]
		} else if len(arg) > 10 && arg[:10] == "--reverse=" {
			fileName := strings.Split(arg, `=`)[1]
			_text, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println("❌ ERROR: File not found")
				os.Exit(1)
			}
			input = string(_text)
			if len(input) == 0 {
				fmt.Println("❌ ERROR: Bad file format")
				os.Exit(1)
			}
			isReverse = true
		} else if len(arg) > 8 && arg[:8] == "--color=" {
			arr := strings.Split(arg, `=`)
			_colorCode := arr[1]
			if _colorCode[0] == '#' {
				colorCode = RGBToANSI(HexToRGB(_colorCode))
			} else if len(_colorCode) >= 12 && _colorCode[:4] == "rgb(" {
				colorCode = RGBToANSI(_colorCode)
			} else {
				colorCode = ansiColors[_colorCode]
			}
			if len(arr) == 3 {
				colorize = arr[2]
			}
		} else if arg == "standard" || arg == "thinkertoy" || arg == "shadow" || arg == "zigzag" {
			typeAscii = arg
		} else {
			input = arg
		}
	}
	return input, typeAscii, outputFile, align, colorCode, colorize, isReverse
}

func IsValid(text []string) bool {
	for _, word := range text {
		for _, char := range word {
			if char < 32 || char > 127 {
				return false
			}
		}
	}
	return true
}

func IsCharacterDelimiter(text [][]rune, line, col int) bool {
	return line+7 <= len(text)-1 &&
		len(text[line]) > 0 && text[line][col] == ' ' &&
		len(text[line+1]) > 0 && text[line+1][col] == ' ' &&
		len(text[line+2]) > 0 && text[line+2][col] == ' ' &&
		len(text[line+3]) > 0 && text[line+3][col] == ' ' &&
		len(text[line+4]) > 0 && text[line+4][col] == ' ' &&
		len(text[line+5]) > 0 && text[line+5][col] == ' ' &&
		len(text[line+6]) > 0 && text[line+6][col] == ' ' &&
		len(text[line+7]) > 0 && text[line+7][col] == ' '
}

func IsAsciiSpace(text [][]rune, line, col int) bool {
	return col+6 <= len(text[line])-1 && IsCharacterDelimiter(text, line, col+1) &&
		IsCharacterDelimiter(text, line, col+2) &&
		IsCharacterDelimiter(text, line, col+3) &&
		IsCharacterDelimiter(text, line, col+4) &&
		IsCharacterDelimiter(text, line, col+5) &&
		IsCharacterDelimiter(text, line, col+6)
}

func PrintAscii(output string) {
	fmt.Print(output)
}

func ParseFile(name string, isJustifying bool) (map[int][][]rune, bool) {
	asciiCharacters := map[int][][]rune{}
	_content, err := os.ReadFile(name)
	content := string(_content)
	if err != nil {
		fmt.Println("ERROR: exit when reading file")
		return asciiCharacters, true
	}
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")
	character := [][]rune{}
	actualChar := 32
	for i := 1; i < len(lines); i++ {
		if i%9 == 0 {
			asciiCharacters[actualChar] = character
			actualChar++
			character = [][]rune{}
			continue
		}
		line := lines[i]
		if actualChar != 32 && isJustifying {
			line = strings.ReplaceAll(line, " ", "R")
		}
		character = append(character, []rune(line))
	}
	return asciiCharacters, false
}

func SaveFile(fileName string, text string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("❌ ERROR: Output file creation error")
	}
	file.WriteString(text)
	file.Close()
}
