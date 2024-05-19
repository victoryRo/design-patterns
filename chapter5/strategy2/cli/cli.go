package cli

import (
	"flag"
	"log"
	"os"
	"victoryRo/design-patterns/chapter5/strategy2/shapes"
)

var output = flag.String("output", "text", "The output to use between 'console' and 'image' file")

func LocalMain() {
	flag.Parse()

	activeStrategy, err := shapes.Factory(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case shapes.TEXT_STRATEGY:
		activeStrategy.SetWrite(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("images/image3.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()

		activeStrategy.SetWrite(w)
	}

	err = activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
