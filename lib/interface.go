package lib

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin) // Will this cause problems?

func GetInput(prompt string, output *string) error {
	fmt.Print(prompt)
	scanner.Scan()

	*output = scanner.Text()

	return scanner.Err()
}

func Next() error {
	return GetInput("Press Enter to continue...", nil)
}
