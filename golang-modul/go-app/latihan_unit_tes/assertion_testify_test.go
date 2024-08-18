/**
go get github.com/stretchr/testify
go get github.com/stretchr/testify/assert@v1.9.0
go get github.com/stretchr/testify/require

assert jika gagal akan tetap berjalan
require jika gagal akan berhenti
*/

package latihan_unit_tes

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/assert"
)

// cek subtest
func callingcoba(t *testing.T) {
	result := HelloWorld("Ugik")
	assert.Equal(t, "Hello Ugik!! ini helper", result, "Message dari assertion")
}

func TestHelloAssertion(t *testing.T) {
	result := HelloWorld("Ugik")
	// t = pointer, args2 experasi, args3 parameter result, message tidak wajib
	// assert.Equal(t, "Hello Ugik!! ini helper", result)
	assert.Equal(t, "Hello Ugik!! ini helper", result, "Message dari assertion")

	fmt.Println("TestHelloAssertion finis")

	// masih tetap berjalan seperti Fail()
}

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("Ugik")
	// t = pointer, args2 experasi, args3 parameter result, message tidak wajib
	require.Equal(t, "Hello Ugik!! ini helper", result, "Message dari assertion require")
	fmt.Println("Tidak akan berjalan karna ini require")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unit test diskip, Windows tidak mendukung")
	}

	result := HelloWorld("Ugik")
	require.Equal(t, "Hello Ugik!! ini helper", result, "Message dari assertion require")
}
