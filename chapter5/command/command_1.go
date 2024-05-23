package command

import (
	"fmt"
)

type Command interface {
	Execute()
}

// ---------------------------------

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")

	return &ConsoleOutput{
		message: s,
	}
}

// ---------------------------------

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, cmd := range p.queue {
			cmd.Execute()
		}

		p.queue = make([]Command, 3)
	}

}

// ---------------------------------

func LocalCommand1() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))

	queue.AddCommand(CreateCommand("Four message"))
	queue.AddCommand(CreateCommand("Five message"))

	// Este ejemplo muestra cómo utilizar un controlador de comandos ( AddCommand )
	// que ejecuta el contenido del comando.
}
