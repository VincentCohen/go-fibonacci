package go_fibonacci_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoFibonacci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoFibonacci Suite")
}
