package latihan_unit_tes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

type DesignTest struct {
	name    string
	request string
	expect  string
}

func TestTable(t *testing.T) {
	tests := []DesignTest{
		{"Ugik", "Ugik", "Hello Ugik"},
		{"Vita", "Vita", "Hello Vita"},
		{"Zea", "Zea", "Hello Zea"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HayHay(test.request)
			assert.Equal(t, test.expect, result, "Message table message")
		})
	}

}
func TestSubTes(t *testing.T) {
	//go test -run=TestSubTest (run all)
	//go test -run=TestSubTest/Ugik (run subnya)
	t.Run("Ugik", func(t *testing.T) {
		result := HelloWorld("Ugik")
		assert.Equal(t, "Hello Ugik!! ini helper", result, "Message dari assertion")
	})
	t.Run("cb", callingcoba)
}
