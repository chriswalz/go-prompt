package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "🚀"},
		{Text: "articles", Description: "🙏"},
		{Text: "long sentence hither", Description: "🙏"},
		{Text: "articles", Description: "🤷‍♂️"},
		{Text: "articles", Description: "🙏"},
		{Text: "articles", Description: "🙏"},
		{Text: "articles", Description: "🙏"},
	}
	return prompt.FilterContains(s, in.GetWordBeforeCursor(), true)
}

func main() {
	in, p := prompt.Input(">>> ", completer,
		prompt.OptionTitle("sql-prompt"),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray))
	p.
	fmt.Println("Your input: " + in)
}
