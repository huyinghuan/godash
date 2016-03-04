package godash

import "testing"

func TestToInt(t *testing.T) {
	tests := []string{"1000", "-123", "abcdef", "100000000000000000000000000000000000000000000"}
	expected := []int64{1000, -123, 0, 0}
	for i := 0; i < len(tests); i++ {
		result, _ := ToInt(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestToBoolean(t *testing.T) {
	tests := []string{"true", "1", "True", "false", "0", "abcdef"}
	expected := []bool{true, true, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		res, _ := ToBoolean(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func toString(t *testing.T, test interface{}, expected string) {
	res := ToString(test)
	if res != expected {
		t.Log("Case ToString: expected ", expected, " when result is ", res)
		t.FailNow()
	}
}

func TestToString(t *testing.T) {
	toString(t, "str123", "str123")
	toString(t, 123, "123")
	toString(t, 12.3, "12.3")
	toString(t, true, "true")
	toString(t, 1.5+10i, "(1.5+10i)")
	// Sprintf function not guarantee that maps with equal keys always will be equal in string  representation
	//toString(t, struct{ Keys map[int]int }{Keys: map[int]int{1: 2, 3: 4}}, "{map[1:2 3:4]}")
}

func TestToFloat(t *testing.T) {
	tests := []string{"", "123", "-.01", "10.", "string", "1.23e3", ".23e10"}
	expected := []float64{0, 123, -0.01, 10.0, 0, 1230, 0.23e10}
	for i := 0; i < len(tests); i++ {
		res, _ := ToFloat(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func TestUnderscoreToCamelCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected string
	}{
		{"a_b_c", "ABC"},
		{"my_func", "MyFunc"},
		{"1ab_cd", "1abCd"},
	}
	for _, test := range tests {
		actual := UnderscoreToCamelCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected UnderscoreToCamelCase(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestCamelCaseToUnderscore(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected string
	}{
		{"MyFunc", "my_func"},
		{"ABC", "a_b_c"},
		{"1B", "1_b"},
	}
	for _, test := range tests {
		actual := CamelCaseToUnderscore(test.param)
		if actual != test.expected {
			t.Errorf("Expected CamelCaseToUnderscore(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}
