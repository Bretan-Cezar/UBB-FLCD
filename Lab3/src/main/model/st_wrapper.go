package model

type STWrapper struct {
	Ids             SymbolTable[string]
	IntConstants    SymbolTable[int64]
	StringConstants SymbolTable[string]
}
