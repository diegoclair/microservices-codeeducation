package main

type somaPackage struct {
}

func soma(first, second int) int {
	return first + second
}

func main() {
	println(soma(5, 5))
}
