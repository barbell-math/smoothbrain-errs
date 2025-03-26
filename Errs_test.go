package customerr

import (
	"fmt"
	"testing"

	"github.com/barbell-math/smoothbrain-test"
)

func TestInverseWrap(t *testing.T) {
	e := InverseWrap(ValOutsideRange, "%d", 100)
	smoothbraintest.Eq(t,
		fmt.Sprintf("100\n  |- %s", ValOutsideRange.Error()), e.Error(),
	)
}

func TestWrap(t *testing.T) {
	e := Wrap(ValOutsideRange, "%d", 100)
	smoothbraintest.Eq(t,
		fmt.Sprintf("%s\n  |- 100", ValOutsideRange.Error()), e.Error(),
	)
}

func TestWrapValueList(t *testing.T) {
	e := WrapValueList(ValOutsideRange, "Val outside range", []WrapListVal{
		{ItemName: "value", Item: 10},
	})
	smoothbraintest.Eq(
		t,
		fmt.Sprintf(
			"%s\n  |- Description: Val outside range\n  |- value (int): 10",
			ValOutsideRange.Error(),
		),
		e.Error(),
	)
}

func TestUnwrap(t *testing.T) {
	e := Wrap(ValOutsideRange, "%d", 100)
	smoothbraintest.Eq(t, ValOutsideRange, Unwrap(e))
	smoothbraintest.ContainsError(t, ValOutsideRange, Unwrap(e))
}

func TestAppendErrorTwoErrors(t *testing.T) {
	e := AppendError(ValOutsideRange, DimensionsDoNotAgree)
	smoothbraintest.Eq(
		t,
		fmt.Sprintf(
			"%s\n%s",
			ValOutsideRange.Error(),
			DimensionsDoNotAgree.Error(),
		),
		e.Error(),
	)
}

func TestAppendErrorOnlyFirst(t *testing.T) {
	e := AppendError(ValOutsideRange, nil)
	smoothbraintest.Eq(t, ValOutsideRange.Error(), e.Error())
	smoothbraintest.Eq(t, ValOutsideRange, e)
}

func TestAppendErrorOnlySecond(t *testing.T) {
	e := AppendError(nil, ValOutsideRange)
	smoothbraintest.Eq(t, ValOutsideRange.Error(), e.Error())
	smoothbraintest.Eq(t, ValOutsideRange, e)
}

func TestAppendErrorBothNil(t *testing.T) {
	e := AppendError(nil, nil)
	smoothbraintest.Nil(t, e)
}
