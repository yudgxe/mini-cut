package cut

import (
	"bufio"
	"io"
	"strings"

	"github.com/yudgxe/cut/value"
)

func Read(r io.Reader, field value.MasValue, delimiter string, separated bool) (string, error) {
	scanner := bufio.NewScanner(r)

	var sb strings.Builder

	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), delimiter)

		if len(strs) == 1 {
			if !separated {
				sb.WriteString(scanner.Text() + "\n")
			}

			continue
		}

		for _, v := range field.Values {
			if v < len(strs) {
				sb.WriteString(strs[v] + " ")
			}
		}

		sb.WriteRune('\n')
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return sb.String(), nil
}
