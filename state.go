package state

import (
	"errors"
	"fmt"
)

type char int

const (
	HS char = iota
	kylling
	korn
	rev
	numchar
)

var west string = "HS rev korn kyllying"
var boat string = ""
var east string = ""

func ViewState() string {
	return "[Kylling rev korn hs ---\\ \\__/ _________________/---]"
}

func Put(item string) string {
	return "[kylling rev korn ---\\ \\_korn_/ _________________/---]"
}

type state
func (c char) eats(e char, s state) bool {
	if c == rev && e == kylling || c == kylling && e == korn {
		if s[c] == s[e] && s[HS] != s[c] {
			return true
		}
	}

	return false
}

func CrossRiverTo(place string) error {
	if place == "west" {
		west = west + boat
		boat = ""
	} else if place == "east" {
		east = east + boat
		boat = ""
	} else {
		return errors.New(fmt.Sprintf("'%s' is an invalid place", place))
	}
	return nil
}
