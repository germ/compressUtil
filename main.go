package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/germ/imageCompress"
)

var (
	flagInput		string
	flagOutput		string
	flagExtracting	bool
	flagCompressing	bool
)

func init() {
	flag.StringVar(&flagInput, "i",	"",		"Path to compressed data.")
	flag.StringVar(&flagOutput, "o",	"",		"Path to uncompressed data.")
	flag.BoolVar(&flagExtracting, "e",	false,	"Preform extract on compressed data.")
	flag.BoolVar(&flagCompressing,	"c",	false,	"Preform compress on file.")

	flag.Parse()

	if flagExtracting && flagCompressing {
		fmt.Println("Error: Cannot use two operations.")
		os.Exit(1)
	}

	if flagInput == "" || flagOutput == "" {
		fmt.Println("Error: No files specified.")
		os.Exit(1)
	}

	if !flagExtracting && !flagCompressing {
		fmt.Println("Error: No operation specified.")
		os.Exit(1)
	}
}
func main() {
	in, err := os.Open(flagInput)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(flagOutput)
	if err != nil {
		panic(err)
	}


	if flagExtracting {
		err = imageCompress.ExtractImage(in, out)
	} else if flagCompressing {
		err = imageCompress.GenerateImage(in, out)
	} else {
		flag.PrintDefaults()
	}

	if err != nil {
		panic(err)
	}
}
