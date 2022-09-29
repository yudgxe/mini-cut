package value_test

import (
	"testing"

	"github.com/yudgxe/cut/value"
)

func TestSet(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{
			input:    "1,2,3,4",
			expected: []int{0, 1, 2, 3},
		},
		{
			input:    "2,2",
			expected: []int{1},
		},
		{
			input:    "3,2",
			expected: []int{1, 2},
		},
		{
			input:    "4,7,2,4",
			expected: []int{1, 3, 6},
		},
	}

	for _, test := range tests {
		m := new(value.MasValue)
		m.Set(test.input)

		if len(m.Values) != len(test.expected) {
			t.Errorf("Incorrect len, Expect %d, got %d", len(test.expected), len(m.Values))
		}

		for i, v := range m.Values {
			if v != test.expected[i] {
				t.Errorf("Incorrect result, Expect %d, got %d", test.expected[i], v)
			}
		}
	}
}

func TestString(t *testing.T) {
	m := new(value.MasValue)
	m.String()

	if m.Values[0] != 1 {
		t.Errorf("Дефолтное значение не работает")
	}

	expected := "1,2,3,4"

	m.Values = []int{1, 2, 3, 4}
	result := m.String()

	if result != expected {
		t.Errorf("Incorrect result, Expect %s, got %s", expected, result)
	}
}
