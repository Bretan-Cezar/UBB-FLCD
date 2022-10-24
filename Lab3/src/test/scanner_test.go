package test

import (
	"UBB-FLCD/Lab3/src/main/model"
	"fmt"
	"os"
	"testing"
)

func TestScanner(t *testing.T) {

	var p1 *os.File
	var err error

	p1, err = os.Open("p1.in")
	//p2, err = os.Open("UBB-FLCD/Lab3/res/p2.in")
	//p3, err = os.Open("UBB-FLCD/Lab3/res/p3.in")
	//p1err, err = os.Open("UBB-FLCD/Lab3/res/p1err.in")

	pif, _, err := model.Scan(p1)

	if err != nil {

		t.Error(err)
		return
	}

	for _, entry := range pif {

		fmt.Println(entry)
	}

}
