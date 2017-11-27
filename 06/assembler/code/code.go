package code

import (
	"strconv"
)

//func divmod(numerator, denominator int64) (quotient, remainder int64) {
//	quotient = numerator / denominator // integer division, decimals are truncated
//	remainder = numerator % denominator
//	return
//}

//func DecTo15BitBinary(num string) string {
//	n := strconv.ParseInt(num, 10, 64)
//	var result string
//	var r int
//
//	for n > 0 {
//		n, r = divmod(n, 2)
//		result = strconv.FormatInt(r) + result
//	}
//
//	missingBits := 15 - len(result)
//
//	for i=0; i < missingBits; i++ {
//		result = "0" + result
//	}
//
//	result = "0" + result
//
//	return result
//}

func DecTo15BitBinary(num string) string {
	n, _ := strconv.ParseInt(num, 10, 64)
	result := strconv.FormatInt(n, 2)


	missingBits := 15 - len(result)
	for i:=0; i < missingBits; i++ {
		result = "0"+result
	}

	return result
}

func Comp(section string) string {
	switch section {
	case "0":
		return "0101010"
	case "1":
		return "0111111"
	case "-1":
		return "0111010"
	case "D":
		return "0001100"
	case "A":
		return "0110000"
	case "M":
		return "1110000"
	case "!D":
		return "0001101"
	case "!A":
		return "0110001"
	case "!M":
		return "1110001"
	case "-D":
		return "0001111"
	case "-A":
		return "0110011"
	case "-M":
		return "1110011"
	case "D+1":
		return "0011111"
	case "A+1":
		return "0110111"
	case "M+1":
		return "1110111"
	case "D-1":
		return "0001110"
	case "A-1":
		return "0110010"
	case "M-1":
		return "1110010"
	case "D+A":
		return "0000010"
	case "D+M":
		return "1000010"
	case "D-A":
		return "0010011"
	case "D-M":
		return "1010011"
	case "A-D":
		return "0000111"
	case "M-D":
		return "1000111"
	case "D&A":
		return "0000000"
	case "D&M":
		return "1000000"
	case "D|A":
		return "0010101"
	case "D|M":
		return "1010101"
	default:
		return "" //error, maybe invalid input
	}
}

func Dest (section string) string {
	switch section {
	case "":
		return "000"
	case "M":
		return "001"
	case "D":
		return "010"
	case "MD":
		return "011"
	case "A":
		return "100"
	case "AM":
		return "101"
	case "AD":
		return "110"
	case "AMD":
		return "111"
	default:
		return "" //error maybe invalid input
	}
}

func Jump (section string) string {
	switch section {
	case "":
		return "000"
	case "JGT":
		return "001"
	case "JEQ":
		return "010"
	case "JGE":
		return "011"
	case "JLT":
		return "100"
	case "JNE":
		return "101"
	case "JLE":
		return "110"
	case "JMP":
		return "111"
	default:
		return "" //error maybe invalid input
	}
}

