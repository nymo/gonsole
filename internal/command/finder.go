package command

import "errors"

type Finder interface {
	FindCommandToExecute(argument string) (Command, error)
}

type finder struct {
	commandList List
}

func (finder *finder) FindCommandToExecute(argument string) (Command, error) {
	//eventuell nutzen wir hier eine tokenizer oder sowas damit der finder nicht den split machen muss

	commands := finder.commandList.Commands()
	for _, com := range commands {
		return com, nil
	}

	return nil, errors.New("command not found")
}

func NewFinder(commandList List) Finder {
	return &finder{commandList}
}
