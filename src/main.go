package main

import "fmt"

func main() {
	parseArgs()

	fmt.Printf("%q\n", returnSomething("something"))
}
