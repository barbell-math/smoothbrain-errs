<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# sberr

```go
import "github.com/barbell-math/smoothbrain-errs"
```

A very simple library that helps with formatting error messages. It provides a way to have consistent, descriptive errors.

## Index

- [func AppendError\(errs ...error\) error](<#AppendError>)
- [func InverseWrap\(origErr error, fmtStr string, vals ...any\) error](<#InverseWrap>)
- [func Unwrap\(err error\) error](<#Unwrap>)
- [func Wrap\(origErr error, fmtStr string, vals ...any\) error](<#Wrap>)
- [func WrapValueList\(origErr error, description string, valsList ...WrapListVal\) error](<#WrapValueList>)
- [type WrapListVal](<#WrapListVal>)


<a name="AppendError"></a>
## func [AppendError](<https://github.com/barbell-math/smoothbrain-errs/blob/main/Errs.go#L95>)

```go
func AppendError(errs ...error) error
```

Given a list of errors it will append them with a predetermined format, as shown below.

```
<original first error>
  |- <wrapped information>
...
<original nth error>
  |- <wrapped information>
```

This allows for consistent error formatting. Special cases are as follows:

- All supplied errors are nil: The returned value will be nil.
- Only one of the supplied errors is nil: The returned value will be the error that is not nil.
- Multiple errors are not nil: The returned error will be a \[MultipleErrorsOccurred\] error with all of the sub\-errors wrapped in it following the above format.

<a name="InverseWrap"></a>
## func [InverseWrap](<https://github.com/barbell-math/smoothbrain-errs/blob/main/Errs.go#L36>)

```go
func InverseWrap(origErr error, fmtStr string, vals ...any) error
```

Wraps an error with a predetermined format, as shown below.

```
<wrapped information>
  |- <original error>
```

This allows for consistent error formatting.

<a name="Unwrap"></a>
## func [Unwrap](<https://github.com/barbell-math/smoothbrain-errs/blob/main/Errs.go#L78>)

```go
func Unwrap(err error) error
```

Unwraps an error. A simple helper function to provide a clean error interface in this module.

<a name="Wrap"></a>
## func [Wrap](<https://github.com/barbell-math/smoothbrain-errs/blob/main/Errs.go#L24>)

```go
func Wrap(origErr error, fmtStr string, vals ...any) error
```

Wraps an error with a predetermined format, as shown below.

```
<original error>
  |- <wrapped information>
```

This allows for consistent error formatting.

<a name="WrapValueList"></a>
## func [WrapValueList](<https://github.com/barbell-math/smoothbrain-errs/blob/main/Errs.go#L49-L53>)

```go
func WrapValueList(origErr error, description string, valsList ...WrapListVal) error
```

Wraps an error with a predetermined format, as shown below,

```
<original error>
  |- <description>
  |- value1 name (value1 type): value1
  |- value2 name (value2 type): value2
```

This allows for consistent error formatting.

<a name="WrapListVal"></a>
## type [WrapListVal](<https://github.com/barbell-math/smoothbrain-errs/blob/main/Errs.go#L12-L15>)



```go
type WrapListVal struct {
    ItemName string
    Item     any
}
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

## Helpful Developer Cmds

To build the build system:

```
go build -o ./bs/bs ./bs
```

The build system can then be used as usual:

```
./bs/bs --help
./bs/bs buildBs # Builds the build system!
```
