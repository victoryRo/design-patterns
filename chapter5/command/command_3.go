package command

import (
	"fmt"
	"time"
)

type Comando interface {
	Info() string
}

type ChainLog interface {
	Next(Comando)
}

// ---------------------------------

type TimePass struct {
	start time.Time
}

func (t *TimePass) Info() string {
	return time.Since(t.start).String()
}

// ---------------------------------

type Logger struct {
	NextChain ChainLog
}

func (l *Logger) Next(c Comando) {
	time.Sleep(time.Second)

	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if l.NextChain != nil {
		l.NextChain.Next(c)
	}
}

// ---------------------------------

func LocalCommand3() {
	second := new(Logger)
	first := &Logger{NextChain: second}

	command := &TimePass{start: time.Now()}

	first.Next(command)
}
