package variadic

func simpleVariadicToSlice(numbers ...int) []int {
	return numbers
}

func mixedVariadicToSlice(name string, numbers ...int) (string, []int) {
	return name, numbers
}

func Add(x int, y int) (z int) {
	z = x + y
	return
}

type ForTest struct {
	num int
}

func (this * ForTest) Loops() {
	for i:=0;i <10000;i++ {
		this.num++
	}
}