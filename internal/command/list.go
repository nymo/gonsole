package command

type List interface {
	Add(command Command)
	Commands() []Command
}

type list struct {
	commands []Command
}

func (list *list) Add(command Command) {
	list.commands = append(list.commands, command)
}

func (list *list) Commands() []Command {
	return list.commands
}

func NewList() List {
	return &list{}
}
