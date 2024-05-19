package shapes

import "victoryRo/design-patterns/chapter5/strategy2"

type TextSquare struct {
	strategy2.DrawOutput
}

func (t *TextSquare) Draw() error {
	_, _ = t.Writer.Write([]byte("Circle"))
	return nil
}
