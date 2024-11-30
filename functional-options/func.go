package functional_options

import "fmt"

type myFuncConfig struct {
	namedArg1 string
	namedArg2 bool
}

type MyFuncOption func(*myFuncConfig)

func WithNamedArg1(val string) MyFuncOption {
	return func(config *myFuncConfig) {
		config.namedArg1 = val
	}
}

func WithNamedArg2(val bool) MyFuncOption {
	return func(config *myFuncConfig) {
		config.namedArg2 = val
	}
}

func MyFunc(arg1, arg2 string, opts ...MyFuncOption) {
	c := &myFuncConfig{
		namedArg1: arg1,
		namedArg2: true,
	}
	for _, opt := range opts {
		opt(&myFuncConfig{})
	}
	fmt.Println(c.namedArg1, c.namedArg2)
}
