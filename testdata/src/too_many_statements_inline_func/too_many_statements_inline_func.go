package main

func main() { // want `Function 'main' has too many statements \(8 > 1\)`
	print("Hello, world!")
	if true {
		y := []int{1, 2, 3, 4}
		for k, v := range y {
			f := func() { print("test", k, v) }
			f()
		}
	}
	print("Hello, world!")
}
