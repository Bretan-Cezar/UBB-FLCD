package model

import "fmt"

type STWrapper struct {
	Ids             SymbolTable[string]
	IntConstants    SymbolTable[int64]
	StringConstants SymbolTable[string]
}

func (stw *STWrapper) String() string {

	return fmt.Sprintf("STWrapper{\n\tIDs: %s\n\tIntConstants: %s\n\tStringConstants: %s\n}",
		stw.Ids.String(),
		stw.IntConstants.String(),
		stw.StringConstants.String())
}
