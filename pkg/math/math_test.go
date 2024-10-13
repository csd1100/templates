package math

import "testing"

func TestAdd(t *testing.T) {
	num1 := 10
	num2 := 13
	expected := 23

	result := Add(num1, num2)

	if result != expected {
		t.Errorf("math.Add(%d, %d) should be equals %d, not %d", num1, num2, expected, result)
	}
}
