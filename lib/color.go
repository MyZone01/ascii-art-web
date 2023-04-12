package ascii_art

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ansiColors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

func Colorize(colorize string, _char string, char string, color string) (string, bool) {
	if strings.Contains(colorize, _char) && char != " " && char != "R" {
		char = fmt.Sprintf("%s%s%s", color, string(char), "\033[0m")
		return char, true
	}
	return char, false
}

func HexToRGB(hex string) string {
	// Remove the '#' character if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Convert the hex string to an integer
	hexInt, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		fmt.Println("❌ Error: Invalid color")
	}

	// Extract the RGB components by bitwise shifting and masking
	red := (hexInt >> 16) & 0xff
	green := (hexInt >> 8) & 0xff
	blue := hexInt & 0xff

	return fmt.Sprintf("rgb(%d, %d, %d)", red, green, blue)
}

func RGBToANSI(input string) string {
	// Remove "rgb(" and ")"
	input = strings.TrimPrefix(input, "rgb(")
	input = strings.TrimSuffix(input, ")")

	// Split into components
	components := strings.Split(input, ",")
	if len(components) != 3 {
		fmt.Println("❌ ERROR: Invalid format")
		os.Exit(1)
	}

	// Convert to integers
	red, err := strconv.Atoi(strings.TrimSpace(components[0]))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	green, err := strconv.Atoi(strings.TrimSpace(components[1]))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	blue, err := strconv.Atoi(strings.TrimSpace(components[2]))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)
}
