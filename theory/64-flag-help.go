package theory

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

func MainStartupFlagsHelp() {
	imgFile := flag.String("file", "", "input image file")

	flag.Usage = func() {
		fmt.Println("Package version:", version)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	fmt.Println("Image file:", *imgFile)
}
