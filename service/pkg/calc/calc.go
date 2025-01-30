package calc; import __sealights__ "github.com/liornabat-sealights/go-calc-demo/__sealights__"

import "fmt"

func Add(a, b float64) float64 {__sealights__.TraceFunc("642be340b378ec794b");
	fmt.Println("Add", a, b)
	return a + b
}

func Subtract(a, b float64) float64 {__sealights__.TraceFunc("a2c53af5233101e409");
	fmt.Println("Subtract", a, b)
	return a - b
}

func Multiply(a, b float64) float64 {__sealights__.TraceFunc("721569a6aee7553cd8");
	fmt.Println("Multiply", a, b)
	return a * b
}

func Divide(a, b float64) float64 {__sealights__.TraceFunc("b271c2366633b74b36");
	fmt.Println("Divide", a, b)
	return a / b
}
