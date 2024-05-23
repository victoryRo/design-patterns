package command

import (
	"fmt"
	"time"
)

// Forma común de utilizar un patrón de Comando es delegar la información,
// en lugar de la ejecución, a un objeto diferente.

type Comand interface {
	Info() string
}

// ---------------------------------

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

// ---------------------------------

type HelloMessage struct{}

func (h *HelloMessage) Info() string {
	return "Hello world !"
}

// ---------------------------------

func LocalCommand2() {
	var timeCommand Comand = &TimePassed{time.Now()}
	var helloCommand Comand = &HelloMessage{}

	time.Sleep(time.Second)

	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}
