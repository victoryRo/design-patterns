package bridge

import (
	"fmt"
	"io"
)

// --------------------------------------------- interface

type PrinterAPI interface {
	PrintMessage(string) error
}

// --------------------------------------------- errors for pkg

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	errPrint2 = constError("You need to pass an io.Writer to PrintImple2")
)

// --------------------------------------------- two implemented

type PrintImple1 struct{}

func (p *PrintImple1) PrintMessage(msg string) error {
	fmt.Printf("%s", msg)
	return nil
}

type PrintImple2 struct {
	Writer io.Writer
}

func (d *PrintImple2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errPrint2
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

// --------------------------------------------- abstraction interface

type PrintAbstraction interface {
	Print() error
}

// --------------------------------------------- implement abstraction

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (n *NormalPrinter) Print() error {
	_ = n.Printer.PrintMessage(n.Msg)
	return nil
}

// --------------------------------------------- implement abstraction

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *PacktPrinter) Print() error {
	_ = p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Msg))
	return nil
}
