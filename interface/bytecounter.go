package main

type ByteCounter int

func (c *ByteCounter) Write(str string) (int, error) {
	return len(str), nil
}
