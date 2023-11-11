package utils

import (
	"fmt"
	"strings"
)

func FormatHex(data []byte) string {
	var builder strings.Builder
	builder.Grow(len(data) * 4)
	builder.WriteString(fmt.Sprintf("\r\n"))
	count := 0
	pass := 1
	for _, b := range data {
		if count == 0 {
			builder.WriteString(fmt.Sprintf("%-6s\t", "["+fmt.Sprint((pass-1)*16)+"]"))
		}

		count++
		builder.WriteString(fmt.Sprintf("%02X", b))
		if count == 4 || count == 8 || count == 12 {
			builder.WriteString(" ")
		}
		if count == 16 {
			builder.WriteString("\t")
			for i := (pass * count) - 16; i < (pass * count); i++ {
				c := rune(data[i])
				if c > 0x1f && c < 0x80 {
					builder.WriteRune(c)
				} else {
					builder.WriteString(".")
				}
			}
			builder.WriteString("\r\n")
			count = 0
			pass++
		}
	}
	builder.WriteString("\r\n")
	builder.WriteString("\r\n")
	return builder.String()
}
