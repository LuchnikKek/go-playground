package theory

import (
	"flag"
	"fmt"
	"os"
)

func MainStartupFlagsSet() {
	parseArgsSet()
}

// go run main.go cnv -dest ./output_dir
// go run main.go filter -gray
func parseArgsSet() {
	// название, стратегия в случае ошибки flag.Parse()
	cnvFlags := flag.NewFlagSet("cnv", flag.ExitOnError)
	filterFlags := flag.NewFlagSet("filter", flag.ExitOnError)

	// декларируем флаги набора cnvFlags
	destDir := cnvFlags.String("dest", "./output", "destination folder")

	// флаги набора filterFlags
	isGray := filterFlags.Bool("gray", false, "convert to grayscale")

	if len(os.Args) < 2 {
		fmt.Println("one of subcommand required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cnv":
		cnvFlags.Parse(os.Args[2:])
	case "filter":
		filterFlags.Parse(os.Args[2:])
	default:
		flag.PrintDefaults() // справка
		os.Exit(1)
	}

	if cnvFlags.Parsed() {
		fmt.Println("cnv flags parsed")
		fmt.Println("Destination folder:", *destDir)
	} else if filterFlags.Parsed() {
		fmt.Println("filter flags parsed")
		fmt.Println("Convert to grayscale:", *isGray)
	}
}
