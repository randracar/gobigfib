package main

import (
	"fmt"
	"math/big"
	"time"
	"os"
)

func fib(max *big.Int) {
	var mem, num1, num2, c big.Int
	one := big.NewInt(1)

	num1.SetInt64(0)
	num2.SetInt64(1)
	for num1.Cmp(max) <= 0 {
		fmt.Println(&num1)
		mem.Set(&num1)
		mem.Add(&mem, &num2)
		num1.Set(&num2)
		num2.Set(&mem)
		c.Add(&c, one)
	}
}

func main(){
	start := time.Now()
	if len(os.Args) != 2 {
		panic("PANIC: INVALID CLI NUMBER OF ARGUMENTS")
	}
	
	// sets the max num for the sequence through cli
	max := new(big.Int)
	_, err := max.SetString(os.Args[1], 10)
	if !err {
		panic(err)
	}
	
	fmt.Println("Fibonacci up to ", max)
	fib(max)
	fmt.Println("Took ", time.Since(start), " to run")
}