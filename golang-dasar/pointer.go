package main

import "fmt"

type Address struct {
	Name         string
	streetNumber string
	active       bool
}

// saat akan merubah active mengunakan methode

func (address *Address) nonActive() {
	address.active = false
	tmpname := &address.Name
	*tmpname = "gugu"
	fmt.Println(*tmpname)
	fmt.Println(address.Name)
}
func main() {

	slice := []string{"Jeruk", "Mangga"}
	copyBiasa := slice
	copyBiasa[0] = "Anggur"
	fmt.Println(slice)
	fmt.Println(copyBiasa)

	address1 := Address{"Alamat 1", "30a", true}
	address2 := address1
	address3 := &address1 // ini mengunakan pointer sehingga nilai disimpan dimemory yg sama
	address2.Name = "Alamat Baru"
	address3.Name = "Alamat Ter Baru"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// membuat address 3 menjadi pointer baru tampa merubah parent
	// address3 = &Address{"Alamat via asteris", "30a"}

	// semua parent dari pointer ikut berubah
	*address3 = Address{"Alamat via asteris", "30a", true}
	fmt.Println("-- resign pointer baru --")
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// changeAlamat(&address1)
	// changeAlamat(&address2)
	changeAlamat(address3) //jika sudah diinisialisasi di awal bahwa ini pointer maka tidak perlu lagi tanda&
	fmt.Println("-- resign pointer dari fungsi --")
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	/**
	pada saat method biasa parameter yang digunakan diobject dicopy tidak mengunakan memory yang sama
	sehingga harus mengunakan * dan & pada inisialisasi fungsi
	*/
	fmt.Println("-- from method --")
	address3.nonActive()
	fmt.Println(address3)
}

func changeAlamat(address *Address) {
	address.Name = "Alamat dari fungsi"
}
