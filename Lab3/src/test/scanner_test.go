package test

import (
	"UBB-FLCD/Lab3/src/main/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestP1(t *testing.T) {

	var p1 *os.File
	var err error
	var pifStr string

	p1, err = os.Open("res/p1.in")

	pif, st, err := model.Scan(p1)

	assert.Equal(t, nil, err)

	pifStr = "PIF{\n"

	for _, entry := range pif {

		pifStr += "\t" + entry.String() + "\n"
	}

	pifStr += "}"

	_ = p1.Close()

	pifFile, err := os.Create("res/PIF.out")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = pifFile.Write([]byte(pifStr + "\n"))
	if err != nil {
		t.Error(err)
		return
	}

	_ = pifFile.Close()

	stFile, err := os.Create("res/ST.out")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = stFile.Write([]byte(st.String()))
	if err != nil {
		t.Error(err)
		return
	}

	_ = stFile.Close()

	t.Log("lexically correct\n")
}

func TestP2(t *testing.T) {

	var p2 *os.File
	var err error
	var pifStr string

	p2, err = os.Open("res/p2.in")

	pif, st, err := model.Scan(p2)

	assert.Equal(t, nil, err)

	pifStr = "PIF{\n"

	for _, entry := range pif {

		pifStr += "\t" + entry.String() + "\n"
	}

	pifStr += "}"

	_ = p2.Close()

	pifFile, err := os.Create("res/PIF.out")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = pifFile.Write([]byte(pifStr + "\n"))
	if err != nil {
		t.Error(err)
		return
	}

	_ = pifFile.Close()

	stFile, err := os.Create("res/ST.out")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = stFile.Write([]byte(st.String()))
	if err != nil {
		t.Error(err)
		return
	}

	_ = stFile.Close()

	t.Log("lexically correct\n")
}

func TestP3(t *testing.T) {

	var p3 *os.File
	var err error
	var pifStr string

	p3, err = os.Open("res/p3.in")

	pif, st, err := model.Scan(p3)

	assert.Equal(t, nil, err)

	pifStr = "PIF{\n"

	for _, entry := range pif {

		pifStr += "\t" + entry.String() + "\n"
	}

	pifStr += "}"

	_ = p3.Close()

	pifFile, err := os.Create("res/PIF.out")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = pifFile.Write([]byte(pifStr + "\n"))
	if err != nil {
		t.Error(err)
		return
	}

	_ = pifFile.Close()

	stFile, err := os.Create("res/ST.out")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = stFile.Write([]byte(st.String()))
	if err != nil {
		t.Error(err)
		return
	}

	_ = stFile.Close()

	t.Log("lexically correct\n")
}

func TestPErr(t *testing.T) {

	var p1err *os.File
	var err error

	p1err, err = os.Open("res/p1err.in")

	_, _, err = model.Scan(p1err)

	assert.Error(t, err)

	t.Logf("%s\n", err)

	_ = p1err.Close()
}