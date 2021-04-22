package mapper

import (
	"unicode"
)

// CapitalizeEveryThirdAlphanumericChar capitalizes *only* every third alphanumeric character of a string,
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// calculate valid index as per description
	// capitalizes *only* every third alphanumeric character of a string
	var result = []rune(s)
	var validIndex int
	for i, v := range result {
		if unicode.IsLetter(v) {
			validIndex = validIndex + 1
		} else {
			continue
		}
		if validIndex%3 == 0 {
			if unicode.IsLower(v) {
				result[i] = (unicode.ToUpper(v))
				//s =  string(s[i]=) //strings.Replace(s, string(v), string(unicode.ToUpper(v)), 1)
			}
		} else {
			if unicode.IsUpper(v) {
				result[i] = unicode.ToLower(v)
				//s = //strings.Replace(s, string(v), string(unicode.ToLower(v)), 1)
			}
		}
	}
	return string(result)
}

type Mapper struct {
	SkipCount int
	Data      string
	r         []rune
}

func NewMapper(skipCount int, data string) Mapper {
	return Mapper{
		Data:      data,
		r:         []rune(data),
		SkipCount: skipCount,
	}
}

func (m *Mapper) TransformRune(pos int) {
	var validIndex int
	for i, v := range m.r[:pos] {
		if unicode.IsLetter(v) {
			validIndex = validIndex + 1
		} else {
			continue
		}
		if validIndex%m.SkipCount == 0 {
			if unicode.IsLower(v) {
				m.r[i] = (unicode.ToUpper(v))
			}
		} else {
			if unicode.IsUpper(v) {
				m.r[i] = unicode.ToLower(v)
			}
		}
	}
}
func (m *Mapper) GetValueAsRuneSlice() []rune {
	return m.r
}

func (m Mapper) String() string {
	return string(m.r)
}
