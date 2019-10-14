package array

import (
	"errors"
	"fmt"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 *
 */

type Array struct {
	data   []int
	length uint
}

// get array length
func (a *Array) Len() uint {
	return a.length
}

// init array
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// print array
func (a *Array) Print() {
	fmt.Println(a.data)
}

// IsIndexOutOfRange
func (a *Array) IsIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return true
	}
	return false
}

// get a data by index
func (a *Array) Find(index uint) (int, error) {
	if a.IsIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

// append a data in array
func (a *Array) Append(data int) {
	a.data = append(a.data, data)
	a.length++
}

// insert a data in the index
func (a *Array) Insert(index uint, v int) error {
	// data is full
	if a.Len() == uint(cap(a.data)) {
		return errors.New("array is full")
	}

	// index is out of range
	if a.IsIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := a.Len(); i > index; i++ {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

// delete a data in array
func (a *Array) Delete(v int) error {
	var deleted bool
	for i, d := range a.data {
		if d == v {
			deleted = true
			for ; i < int(a.length); i++ {
				a.data[i] = a.data[i+1]
				break
			}
		}
	}
	if deleted {
		a.length--
		return nil
	}
	return errors.New("not found in array")
}

// delete a data by index
func (a *Array) DeleteByIndex(index uint) error {
	if a.IsIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := index; i < a.length; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return nil
}
