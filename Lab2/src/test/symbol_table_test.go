package test

import (
	"UBB-FLCD/Lab2/src/main/model"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymbolTable(t *testing.T) {

	idTable := model.SymbolTable[string]{}
	intConstTable := model.SymbolTable[int64]{}
	stringConstTable := model.SymbolTable[string]{}

	h, i, err := idTable.SetSymbol("var1")

	assert.Equal(t, nil, err)

	_, _, err = idTable.SetSymbol("var1")

	assert.Error(t, err, "symbol already defined")

	h, i, err = idTable.SetSymbol("var2")

	assert.Equal(t, nil, err)

	h, i, err = intConstTable.SetSymbol(23)

	assert.Equal(t, nil, err)
	assert.Equal(t, h, 23)
	assert.Equal(t, i, 0)

	h, i, err = intConstTable.SetSymbol(12708)
	assert.Equal(t, nil, err)
	assert.Equal(t, h, 420)
	assert.Equal(t, i, 0)

	b, err := intConstTable.HasValue(12708)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, b)

	val, err := intConstTable.FindByHashAndIndex(h, i)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(12708), val)

	_, err = intConstTable.FindByHashAndIndex(h, 2)

	assert.Error(t, err, "index out of range")

	h, i, err = stringConstTable.SetSymbol("sample text")

	assert.Equal(t, nil, err)
	fmt.Println(h, " ", i)

	_, _, err = stringConstTable.SetSymbol("sample text")

	assert.Error(t, err, "symbol already defined")
}
