package main

import (
	"errors"
	"fmt"
	"os"
	"io/ioutil"
)

// ExpectedArgCount is the number of expected parameters, including the path.
const ExpectedArgCount int = 2

// FileNameArg is the expected index of the input file name.
const FileNameArg int = 1

func main() {

	theFileName, theParamError := handleArguments(os.Args)

	if theParamError != nil {
		fmt.Println( "Error: " + theParamError.Error() )
	}

	fmt.Println("The Filename: " + theFileName )
	
	theFileContents := read( theFileName )

	fmt.Println(theFileContents)

	return
}

// Pulls the bytes from the file and converts it into a string.
func read(inFileName string) (string) {
	var outFileContents string
	theFileBytes, theFileReadError := ioutil.ReadFile( inFileName );

	// If the byte read didn't provide an error then convert the contents to the string.
	if theFileReadError == nil {
		outFileContents = string(theFileBytes)
	} else {
		fmt.Println( "Error: " + theFileReadError.Error() )
	}

	return outFileContents
}

// Parses the command line arguments for validity.
func handleArguments(inParams []string) (string, error) {
	var output string
	var outputErr error

	// Only two parameters expected.
	if len(inParams) != ExpectedArgCount {
		output = ""
		outputErr = errors.New("Bad Arguments: Two parameters expected including path")
	} else {
		output = inParams[FileNameArg]
		outputErr = nil
	}
	return output, outputErr

}
