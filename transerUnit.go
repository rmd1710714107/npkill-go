package main

import (
	"fmt"
	"strconv"
)

type transferRes struct {
	val  string
	unit string
}

func transferUnit(originSize int64) string {
	temSize := float64(originSize)
	var TransferRes transferRes
	switch {
	case temSize < (1 << 10):
		TransferRes.val = strconv.FormatFloat(temSize/(1<<10), 'f', 2, 64)
		TransferRes.unit = "B"

	case temSize >= (1>>10) && temSize < (1<<20):
		TransferRes.val = strconv.FormatFloat(temSize/(1<<10), 'f', 2, 64)
		TransferRes.unit = "KB"

	case temSize >= (1<<20) && temSize < (1<<30):
		TransferRes.val = strconv.FormatFloat(temSize/(1<<20), 'f', 2, 64)
		TransferRes.unit = "MB"
	case temSize >= (1<<30) && temSize < (1<<40):
		TransferRes.val = strconv.FormatFloat(temSize/(1<<30), 'f', 2, 64)
		TransferRes.unit = "GB"
	default:
		TransferRes.val = strconv.FormatFloat(temSize/(1<<40), 'f', 2, 64)
		TransferRes.unit = "TB"
	}
	return fmt.Sprintf("%s%s", TransferRes.val, TransferRes.unit)
}
