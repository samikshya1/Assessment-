package main

import (
	"fmt"

	"github.com/samikshya/mapper"
)

func NewSkipString(skipCount int, s string) mapper.Mapper {
	return mapper.NewMapper(skipCount, s)
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)
}
