package type_helper

import (
	"fmt"
	"log/slog"
	"strconv"
)

func String(n int32) string {
	slog.Debug(fmt.Sprint("Converting", n, "from int32 to String"))
	buf := [11]byte{}
	pos := len(buf)
	i, q := int64(n), int64(0)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		q = i / 10
		buf[pos], i = '0'+byte(i-10*q), q
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func Int64Parser(str string) int64 {
	ret, err := strconv.ParseInt(str, 10, 64)
	fmt.Errorf("Error parsing time value", err)

	return ret
}
