package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	NoValue = "<NA>"
)

func Atoa(v *string) string {
	if v == nil {
		return NoValue
	}
	return *v
}

func Itoa(v *int) string {
	if v == nil {
		return NoValue
	}
	return strconv.Itoa(*v)
}

func Ftoa(v *float32) string {
	if v == nil {
		return NoValue
	}
	return strconv.FormatFloat(float64(*v), 'f', -1, 64)
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

func FormatName(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

func PrintId(format string, id int) {
	if format == "json" {
		fmt.Printf("{\"id\":\"%d\"}\n", id)
	} else {
		fmt.Println(id)
	}
}
