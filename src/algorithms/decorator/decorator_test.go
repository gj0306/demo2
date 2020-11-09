package decorator

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	fmt.Println(f1(5000))
	fmt.Println(f1(10000))
	fmt.Println(f1(50000))
}

func TestDecorator(t *testing.T) {
	f:= Decorator(f1)
	f(100000)
	f(200000)
}