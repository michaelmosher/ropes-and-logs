package statemachine

import (
	"fmt"
	"testing"
)

func TestSerialize(t *testing.T) {
	var tests = []struct {
		k    string
		v    int
		want string
	}{
		{"a", 0, "Set a to 0."},
		{"key with space", 10, "Set key with space to 10."},
		{"fly me to the moon", 100, "Set fly me to the moon to 100."},
		{"Set to", 30, "Set Set to to 30."},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Serialize: %s, %d", tt.k, tt.v)
		t.Run(testname, func(t *testing.T) {
			c := SetCommand{tt.k, tt.v}
			got := c.Serialize()

			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestDeserialize(t *testing.T) {
	var tests = []struct {
		input string
		want  SetCommand
	}{
		{"Set a to 0.", SetCommand{"a", 0}},
		{"Set key with space to 10.", SetCommand{"key with space", 10}},
		{"Set fly me to the moon to 100.", SetCommand{"fly me to the moon", 100}},
		{"Set Set to to 30.", SetCommand{"Set to", 30}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Deserialize: %s", tt.input)
		t.Run(testname, func(t *testing.T) {
			got := SetCommand{}
			got.Deserialize(tt.input)

			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}
