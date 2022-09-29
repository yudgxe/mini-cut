package cut

import (
	"strings"
	"testing"

	"github.com/yudgxe/cut/value"
)

var (
	test1 string = `245:789	4567	|M:4540	Admin	01:10:1980
535:763	4987	|M:3476	Sales	11:04:1978`
)

func TestCut(t *testing.T) {
	tests := []struct {
		data      string
		field     []int
		delimiter string
		separated bool
		expected  string
	}{
		{
			data:      test1,
			field:     []int{0},
			delimiter: "\t",
			expected:  "245:789 \n535:763 ",
		},
		{
			data:      test1,
			field:     []int{1},
			delimiter: "\t",
			expected:  "4567 \n4987 ",
		},
	}

	m := value.MasValue{}
	for _, test := range tests {
		reader := strings.NewReader(test.data)
		m.Values = test.field

		result, err := Read(reader, m, test.delimiter, test.separated)
		if err != nil {
			t.Errorf(err.Error())
		}

		if result != test.expected {
			t.Errorf("\nIncorrect result, Expect\n%s\n, got\n%s", test.expected, result)
		}
	}
}
