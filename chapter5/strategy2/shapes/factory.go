package shapes

import (
	"fmt"
	"os"
	"victoryRo/design-patterns/chapter5/strategy2"
)

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func Factory(s string) (strategy2.Output, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			DrawOutput: strategy2.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			DrawOutput: strategy2.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
