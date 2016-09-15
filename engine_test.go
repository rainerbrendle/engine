//
// Test suite for tsn.go
//

package engine

import (
	"fmt"
	"testing"
)

func TestNewTSN(t *testing.T) {

	var tsn int64

	db, err := GetDatabase("rainer")

	if err != nil {
		fmt.Printf("PANIC %#v\n", err)
		t.FailNow()
	}

	tsn, err = db.NewTSN()
	if err != nil {
		fmt.Printf("PANIC %#v\n", err)
		t.FailNow()
	}

	fmt.Printf("1 x NEW TSN %v\n", tsn)

}

func TestNewTSN100(t *testing.T) {

	var tsn int64

	db, err := GetDatabase("rainer")

	if err != nil {
		fmt.Printf("PANIC %#v\n", err)
		t.FailNow()
	}

	for i := 0; i < 100; i++ {

		tsn, err = db.NewTSN()
		if err != nil {
			fmt.Printf("PANIC %#v\n", err)
			t.FailNow()
		}
	}

	fmt.Printf("100 x NEW TSN %v\n", tsn)

}

func TestNewTSN10000(t *testing.T) {

	var tsn int64

	db, err := GetDatabase("rainer")

	if err != nil {
		fmt.Printf("PANIC %#v\n", err)
		t.FailNow()
	}

	for i := 0; i < 10000; i++ {

		tsn, err = db.NewTSN()
		if err != nil {
			fmt.Printf("PANIC %#v\n", err)
			t.FailNow()
		}
	}

	fmt.Printf("10000 x NEW TSN %v\n", tsn)

}
