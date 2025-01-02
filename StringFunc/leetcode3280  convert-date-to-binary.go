package StringFunc

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertDateToBinary(date string) string {
	parts := strings.Split(date, "-")
	for i, p := range parts {
		n, _ := strconv.Atoi(p)
		parts[i] = fmt.Sprintf("%b", n)
	}
	return strings.Join(parts, "-")
}
