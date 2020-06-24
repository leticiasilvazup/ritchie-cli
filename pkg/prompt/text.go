package prompt

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/manifoldco/promptui"
)


type InputText interface {
	Text(name string, required bool) (string, error)
}

type inputText struct{}

type surveyText struct{}

func NewInputText() inputText {
	return inputText{}
}

// Text show a prompt and parse to string.
func (inputText) Text(name string, required bool) (string, error) {
	var prompt promptui.Prompt

	if required {
		prompt = promptui.Prompt{
			Label:     name,
			Pointer: promptui.PipeCursor,
			Validate:  validateEmptyInput,
			Templates: defaultTemplate(),
		}
	} else {
		prompt = promptui.Prompt{
			Label:     name,
			Pointer: promptui.PipeCursor,
			Templates: defaultTemplate(),
		}
	}

	return prompt.Run()
}

func NewSurveyText() surveyText {
	return surveyText{}
}

func (surveyText) Text(name string, required bool) (string, error) {

	var value string

	var validationQs []*survey.Question


	if required{
		validationQs = []*survey.Question{
			{
				Name:     "name",
				Prompt:   &survey.Input{Message: name},
				Validate: survey.Required,
			},
		}
	}else {
		validationQs = []*survey.Question{
			{
				Name:   "name",
				Prompt: &survey.Input{Message: name},
			},
		}
	}
	return value, survey.Ask(validationQs, &name)
}

