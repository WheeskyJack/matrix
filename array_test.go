// array_test.go
package arrays

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	SetArrayStartIndexToOne()
	fmt.Println("Launching test suite with startOfArrayIndex=", startOfArrayIndex)
	retCode := m.Run()
	os.Exit(retCode)
}

func failOnNil(t *testing.T, err error, msg string) {
	if err == nil {
		fmt.Println(msg)
		t.Fail()
	}
}

func failOnErr(t *testing.T, err error, msg string) {
	if err != nil {
		fmt.Println(msg + " " + err.Error())
		t.Fail()
	}
}

func failOnMatch(t *testing.T, exp, act interface{}, msg string) {
	if exp == act {
		fmt.Println(msg)
		t.Fail()
	}
}

func failOnMisMatch(t *testing.T, exp, act interface{}, msg string) {
	if exp != act {
		fmt.Println(msg+" : expected ", exp, " got ", act)
		t.Fail()
	}
}

func TestInsertOnNonEmptyArray(t *testing.T) {
	a := []int{1, 2, 4, 6}
	b, e := Insert(a, 3, 99) //simple insert
	failOnMisMatch(t, 99, b[2], "")
	failOnErr(t, e, "")

	b, e = Insert(a, 1, 99) //insert at begining
	failOnMisMatch(t, 99, b[0], "")
	failOnErr(t, e, "")

	b, e = Insert(a, len(a)+1, 199) //insert in the end
	failOnMisMatch(t, 199, b[len(a)], "")
	failOnErr(t, e, "")

	_, e = Insert(a, 7, 99) //out of boud test
	failOnNil(t, e, "")
}

func TestInsertOnEmptyArray(t *testing.T) {
	a := []int{}
	_, e := Insert(a, 2, 99) //out of boud test
	failOnNil(t, e, "")

	b, e := Insert(a, 1, 99)
	failOnMisMatch(t, 99, b[0], "") //insert at begining
	failOnErr(t, e, "")

	b, e = Insert(a, len(a)+1, 99)
	failOnMisMatch(t, 99, b[len(a)], "") //insert in the end
	failOnErr(t, e, "")
}

func TestDeleteOnNonEmptyArray(t *testing.T) {
	a := []int{1, 2, 4, 6}
	b, e := Delete(a, 3) //simple delete
	failOnMatch(t, 4, b[2], "")
	failOnErr(t, e, "")

	b, e = Delete(a, 1) //delete the first element
	failOnMatch(t, 1, b[0], "")
	failOnErr(t, e, "")

	b, e = Delete(a, len(a)) //delete the last element
	failOnMatch(t, 6, b[len(b)-1], "")
	failOnErr(t, e, "")

	_, e = Delete(a, -5) //delete the negative index element
	failOnNil(t, e, "")

	_, e = Delete(a, 6) //out of bound delete
	failOnNil(t, e, "")
}

func TestDeleteOnEmptyArray(t *testing.T) {
	a := []int{}
	_, e := Delete(a, 3) //simple delete
	failOnNil(t, e, "")

	_, e = Delete(a, 1) //delete the first element
	failOnNil(t, e, "")

	_, e = Delete(a, len(a)) //delete the last element
	failOnNil(t, e, "")

	_, e = Delete(a, -5) //delete the negative index element
	failOnNil(t, e, "")
}

func TestReverseOnEvenLengthArray(t *testing.T) {
	a := []int{1, 2, 4, 6}
	b := Reverse(a)
	failOnMisMatch(t, b[0], 6, "")
	failOnMisMatch(t, b[1], 4, "")
	failOnMisMatch(t, b[2], 2, "")
	failOnMisMatch(t, b[3], 1, "")
}

func TestReverseOnOddLengthArray(t *testing.T) {
	a := []int{1, 2, 4}
	b := Reverse(a)
	failOnMisMatch(t, b[0], 4, "")
	failOnMisMatch(t, b[1], 2, "")
	failOnMisMatch(t, b[2], 1, "")
}

func TestReverseOnTrivialLengthArray(t *testing.T) {
	a := []int{} //zero length array
	b := Reverse(a)
	failOnMisMatch(t, len(a), len(b), "")

	a = []int{5} //unit length array
	b = Reverse(a)
	failOnMisMatch(t, a[0], b[0], "")
}

func TestReverseOnPalindrome(t *testing.T) {
	a := []int{5, 1, 4, 1, 5} //palindrome
	b := Reverse(a)
	size := len(a)
	for i, v := range b {
		failOnMisMatch(t, a[size-1-i], v, "")
	}
}

func TestSearchOnNonEmptyArray(t *testing.T) {
	a := []int{1, 2, 4, 6, 2}
	b := Search(a, 1)
	failOnMisMatch(t, 1, b[0], "")
	failOnMisMatch(t, 1, len(b), "")

	b = Search(a, 2)
	failOnMisMatch(t, 2, b[0], "")
	failOnMisMatch(t, 5, b[1], "")
	failOnMisMatch(t, 2, len(b), "")

	b = Search(a, 88)
	failOnMisMatch(t, 0, len(b), "")
}

func TestSearchOnEmptyArray(t *testing.T) {
	a := []int{}
	b := Search(a, 1)
	failOnMisMatch(t, 0, len(b), "")
}

func TestIterFunc(t *testing.T) {
	a := []int{1, 2, 4, 6}
	f := func(x int) int {
		x *= x
		return x
	}
	b := IterFunc(a, f)
	for i, v := range b {
		failOnMisMatch(t, a[i]*a[i], v, "")
	}
}

func TestSum(t *testing.T) {
	a := []int{1, 2, 4, 6}

	s := Sum(a)
	failOnMisMatch(t, 13, s, "")

	a = []int{}
	s = Sum(a)
	failOnMisMatch(t, 0, s, "")
}

func TestAddArr(t *testing.T) {
	a := []int{1, 2, 4, 6}
	b := []int{1, 2, 4, 6}
	e := []int{2, 4, 8, 12}
	s, err := AddArr(a, b)

	failOnErr(t, err, "")
	for i, v := range s {
		failOnMisMatch(t, e[i], v, "")
	}

	c := []int{1, 2}
	_, err = AddArr(a, c)
	failOnNil(t, err, "")

	z1, z2 := []int{}, []int{}
	zs, err := AddArr(z1, z2)
	failOnMisMatch(t, 0, len(zs), "")
	failOnErr(t, err, "")
}
