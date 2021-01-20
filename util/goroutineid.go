package Util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func GetCurrentGoroutineId() int64 {
	var (
		buf [64]byte
		n = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)
	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}
	return int64(id)
}

// 建议使用此接口，少一次 Atoi 的转换，反正最终打印也是%s
func GetCurrentGoroutineIdStr() string {
	var (
		buf [64]byte
		n = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)
	idField := strings.Fields(stk)[0]
	return idField
}
