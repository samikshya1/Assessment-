package mapper_test

import (
	"fmt"
	"testing"

	"github.com/samikshya/mapper"
)

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {

	cases := []struct {
		input, output string
	}{
		{
			input:  "blabla",
			output: "blAblA",
		},
		{
			input:  "BLaBLa",
			output: "blAblA",
		},
		{
			input:  "Aspiration.com",
			output: "asPirAtiOn.cOm",
		},
		{
			input:  "special#12#123test",
			output: "spEciAl#12#123tEst",
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("test_%s", tc.input), func(t *testing.T) {
			output := mapper.CapitalizeEveryThirdAlphanumericChar(tc.input)
			if tc.output != output {
				t.Errorf("expected %s got %s", tc.output, output)
			}
		})
	}

}
