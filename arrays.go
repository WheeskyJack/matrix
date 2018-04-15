// arrays.go
package arrays

import (
	"errors"
)

var startOfArrayIndex = 0

func SetArrayStartIndexToOne() {
	startOfArrayIndex = 1
}

//Insert returns the array with element inserted at given position.
func Insert(a []int, pos, val int) (b []int, err error) {
	pos -= startOfArrayIndex
	if len(a) < pos || pos < 0 {
		b, err = a, errors.New("input array slice range smaller than given position")
	} else {
		b = append(b, a[:pos]...)
		b = append(b, val)
		b = append(b, a[pos:]...)
	}
	return
}

//Delete deletes the element in the array present at given postion
func Delete(a []int, pos int) (b []int, err error) {
	pos -= startOfArrayIndex
	if len(a) <= pos || pos < 0 {
		b, err = a, errors.New("input array slice range smaller than given position")
	} else {
		b = append(b, a[:pos]...)
		b = append(b, a[pos+1:]...)
	}
	return
}

//Reverse reverses the array of integer.
func Reverse(a []int) (b []int) {
	if len(a) <= 1 {
		b = a
	} else {
		size := len(a)
		b = make([]int, size)
		for i := 0; i <= len(a)/2; i++ {
			b[i], b[size-1-i] = a[size-1-i], a[i]
		}
	}
	return
}

//Search searches the given element value in array and
//returns all the indices matched
func Search(a []int, val int) (b []int) {
	for i, v := range a {
		if v == val {
			b = append(b, i+startOfArrayIndex)
		}
	}
	return
}

//IterFunc operates on each of the element of the array with given function
//and returns the array with corresponding output.
func IterFunc(a []int, f func(int) int) (b []int) {
	done := make(chan bool, len(a))
	b = make([]int, len(a))
	for i, v := range a {
		go parallelRuns(b, i, v, f, done)
	}
	for i := 0; i < len(a); i++ {
		<-done
	}
	return
}

func parallelRuns(b []int, ind, val int, f func(int) int, done chan bool) {
	b[ind] = f(val)
	done <- true
}

//Sum adds all the elements in an array
func Sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

//AddArr adds the two arrays.
func AddArr(a, b []int) (s []int, err error) {
	if len(a) != len(b) {
		err = errors.New("array length mismatch")
	} else {
		s = make([]int, len(a))
		for i, v := range a {
			s[i] = v + b[i]
		}
	}
	return
}
