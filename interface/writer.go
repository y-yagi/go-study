package main

type Writer interface {
	Write(str string) (n int, err error)
}
