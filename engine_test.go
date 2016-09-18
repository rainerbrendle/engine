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

func TestPutPowerData(t *testing.T) {

	var in_value = "hello world"

	db, err := GetDatabase("rainer")

	if err != nil {
		fmt.Printf("PANIC %#v\n", err)
		t.FailNow()
	}

	err = db.PutPowerData("TESTX", in_value)

	if err != nil {
		fmt.Printf("PANIC %#v\n", err)
		t.FailNow()
	}

	fmt.Printf("TESTX put with value %v\n", in_value)

}

func TestPutPowerData10(t *testing.T) {
	testPutPowerDataX(t, 10)
}

func TestPutPowerData1000(t *testing.T) {
	testPutPowerDataX(t, 1000)
}

func testPutPowerDataX(t *testing.T, count int) {

	var in_value = "Hello"

	db, err := GetDatabase("rainer")

	if err != nil {
		fmt.Printf("PANIC %#v\n", err)
		t.FailNow()
	}

	for i := 0; i < count; i++ {

		err = db.PutPowerData("TESTX", fmt.Sprintf("%v %v", in_value, i))

		if err != nil {
			fmt.Printf("PANIC %#v\n", err)
			t.FailNow()
		}
	}

	fmt.Printf("TESTX put with value %v (%v times)\n", in_value, count)

}

func TestGetPowerData(t *testing.T) {

	db, errdb := GetDatabase("rainer")

	if errdb != nil {
		fmt.Printf("PANIC %#v\n", errdb)
		t.FailNow()
	}

	out_value, err := db.GetPowerData("TESTX")

	if err != nil {
		fmt.Printf("PANIC %#v\n", err)
		t.FailNow()
	}

	fmt.Printf("TESTX get  %v\n", out_value)

}
