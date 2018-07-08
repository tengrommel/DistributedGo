package variadic

import "testing"

func TestSimpleVariadicToSlice(t *testing.T)  {
	// Test for no arguments
	if val := simpleVariadicToSlice(); val!=nil{
		t.Error("value should be nil", nil)
	} else {
		t.Log("simpleVariadicToSlic() -> nil")
	}
	// Test for random set of values
	vals := simpleVariadicToSlice(1,2,3)
	expected := []int{1,2,3}
	isErr := false
	for i:=0;i<3;i++ {
		if vals[i] != expected[i]{
			isErr = true
			break
		}
	}
	if isErr{
		t.Error("value should be []int{1,2,3}", vals)
	} else {
		t.Log("simpleVariadicToSlic(1,2,3) -> []int{1,2,3}")
	}
}

func TestMixedVariadicToSlice(t *testing.T) {
	// Test for simple argument & no variadic arguments
	name, numbers := mixedVariadicToSlice("Bob")
	if name == "Bob" && numbers == nil {
		t.Log("Recieved as expected: Bob, <nil slice>")
	} else {
		t.Errorf("Received unexpected values: %s, %s", name, numbers)
	}
}