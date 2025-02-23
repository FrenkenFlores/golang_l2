package main

import "testing"

func TestDecompile(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		fail   bool
	}{
		{"valid string with digits", "a4bc2d5e", "aaaabccddddde", false},
		{"valid string without digits", "abcd", "abcd", false},
		{"invalid string starting with digit", "45", "", true},
		{"empty string", "", "", false},
		{"escaped digits", `qwe\4\5`, "qwe45", false},
		{"escaped digit followed by digit", `qwe\45`, "qwe44444", false},
		{"escaped backslash followed by digit", `qwe\\5`, `qwe\\\\\`, false},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := decompile(testCase.input)
			if err != nil && testCase.fail == false {
				t.Errorf("%s: unexpected error status: %v. Input: %s, output: %s", testCase.name, err, testCase.input, result)
			}
			if result != testCase.output {
				t.Errorf("%s: expected %q, got %q", testCase.name, testCase.output, result)
			}
		})
	}
}
