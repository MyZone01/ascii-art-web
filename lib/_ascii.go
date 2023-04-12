/***********************************************************************************************/
/************************************** CODE SOURCE ********************************************/
package src

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const fontHeight = 8

func readFont(fileName string) map[string][]string {
	input, err := os.ReadFile("./font/" + fileName + ".txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	font := make(map[string][]string)
	for i, s := range strings.Split(string(input), "\n\n") {
		font[string(rune(i+32))] = strings.Split(s, "\n")
	}
	return font
}

func printline(input string, font map[string][]string) []byte {
	var output bytes.Buffer
	for _, mot := range strings.Split(input, "\n") {
		for row := 0; row < fontHeight; row++ {
			if len(string(mot)) == 0 {
				output.WriteString("\n")
				break
			}
			for _, char := range mot {
				if char < 32 || char > 126 {
					output.WriteString("We don't deal with this type of char Son of devil !!")
					break
				} else {
					output.WriteString(font[string(char)][row])
				}
			}
			output.WriteString("\n")
		}
	}
	return output.Bytes()
}

func Ascii(input string, fontName string) []byte {
	text := strings.ReplaceAll(input, "\\n", "\n")
	font := readFont(fontName)
	output := printline(text, font)
	return output
}

/***********************************************************************************************/
/***********************************************************************************************/
