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
	index   int
}

func (c Choice) String() string {
	return fmt.Sprintf("%s%s) %s", c.padding, c.key, c.text)
}

type Choices []Choice

func MakeChoices(options []string) Choices {
	choices := make(Choices, 0)
	maxChoiceKeyLength := 0
	for i, text := range options {
		choice := Choice{
			key:   fmt.Sprintf("%d", i),
			text:  text,
			index: i,
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

	return choices
}

func (c Choices) Map() map[string]Choice {
	choiceMap := make(map[string]Choice)
	for _, choice := range c {
		choiceMap[choice.key] = choice
	}

	return choiceMap
}

func (c Choices) SelectionIndex(selection string) (int, error) {
	// TODO:  Update to use c.Map()
	for i, choice := range c {
		if selection == choice.key {
			return i, nil
		}
	}

	return -1, errors.New("selection invalid")
}

func (c Choices) SelectionIndices(selection string, maximum int) ([]int, error) {
	if selection == "" {
		return []int{}, nil
	}

	choiceMap := c.Map()

	var indices []int
	selections := strings.Split(selection, ",")

	if maximum > -1 && len(selections) > maximum {
		fmt.Printf("Please choose up to %d items.\n", maximum)
		return nil, errors.New("selection invalid")
	}

	for _, selection := range selections {
		selection = strings.TrimSpace(selection)
		if choice, ok := choiceMap[selection]; ok {
			indices = append(indices, choice.index)
		} else {
			fmt.Printf("%s is not a valid selection.\n", selection)
			return nil, errors.New("selection invalid")
		}
	}

	// Check for duplicates.
	indexMap := make(map[int]struct{})
	for _, index := range indices {
		if _, ok := indexMap[index]; ok {
			fmt.Println("Each item can only be selected up to once.")
			return nil, errors.New("selection invalid")
		} else {
			indexMap[index] = struct{}{}
		}
	}

	return indices, nil
}

func GetChoice(prompt string, options []string) int {
	choices := MakeChoices(options)

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

func GetChoices(prompt string, options []string, maximum int) []int {
	choices := MakeChoices(options)

	fmt.Println(prompt)
	for _, choice := range choices {
		fmt.Println(choice)
	}

	var selection string
	var indices []int
	err := errors.New("")
	for ; err != nil; indices, err = choices.SelectionIndices(selection, maximum) {
		GetInput("Your choices (comma separated):  ", &selection)
	}

	return indices
}
