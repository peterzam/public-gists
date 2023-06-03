package main

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	// Generate a 'dense' qrcode with the 'Low' level error correction and write it to Stdout
	var input string
	fmt.Scanf("%s", &input)
	qrterminal.GenerateWithConfig(input, qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	})
}
