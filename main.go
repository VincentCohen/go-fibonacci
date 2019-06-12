package go_fibonacci

type Fibonacci struct {}

func (f Fibonacci) Fib(steps int) int {

	if steps == 0 {
		return 0
	}

	if steps == 1 || steps == 2 {
		return 1
	}

	return f.Fib(steps-1) + f.Fib(steps-2)
}