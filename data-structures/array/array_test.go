package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray(10)

	arr.Append(1)
	if err := arr.Insert(5, 2); err != nil {
		t.Log(err)
	}

	arr.Print()

	if v, err := arr.Find(1); err != nil {
		t.Log(err)
	} else {
		fmt.Println("found data is ", v)
	}

	if err := arr.Delete(1); err != nil {
		t.Logf("delete val 1 failed with:%v", err)
	}

	if err := arr.DeleteByIndex(0); err != nil {
		t.Logf("delete index 0 failed with:%v", err)
	}
	arr.Print()
}
