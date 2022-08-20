package mapper

import (
	"fmt"
	"regexp"
)

var (
	alnum = regexp.MustCompile("[[:alnum:]]")
)

// Interface exposes two functions
type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

// MapString will capitalize given input
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

// SkipString type to implement Interface interface
type SkipString struct {
	str        string
	desiredPos int
	posCount   int
}

// NewSkipString to create a SkipString struct
func NewSkipString(pos int, str string) SkipString {
	return SkipString{desiredPos: pos, str: str}
}

// TransformRune capitalizes the character at given the position
func (ss *SkipString) TransformRune(pos int) {
	runeArray := []rune(ss.str)
	runeArray[pos], ss.posCount = transform(runeArray[pos], ss.posCount, ss.desiredPos)
	ss.str = string(runeArray)
}

// GetValueAsRuneSlice retunrs rune arry for a string
func (ss *SkipString) GetValueAsRuneSlice() []rune {
	return []rune(ss.str)
}

func (ss SkipString) String() string {
	return fmt.Sprintf("%v", ss.str)
}

// CapitalizeEveryThirdAlphanumericChar function to capitalize every third character
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	runeArray := []rune(s)
	desiredPos := 3
	posCount := 0

	for n := range runeArray {
		runeArray[n], posCount = transform(runeArray[n], posCount, desiredPos)
	}
	return string(runeArray)
}

func transform(input rune, posCount int, desiredPos int) (rune, int) {
	if alnum.MatchString(string(input)) {
		if posCount == desiredPos-1 { //For every 3rd char
			if input >= 97 && input <= 122 { // Small letters
				input = input - 32
			}
			posCount = 0
		} else {
			if input >= 65 && input <= 90 { // Capital letters
				input = input + 32
			}
			posCount++
		}
	}
	return input, posCount
}
