package main

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time `protobuf:"-"`
}

// ToUnstructured implements the value.UnstructuredConverter interface.
func (t Time) ToUnstructured() interface{} {
	if t.IsZero() {
		return nil
	}
	buf := make([]byte, 0, len(time.RFC3339))
	buf = t.UTC().AppendFormat(buf, time.RFC3339)
	return string(buf)
}

func (t Time) WithLocalToUnstructured() interface{} {
	if t.IsZero() {
		return nil
	}
	buf := make([]byte, 0, len(time.RFC3339))
	buf = t.Local().AppendFormat(buf, time.RFC3339)
	return string(buf)
}

func tz1() {
	// 获取当前时间
	t := time.Now()

	// 获取当前时区
	loc := t.Location()

	// 打印时区信息
	fmt.Println("当前时区:", loc)
}
func main() {
	tz1()
	t1 := Time{time.Now()}
	fmt.Println(t1.ToUnstructured())
	fmt.Println(t1.WithLocalToUnstructured())
	// time.Sleep(24 * 8 * time.Minute)
}
