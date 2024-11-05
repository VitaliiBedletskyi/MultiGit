package log

import "fmt"

// ANSI escape codes for custom badge styles with high-contrast text
const (
	reset          = "\033[0m"
	bold           = "\033[1m"
	bgBrightRed    = "\033[101m"
	bgBrightYellow = "\033[103m"
	bgBrightCyan   = "\033[106m"
	bgBrightGreen  = "\033[102m"
	fgBlack        = "\033[30m"
)

// Badge function to create center-aligned, high-contrast badges with a fixed width
func badge(label, bgColor string) string {
	labelLength := len(label)

	// Define fixed-width badges with high-contrast text color
	width := 9

	// Calculate left and right padding to center the label
	totalPadding := width - labelLength
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding

	// Construct the padded label with spaces for centering
	paddedLabel := fmt.Sprintf("%*s%s%*s", leftPadding, "", label, rightPadding, "")
	return fmt.Sprintf("%s%s%s%s%s", bgColor, fgBlack, bold, paddedLabel, reset)
}

func Error(err string) {
	fmt.Println(badge("ERROR", bgBrightRed), err)
}
