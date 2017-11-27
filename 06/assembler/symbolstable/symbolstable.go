package symbolstable

import (
	"strconv"
)



type SymbolsTable struct {
	nextVarAddrSpace int
	symMap map[string]string
	symList []string
}


func NewSymbolsTable() *SymbolsTable {
	st := new(SymbolsTable)

	st.nextVarAddrSpace = 16

	st.symMap = make(map[string]string)
	st.symMap["R0"] = "0"
	st.symMap["R1"] = "1"
	st.symMap["R2"] = "2"
	st.symMap["R3"] = "3"
	st.symMap["R4"] = "4"
	st.symMap["R5"] = "5"
	st.symMap["R6"] = "6"
	st.symMap["R7"] = "7"
	st.symMap["R8"] = "8"
	st.symMap["R9"] = "9"
	st.symMap["R10"] = "10"
	st.symMap["R11"] = "11"
	st.symMap["R12"] = "12"
	st.symMap["R13"] = "13"
	st.symMap["R14"] = "14"
	st.symMap["R15"] = "15"
	st.symMap["SCREEN"] = "16384"
	st.symMap["KBD"] = "24576"
	st.symMap["SP"] = "0"
	st.symMap["LCL"] = "1"
	st.symMap["ARG"] = "2"
	st.symMap["THIS"] = "3"
	st.symMap["THAT"] = "4"

	st.symList = make([]string, 24577, 24577)
	st.symList[0] = "R0"
	st.symList[1] = "R1"
	st.symList[2] = "R2"
	st.symList[3] = "R3"
	st.symList[4] = "R4"
	st.symList[5] = "R5"
	st.symList[6] = "R6"
	st.symList[7] = "R7"
	st.symList[8] = "R8"
	st.symList[9] = "R9"
	st.symList[10] = "R10"
	st.symList[11] = "R11"
	st.symList[12] = "R12"
	st.symList[13] = "R13"
	st.symList[14] = "R14"
	st.symList[15] = "R15"

	for i := 16384; i < 24576; i++ {
		st.symList[i] = "SCREEN"
	}

	st.symList[24576] = "KBD"

	return st
}

func (st *SymbolsTable) Insert(label string, index int) {
	//add to Slice at index
	//listCap := cap(st.symList)
	//
	//if index > listCap {
	//	// if index that the symbol is to be inserted at is greater then the
	//	// capacity of the array, then increase the capacity of array
	//	newSymList := make([]string, 2*listCap, 2*listCap)
	//	for i := 0; i < len(st.symList); i++ {
	//		newSymList[i] = st.symList[i]
	//	}
	//	st.symList = newSymList
	//}
	//
	//st.symList[index] = label

	//add to hashmap
	st.symMap[label] = strconv.Itoa(index)
}

func (st *SymbolsTable) Get(label string) string {
	if val, ok := st.symMap[label]; ok {
		return val
	}

	return ""
}

func (st *SymbolsTable) Add(label string) string {
	address := 0

	//Add to first empty register found after 16
	for index := 16; index < len(st.symList); index++ {
		value := st.symList[index]
		if value == "" {
			st.symList[index] = label
			address = index
			break
		}
	}

	if address == 0 {
		// then no place empty to insert so need to expand memory and add at end register
		st.symList = append(st.symList, label)
		address = len(st.symList)-1
	}

	//Add to hashmap
	addrStr := strconv.Itoa(address)
	st.symMap[label] = addrStr
	return addrStr
}

