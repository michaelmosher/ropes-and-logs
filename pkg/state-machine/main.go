package statemachine

import (
	"fmt"
	"regexp"
	"strconv"
)

type SetCommand struct {
	Key   string
	Value int
}

func (s *SetCommand) Serialize() string {
	return fmt.Sprintf("Set %s to %d.", s.Key, s.Value)
}

func (s *SetCommand) Deserialize(input string) error {
	re := regexp.MustCompile(`Set (?P<key>[a-zA-Z ]+) to (?P<value>[0-9]+).`)

	if !re.MatchString(input) {
		return fmt.Errorf("input does not resemble a command")
	}

	matches := re.FindStringSubmatch(input)

	v, err := strconv.Atoi(matches[re.SubexpIndex("value")])
	if err != nil {
		return fmt.Errorf("value in command does not resemble an integer")
	}

	s.Key = matches[re.SubexpIndex("key")]
	s.Value = v

	return nil
}

type data struct {
	kv map[string]int
}

func New() (d data) {
	d.kv = make(map[string]int)
	return
}

func (d *data) Apply(s SetCommand) {
	d.kv[s.Key] = s.Value
}

func (d *data) Get(k string) int {
	return d.kv[k]
}
