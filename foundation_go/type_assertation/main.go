package main

func main() {
	var minhaVar interface{} = "Leonardo"

	res, ok := minhaVar.(int)

	println(res, ok)
}
