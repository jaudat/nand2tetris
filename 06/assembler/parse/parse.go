package parse

import (
	"strings"
)


func LabelVariable(line string) string {
	return line[1:len(line)-1]
}

func Address(line string) string {
	return line[1:len(line)]
}

func Comp(line string) string {

	equalSignIndex := strings.Index(line, "=")

	// if no equalSign then the value returned will be -1 and the following code
	// will change it to 0 as wanted as no equalSign means there are no Dest Instructions
	// if there is an equal sign we will have it's index which we will increment by 1, as
	// the comp instructions will start after the equal sign
	equalSignIndex += 1

	semicolonSignIndex := strings.Index(line, ";")
	if semicolonSignIndex == -1 {
		semicolonSignIndex = len(line)
	}

	return line[equalSignIndex:semicolonSignIndex]


}

func Dest(line string) string {
	equalSignIndex := strings.Index(line, "=")
	if equalSignIndex == -1 {
		return "" //there is no equal sign therefore no dest
	}

	return line[0:equalSignIndex] //substring from beginning of string till just before the equal sign
}

func Jump(line string) string {
	semicolonSignIndex := strings.Index(line, ";")
	if semicolonSignIndex == -1 {
		return "" //there is no semicolon therefore no jump instruction
	}
	semicolonSignIndex += 1
	return line[semicolonSignIndex:len(line)] //substring from after the semicolon to end of string
}

