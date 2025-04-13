package sberr

import (
	"errors"
	"fmt"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

var testErrOne = errors.New("err1")
var testErrTwo = errors.New("err2")

func TestInverseWrap(t *testing.T) {
	e := InverseWrap(testErrOne, "%d", 100)
	sbtest.Eq(t,
		fmt.Sprintf("100\n  |- %s", testErrOne.Error()), e.Error(),
	)
}

func TestWrap(t *testing.T) {
	e := Wrap(testErrOne, "%d", 100)
	sbtest.Eq(t,
		fmt.Sprintf("%s\n  |- 100", testErrOne.Error()), e.Error(),
	)
}

func TestWrapValueList(t *testing.T) {
	e := WrapValueList(
		testErrOne,
		"Val outside range",
		WrapListVal{ItemName: "value", Item: 10},
	)
	sbtest.Eq(
		t,
		fmt.Sprintf(
			"%s\n  |- Description: Val outside range\n  |- value (int): 10",
			testErrOne.Error(),
		),
		e.Error(),
	)
}

func TestUnwrap(t *testing.T) {
	e := Wrap(testErrOne, "%d", 100)
	sbtest.Eq(t, testErrOne, Unwrap(e))
	sbtest.ContainsError(t, testErrOne, Unwrap(e))
}

func TestAppendErrorTwoErrors(t *testing.T) {
	e := AppendError(testErrOne, testErrTwo)
	sbtest.Eq(
		t,
		fmt.Sprintf(
			"%s\n%s",
			testErrOne.Error(),
			testErrTwo.Error(),
		),
		e.Error(),
	)
}

func TestAppendErrorOnlyFirst(t *testing.T) {
	e := AppendError(testErrOne, nil)
	sbtest.Eq(t, testErrOne.Error(), e.Error())
	sbtest.Eq(t, testErrOne, e)
}

func TestAppendErrorOnlySecond(t *testing.T) {
	e := AppendError(nil, testErrOne)
	sbtest.Eq(t, testErrOne.Error(), e.Error())
	sbtest.Eq(t, testErrOne, e)
}

func TestAppendErrorBothNil(t *testing.T) {
	e := AppendError(nil, nil)
	sbtest.Nil(t, e)
}
