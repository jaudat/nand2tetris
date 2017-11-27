package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"./parse"
	"./code"
	"./symbolstable"
)

func main() {
	allFiles := os.Args[1:]

	for _, file := range allFiles {
		lines := populateSymbolTable(file)

		nameWithoutExtension := strings.Split(file, ".")
		hackFileName := nameWithoutExtension[0]+".hack"
		assemble(hackFileName, lines)
	}
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func populateSymbolTable(fileName string) []string {
	// create symboltree and set predefined symbols
	st := symbolstable.NewSymbolsTable()

	//Add label symbols in this pass
	afterFirstPass := firstPass(fileName, st)

	//Add the variable symbols in this pass
	afterSecondPass := secondPass(afterFirstPass, st)

	return afterSecondPass
}

func firstPass(fileName string, symTable *symbolstable.SymbolsTable) []string {
	// get file to read from
	sourceFile, err := os.Open(fileName);
	check(err)

	defer sourceFile.Close()

	scanner := bufio.NewScanner(sourceFile)

	programLineAddr := 0
	codeLines := []string{}

	// read line by line
	for scanner.Scan() {
		line := scanner.Text()
		// remove emptylines and lines that are comments
		if line != "" && !strings.HasPrefix(line, "//") {
			if strings.HasPrefix(line, "(") {
				// add label variables to symbolTable
				label := parse.LabelVariable(line)
				symTable.Insert(label, programLineAddr)

			} else {
				//remove inline comments
				splitCommentAndInstructions := strings.Split(line, "//")
				instructions := strings.TrimSpace( splitCommentAndInstructions[0] )
				codeLines = append(codeLines, instructions)

				programLineAddr += 1
			}
		}
	}

	check(scanner.Err())
	return codeLines
}

func secondPass(lines []string, symTable *symbolstable.SymbolsTable) []string {
	newLines := []string{}

	for _, line := range lines {
		// Check if A Command, if it is then if variable used lookup variable found in symboltable,
		// if found then replace variable with decimal value of address, if not found then add to
		// symbolTable in unused address space from 16, and then repace variable with decimal value
		// of that address

		if line[0] == '@' {
			value := parse.Address(line)
			_, err := strconv.Atoi(value) //converts to int
			if err != nil {
				// then it errored out meaning it is symbol and needs to be handled
				newLines = append(newLines, "@" + handleSymbol(value, symTable))
			} else {
				newLines = append(newLines, line)
			}
		} else {
			newLines = append(newLines, line)
		}
	}

	return newLines
}

func handleSymbol(symbol string, symTable *symbolstable.SymbolsTable) string {
	// Check if symbol is in SymbolTable if it is then return its address
	// otherwise put it in the SymbolTable and then return it's address

	address := symTable.Get(symbol)
	if address == "" {
		address = symTable.Add(symbol)
	}
	return address
}


func assemble(fileName string, lines []string) {
		// create file to write to
	destFile, err := os.Create(fileName)
	defer destFile.Close()

	check(err)

	w := bufio.NewWriter(destFile)

	// read line by line
	for _, line := range lines{
		//convert instructions to machine code
		machineCode := convertToMachineCode(line)

		// and then write machine code to the HACK file
		w.WriteString(machineCode + "\n")
	}

	w.Flush()

}

func convertToMachineCode(instruction string) string {

	if instruction[0] == '@' {
		// then it is A instruction and will look like following:
		// 0vvv vvvv vvvv vvvv
		// where v are values that are either 0 or 1 and represent
		// the address number in binary
		num := parse.Address(instruction)
		return "0"+code.DecTo15BitBinary(num)
	} else {
		// then it is C instruction and will look like following:
		// 111a c1 c2 c3 c4 c5 c6 d1 d2 d3 j1 j2 j3
		compStr:= parse.Comp(instruction)
		destStr := parse.Dest(instruction)
		jumpStr := parse.Jump(instruction)

		compCode := code.Comp(compStr)
		destCode := code.Dest(destStr)
		jumpCode := code.Jump(jumpStr)

		return "111" + compCode + destCode + jumpCode
	}

}

