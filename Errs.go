// A very simple library that helps with formatting error messages. It
// provides a way to have consistent, descriptive errors.
package customerr

import (
	"errors"
	"fmt"
	"strings"
)

type (
	WrapListVal struct {
		ItemName string
		Item     any
	}
)

// Wraps an error with a predetermined format, as shown below.
//
//	<original error>
//	  |- <wrapped information>
//
// This allows for consistent error formatting.
func Wrap(origErr error, fmtStr string, vals ...any) error {
	fmtStrWithErr := fmt.Sprintf("%%w\n  |- %s", fmtStr)
	args := []interface{}{origErr}
	return fmt.Errorf(fmtStrWithErr, append(args, vals...)...)
}

// Wraps an error with a predetermined format, as shown below.
//
//	<wrapped information>
//	  |- <original error>
//
// This allows for consistent error formatting.
func InverseWrap(origErr error, fmtStr string, vals ...any) error {
	fmtStrWithErr := fmt.Sprintf("%s\n  |- %%w", fmtStr)
	return fmt.Errorf(fmtStrWithErr, append(vals, origErr)...)
}

// Wraps an error with a predetermined format, as shown below,
//
//	<original error>
//	  |- <description>
//	  |- value1 name (value1 type): value1
//	  |- value2 name (value2 type): value2
//
// This allows for consistent error formatting.
func WrapValueList(
	origErr error,
	description string,
	valsList []WrapListVal,
) error {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%%w\n  |- Description: %s", description))
	if len(valsList) > 0 {
		sb.WriteByte('\n')
	}
	for i, v := range valsList {
		if stringer, ok := v.Item.(fmt.Stringer); ok {
			sb.WriteString(fmt.Sprintf(
				"  |- %s (%T): %s", v.ItemName, v.Item, stringer,
			))
		} else {
			sb.WriteString(fmt.Sprintf(
				"  |- %s (%T): %+v", v.ItemName, v.Item, v.Item,
			))
		}
		if i+1 < len(valsList) {
			sb.WriteByte('\n')
		}
	}
	return fmt.Errorf(sb.String(), origErr)
}

// Unwraps an error. A simple helper function to provide a clean error interface
// in this module.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Given a list of errors it will append them with a predetermined format, as
// shown below.
//
//	<original first error>
//	  |- <wrapped information>
//	...
//	<original nth error>
//	  |- <wrapped information>
//
// This allows for consistent error formatting. Special cases are as follows:
//   - All supplied errors are nil: The returned value will be nil.
//   - Only one of the supplied errors is nil: The returned value will be the error that is not nil.
//   - Multiple errors are not nil: The returned error will be a [MultipleErrorsOccurred] error with all of the sub-errors wrapped in it following the above format.
func AppendError(errs ...error) error {
	var rv error
	cntr := 0
	for _, e := range errs {
		if e != nil {
			if cntr == 0 {
				rv = e
			} else {
				rv = fmt.Errorf("%w\n%w", rv, e)
			}
			cntr++
		}
	}
	return rv
}
