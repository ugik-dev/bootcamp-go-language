package latihan_unit_tes

/**
Unit Test
nama file harus diahiri _test.go
untuk menjalankan unit test

go test
go test -v (untuk detail )
go test -v -run=NamaFungsi (single test)

go test harus masuk kedalam dir seperti cd helper dlu
go test ./... (untuk test dari root project)

Fail() => gagal dan di lanjutkan
FailNow => gagal dan prosess terhenti
Error() => menampilkan info dan Fail / lanjut
Fatal() => menampilkan info dan FailNow() / terhenti


Benchmark

unit test akan ikut runing
go test -v -bench=.

tampa unit test
go test -v -run=FungsiYangTidakAdaAtauSalah -bench=.
*/
import (
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		HelloWorld("Eko")
	}
}
func BenchmarkHelloB(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		HelloWorld("Sugi Pramana")
	}
}
func TestHelloWord(t *testing.T) {

	result := HelloWorld("Eko")
	if result != "Hello Eko!! ini helperFAIL" {
		t.Fail()
		// panic("result tidak sesuai")
	}

	// fmt.Println("runing TestHelloWord finish")
}

func TestHelloWordS(t *testing.T) {

	result := HelloWorld("Ugik")
	if result != "Hello Ugik!! ini helper" {
		// panic("result tidak sesuai")
		// t.FailNow()
	}
	// fmt.Println("runing TestHelloWordS finish")

}

func TestError(t *testing.T) {

	result := HelloWorld("Ugik")
	if result != "Hello Ugik!! ini helperFAIL" {
		t.Error("Result tidak sesuai")
	}
}

func TestFatal(t *testing.T) {
	// Fatal akan terhenti seperti FailNow
	result := HelloWorld("Ugik")
	if result != "Hello Ugik!! ini helperFAIL" {
		// t.Fatal("Result tidak sesuai")
	}
}
