package main

import (
	"flag"
	"io"
	"log"
	"os"
)

type ProgramFlags struct {
	help    *bool
	verbose *bool
	version *bool
	debug   *bool
	print   *bool
	arg     *string
	brine   *string
}

var flags = &ProgramFlags{}

func programConfigure() {
	flags.help = flag.Bool("h", false, "help")
	flags.verbose = flag.Bool("verbose", false, "verbose")
	flags.version = flag.Bool("v", false, "verbose")
	flags.debug = flag.Bool("d", false, "debug")
	flags.print = flag.Bool("p", false, "print")
	flags.arg = flag.String("a", "", "argument")
	flags.brine = flag.String("b", "", "brine")
	flag.Parse()

	if *flags.help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if !*flags.verbose {
		log.SetOutput(io.Discard)
	} else {
		log.Print("running program in verbose mode")
	}

}
