package main

import "fmt"

func main() {
	var w Writer
	w = new(ByteCounter)
	count, _ := w.Write("a")
	fmt.Printf("count: %d\n", count)
}
