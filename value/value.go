package value

import (
	"sort"
	"strconv"
	"strings"
)

type MasValue struct {
	Values []int
}

func (m *MasValue) String() string {
	if m.Values != nil {
		var sb strings.Builder

		for i := 0; i < len(m.Values)-1; i++ {
			sb.WriteString(strconv.Itoa(m.Values[i]))
			sb.WriteRune(',')
		}

		if len(m.Values)-1 >= 0 {
			sb.WriteString(strconv.Itoa(m.Values[len(m.Values)-1]))
		}

		return sb.String()
	}

	m.Values = append(m.Values, 1)
	return ""

}

func (m *MasValue) Set(s string) error {
	values := strings.Split(s, ",")
	var buff []int

	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		buff = appendIfMissing(buff, num-1)
	}

	sort.Slice(buff, func(i, j int) bool {
		return buff[i] < buff[j]
	})

	m.Values = buff

	return nil
}

func appendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
