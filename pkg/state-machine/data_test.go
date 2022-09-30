package statemachine

import (
	"fmt"
	"testing"
)

func TestDataApplyOnce(t *testing.T) {
	var tests = []struct {
		c    SetCommand
		k    string
		want int
	}{
		{SetCommand{"foo", 1}, "foo", 1},
		{SetCommand{"hello", 5}, "hello", 5},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("ApplyOnce: %+v", tt.c)
		t.Run(testname, func(t *testing.T) {
			data := New()
			data.Apply(tt.c)
			got := data.Get(tt.k)

			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestDataApplyMulitple(t *testing.T) {
	type result struct {
		key   string
		value int
	}

	tests := []struct {
		cs     []SetCommand
		checks []result
	}{
		{
			[]SetCommand{{"a", 1}, {"b", 2}, {"c", 3}},
			[]result{{"a", 1}, {"b", 2}, {"c", 3}},
		},
		{
			[]SetCommand{{"a", 1}, {"b", 2}, {"a", 3}},
			[]result{{"a", 3}, {"b", 2}},
		},
	}

	for _, tt := range tests {
		data := New()

		for _, c := range tt.cs {
			data.Apply(c)
		}

		testname := fmt.Sprintf("ApplyMulitple: %+v", tt.cs)

		t.Run(testname, func(t *testing.T) {
			for _, r := range tt.checks {
				got := data.Get(r.key)
				want := r.value

				if got != want {
					t.Errorf("got %+v, want %+v", got, want)
				}
			}
		})
	}
}
