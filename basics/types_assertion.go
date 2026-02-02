package main

func main() {
	var x any = 42
	if v, ok := x.(int); ok {
		println(v)
	} else {
		println("not ok")
	}
}
