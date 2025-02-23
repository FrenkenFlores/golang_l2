package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

func decompile(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	tmp := []interface{}{}
	// Move from left to right and escape the values.
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' && i+1 < len(str) {
			if unicode.IsDigit(rune(str[i+1])) {
				tmp = append(tmp, string(str[i+1]))
			} else {
				tmp = append(tmp, "\\")
			}
			i++
		} else if unicode.IsDigit(rune(str[i])) {
			n, _ := strconv.Atoi(string(str[i]))
			tmp = append(tmp, n)
		} else {
			tmp = append(tmp, string(str[i]))
		}
	}
	s := ""
	// Move from right to left and decompile the string.
	for i := len(tmp) - 1; i >= 0; i-- {
		repeat := 1
		if reflect.ValueOf(tmp[i]).Kind() == reflect.Int {
			if tmp[i].(int) > 0 {
				repeat = tmp[i].(int)
			}
			i--
		}
		for j := 0; j < repeat; j++ {
			if tmp_s, ok := tmp[i].(string); ok {
				s += tmp_s
			}
		}
	}
	return_value := ""
	// Reverse the string.
	for i := len(s) - 1; i >= 0; i-- {
		return_value += string(s[i])
	}
	if len(return_value) == 0 {
		return "", errors.ErrUnsupported
	}
	return return_value, nil
}

func main() {
	result, err := decompile(`Hello0, world!`)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(result)
	}
}
