package adapter

import "fmt"

// ------------------------------------------ interface

type LegacyPrinter interface {
	Print(s string) string
}

// ------------------------------------------ first implementation

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy printer: %s", s)
	println(newMsg)
	return
}

// ------------------------------------------ adapter

type ModernPrinter interface {
	PrintStore() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStore() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
