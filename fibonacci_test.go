package go_fibonacci_test

import (
. "github.com/onsi/ginkgo"
. "github.com/onsi/ginkgo/extensions/table"
. "github.com/onsi/gomega"

. "github.com/VincentCohen/go-fibonacci"
)

var _ = Describe("Fibonacci", func() {
	var fibonacci Fibonacci

	DescribeTable("Calculating Fibonacci numbers",
		func(steps int, expected int) {
			Expect(fibonacci.Fib(steps)).To(Equal(expected))
		},
		Entry("fib(0)", 0, 0),
		Entry("fib(1)", 1, 1),
		Entry("fib(2)", 2, 1),
		Entry("fib(3)", 3, 2),
		Entry("fib(4)", 4, 3),
		Entry("fib(5)", 5, 5),
	)

})

