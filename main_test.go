package main

import (
	"testing"
)

func TestGetDataList(t *testing.T) {
	var v []string

	v = GetDataList(src)
	if len(v) < 1 {
		t.Error("файлов для обработки нет")
	}
}
