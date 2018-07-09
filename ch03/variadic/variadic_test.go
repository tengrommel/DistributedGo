package variadic

import "testing"

type AddArray struct {
	result  int;
	add_1   int;
	add_2   int;
}

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

func BenchmarkLoops(b *testing.B)  {
	var test ForTest
	ptr := &test
	// 必须循环b.N次。这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i:=0;i<b.N;i++ {
		ptr.Loops()
	}
}

func BenchmarkLoopsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var test ForTest
		ptr := &test
		for pb.Next() {
			ptr.Loops()
		}
	})
}

func TestBench(t *testing.T) {
	var test_data = [3] AddArray {
		{ 2, 1, 1},
		{ 5, 2, 3},
		{ 4, 2, 2},
	}
	for _ , v := range test_data {
		if v.result != Add(v.add_1, v.add_2) {
			t.Errorf("Add( %d , %d ) != %d \n", v.add_1 , v.add_2 , v.result);
		}
	}
}