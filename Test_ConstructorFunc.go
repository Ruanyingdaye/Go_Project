package main

type Executor struct {
	Name string
	Val  int
}

type ConstructorFunc func() *Executor

type Base interface {
	Constructor()
}

func ConstructorA() *Executor {
	println("BaseAImpl")
	return &Executor{Name: "A", Val: 1}
}

func ConstructorB() *Executor {
	println("BaseBImpl")
	return &Executor{Name: "B", Val: 2}
}

func ConstructorC() *Executor {
	println("BaseCImpl")
	return &Executor{Name: "C", Val: 3}
}

func main() {
	Total := 0
	constructorMap := map[string]ConstructorFunc{
		"A": ConstructorA,
		"B": ConstructorB,
		"C": ConstructorC,
	}
	for _, outputFunc := range constructorMap {
		Total += outputFunc().Val
	}

	println(Total)
}
