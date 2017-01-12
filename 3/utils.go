package main

import (
	"fmt"
	"regexp"
)

var commandPattern = regexp.MustCompile(`^(?:[:](\S+) )?(\S+)(?: ([^:].+?))?(?: [:](.+))?$`)

func ParseCommand(line string) (*Command, error) {
	fmt.Printf("Line: %v ---\n", line)
	if commandPattern.MatchString(line) {
		matches := commandPattern.FindStringSubmatch(line)
		fmt.Printf("Matches: %v\n", matches)

		return &Command{
			Source:  matches[1],
			Target:  matches[3],
			Command: matches[2],
			Params:  matches[4]}, nil
	}
	return nil, fmt.Errorf("Failed to parse command from [%s]", line)
}
