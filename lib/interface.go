package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var scanner = bufio.NewScanner(os.Stdin) // Will this cause problems?

func GetInput(prompt string, output *string) error {
	fmt.Print(prompt)
	scanner.Scan()

	*output = scanner.Text()

	return scanner.Err()
}

func Next() error {
	var input string

	return GetInput("Press Enter to continue...", &input)
}

type Choice struct {
	key     string
	padding string
	text    string
}

type Choices []Choice

func (c Choice) String() string {
	return fmt.Sprintf("%s%s) %s", c.padding, c.key, c.text)
}

func (c Choices) SelectionIndex(selection string) (int, error) {
	for i, choice := range c {
		if selection == choice.key {
			return i, nil
		}
	}

	return -1, errors.New("selection invalid")
}

func GetChoice(prompt string, options []string) int {
	choices := make(Choices, 0)
	maxChoiceKeyLength := 0
	for i, text := range options {
		choice := Choice{
			key:  fmt.Sprintf("%d", i),
			text: text,
		}
		choices = append(choices, choice)

		choiceKeyLength := utf8.RuneCountInString(choice.key)
		if choiceKeyLength > maxChoiceKeyLength {
			maxChoiceKeyLength = choiceKeyLength
		}
	}

	for i, choice := range choices {
		choices[i].padding = strings.Repeat(
			" ",
			maxChoiceKeyLength-utf8.RuneCountInString(choice.key),
		)
	}

	fmt.Println(prompt)
	for _, choice := range choices {
		fmt.Println(choice)
	}

	var selection string
	var index int
	err := errors.New("")
	for ; err != nil; index, err = choices.SelectionIndex(selection) {
		GetInput("Your choice:  ", &selection)
	}

	return index
}
