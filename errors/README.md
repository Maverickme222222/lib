# errors
Package errors provides convenience functions to handle errors.

Download:
```shell
go get github.com/kappapay/backend/lib/golang/src/kappa/errors
```

* * *
Package errors provides convenience functions to handle errors.
Errors can be created using the Error struct and can be printed using `%s` when formatting the error to be logged.
Errors should be embedded in the other errors using `%w` to preserve the underlying error and can be compared
to determine if it's of a specific error Type using `Is` function. The `As` function can be used to determine if the
error is in fact a gokit error and the `Unwrap function can be used to get to the original error.

Examples: (in these examples kiterror is the backend/errors package. The standard library errors package is used as errors)

``` golang
var customErr = fmt.Errorf("an error occured")

func TestErrors(t *testing.T) {
	err := ErrorTestHelper1()
	if !errors.Is(err, kiterror.ErrSystem) {
		t.Errorf("expected system error")
	}

	var ee kiterror.Error
	if !errors.As(err, &ee) {
		t.Errorf("expected system error type")
	}

	e := errors.Unwrap(err)
	if e, ok := e.(kiterror.Error); !ok || e.Err != customErr {
		t.Errorf("expected error to be system error")
	}
}

func ErrorTestHelper1() error {
	err := ErrorTestHelper2()
	return fmt.Errorf("wrapping the first error here: %w", err)
}

func ErrorTestHelper2() error {
	return kiterror.Error{
		Err:       customErr,
		Type:      kiterror.ErrSystem,
		Message:   "system error",
		Timestamp: time.Now(),
	}
}
```

In addition when a service makes a grpc call to another service returning the error to the caller as kiterror.Error
is automatically transformed using gokit's grpc server and client interceptors. For example this makes it so that an error
of type errors.NotFound that is returned from one service can be examined by the caller as follows:

```
if ok := kiterror.Is(err, kiterror.NotFound); ok {
     // the error is of type NotFound
}
```



* * *
