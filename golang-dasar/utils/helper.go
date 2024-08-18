package utils

/**

Saat deklarasi fungsi variable parameter struct dan interface
penamaan huruf kecil untuk private
huruf besar diawal agar bisa diaksess dari package lain

*/

var connection string

func init() {
	connection = "MongoDB"
}
func SayHello(name string) string {
	return "Hay, " + name + "!!"
}
