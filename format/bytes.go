package format

import (
	"fmt"
	"strconv"
)

const (
	Byte = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
	TiByte
	PiByte
	EiByte
)

const (
	IByte = 1
	KByte = IByte * 1000
	MByte = KByte * 1000
	GByte = MByte * 1000
	TByte = GByte * 1000
	PByte = TByte * 1000
	EByte = PByte * 1000
)

//
// Bytes Formats given bytes to string
//
func Bytes(bytes uint64, decimalPlaces byte) string {
	tpl := fmt.Sprintf("%%.%df", decimalPlaces)

	switch {
	case bytes > EByte:
		return fmt.Sprintf(tpl, float64(bytes)/EByte) + " EB"
	case bytes > PByte:
		return fmt.Sprintf(tpl, float64(bytes)/PByte) + " PB"
	case bytes > TByte:
		return fmt.Sprintf(tpl, float64(bytes)/TByte) + " TB"
	case bytes > GByte:
		return fmt.Sprintf(tpl, float64(bytes)/GByte) + " GB"
	case bytes > MByte:
		return fmt.Sprintf(tpl, float64(bytes)/MByte) + " MB"
	case bytes > KByte:
		return fmt.Sprintf(tpl, float64(bytes)/KByte) + " kB"
	default:
		return fmt.Sprintf("%d", bytes) + " Byte"
	}
}

func Number(num float64, decimalPlaces byte) string {
	var ret string
	tmp := fmt.Sprintf("%.0f", num)

	c := 0
	for i := len(tmp) - 1; i >= 0; i-- {
		if c > 0 && c%3 == 0 {
			ret = "," + ret
		}

		ret = string(tmp[i]) + ret

		c++
	}

	if decimalPlaces > 0 {
		ret += fmt.Sprintf(
			"%."+strconv.FormatInt(int64(decimalPlaces), 10)+"f",
			num-float64(uint64(num)),
		)[1:]
	}

	return ret
}
