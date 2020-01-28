package main

import (
	"strings"
	"testing"
)

var versions string = "1.0.0\n1.0.1\n1.1.0\n2.0.0\n1.1.2\n2.0.1\n1.2.0\n3.0.0\n2.1.0\n2.2.0"

func TestVersionsMajor(t *testing.T) {

	t.Parallel() // marks TLog as capable of running in parallel with other tests
	tests := []struct {
		in  string
		out string
	}{
		{"3.0.1", "3"},
		{"2.2.1", "2"},
		{"2.3.0", "2"},
		{"1.1.1", ""},
		{"1.1.3", ""},
		{"1.2.1", "1"},
		{"2.1.1", ""},
		{"4.0.0", "4"},
	}
	for _, tt := range tests {
		tt := tt // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel() // marks each test case as capable of running in parallel with each other

			out := major(run(strings.NewReader(versions), tt.in))
			if out != tt.out {
				t.Errorf("got %v, want %v", out, tt.out)
			}
		})
	}
}

func TestVersionsMinor(t *testing.T) {

	t.Parallel() // marks TLog as capable of running in parallel with other tests
	tests := []struct {
		in  string
		out string
	}{
		{"3.0.1", "3.0"},
		{"2.2.1", "2.2"},
		{"2.3.0", "2.3"},
		{"1.1.1", ""},
		{"1.1.3", "1.1"},
		{"1.2.1", "1.2"},
		{"2.1.1", "2.1"},
		{"4.0.0", "4.0"},
	}
	for _, tt := range tests {
		tt := tt // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel() // marks each test case as capable of running in parallel with each other

			out := minor(run(strings.NewReader(versions), tt.in))
			if out != tt.out {
				t.Errorf("got %v, want %v", out, tt.out)
			}
		})
	}
}
