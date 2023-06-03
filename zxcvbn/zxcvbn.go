package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nbutton23/zxcvbn-go"
)

func main() {
	for {
		fmt.Println("Enter password or Ctrl-c to exit:")
		reader := bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')
		result := zxcvbn.PasswordStrength(password, nil)
		fmt.Println(result.CalcTime)
		fmt.Println(result.CrackTime)
		fmt.Println(result.CrackTimeDisplay)
		fmt.Println(result.Entropy)
		fmt.Println(result.MatchSequence)
		fmt.Println(result.Score)
	}
}
