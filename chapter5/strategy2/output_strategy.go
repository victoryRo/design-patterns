package strategy2

import "io"

type Output interface {
	Draw() error
	SetLog(io.Writer)
	SetWrite(io.Writer)
}

type DrawOutput struct {
	LogWriter io.Writer
	Writer    io.Writer
}

func (d *DrawOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

func (d *DrawOutput) SetWrite(w io.Writer) {
	d.Writer = w
}
