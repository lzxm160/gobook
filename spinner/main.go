package main

import "fmt"
import "time"
func main() {
	go spinner(100*time.Millisecond)
	const n=45
	fibn:=fib(n)//slow
	
	fmt.Printf("\rfib(%d)=%d\n",n,fibn)
}
func spinner(delay time.Duration){
	for{
		for _,r := range `-/|\`{
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x<2{
		return x
	}
	return fib(x-1)+fib(x-2)
}
