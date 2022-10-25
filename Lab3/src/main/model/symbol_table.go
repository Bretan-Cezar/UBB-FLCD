package model

import (
	"errors"
	"fmt"
)

type SymbolTable[V Variant] [SIZE][]V

func (st *SymbolTable[V]) GetHashAndIndex(value1 V) (int, int, error) {

	hash, err := ToyHashFunction(value1)

	if err != nil {

		return SIZE, SIZE, err
	}

	for index, value2 := range st[hash] {

		b, _ := equals(value1, value2)

		if b {

			return hash, index, nil
		}
	}

	return SIZE, SIZE, errors.New("undefined symbol")
}

func (st *SymbolTable[V]) SetSymbol(value1 V) (int, int, error) {

	hash, err := ToyHashFunction(value1)

	if err != nil {

		return SIZE, SIZE, err
	}

	for _, value2 := range st[hash] {

		b, _ := equals(value1, value2)

		if b {

			return SIZE, SIZE, errors.New("symbol already defined")
		}
	}

	st[hash] = append(st[hash], value1)

	return hash, len(st[hash]) - 1, nil
}

func (st *SymbolTable[V]) HasValue(value1 V) (bool, error) {

	hash, err := ToyHashFunction(value1)

	if err != nil {

		return false, err
	}

	for _, value2 := range st[hash] {

		b, _ := equals(value1, value2)

		if b {

			return true, nil
		}
	}

	return false, nil
}

func (st *SymbolTable[V]) FindByHashAndIndex(hash int, index int) (V, error) {

	if index >= len(st[hash]) {

		var v V
		return v, errors.New("index out of range")
	}

	return st[hash][index], nil
}

func (st *SymbolTable[V]) String() (str string) {

	str = "SymbolTable{ "

	for h := 0; h < 4096; h++ {

		if len(st[h]) != 0 {

			for i, sym := range st[h] {

				sym_i, ok := any(sym).(int64)

				if ok {

					str += fmt.Sprintf("(hash=%d; index=%d; symbol=%d); ", h, i, sym_i)

				} else {

					str += fmt.Sprintf("(hash=%d; index=%d; symbol=%s); ", h, i, sym)
				}
			}
		}
	}

	str += " }"

	return
}
