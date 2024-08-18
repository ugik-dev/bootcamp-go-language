package latihangoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	// channel <- "Zea"
	// parsing data dari chanel

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ikan"
		channel <- "Ayam"
		fmt.Println("Mulai berburu")
	}()

	data2 := <-channel
	fmt.Println("Hasil tangkapan1 = ", data2)
	data := <-channel
	fmt.Println("Hasil tangkapan2 = ", data)

	time.Sleep(5 * time.Second)

}

func TestTakeResponse(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go takeResponse(channel)

	data2 := <-channel
	fmt.Println("Hasil tangkapan1 = ", data2)
	data := <-channel
	fmt.Println("Hasil tangkapan2 = ", data)

	time.Sleep(5 * time.Second)
}

func takeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Mulai berburu")
	channel <- "Kucing"
	channel <- "Anjing"
}

// In Out Channel
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go onlyOutOrReceiver(channel)
	go onlyInSending(channel)
	time.Sleep(5 * time.Second) //untuk delay
}

func onlyInSending(channel chan<- string) {
	fmt.Println("onlyInSending()")
	time.Sleep(2 * time.Second) //delay
	channel <- "Kucing"
}

func onlyOutOrReceiver(channel <-chan string) {
	fmt.Println("onlyOutOrReceiver()")
	data := <-channel
	fmt.Println("Hasil onlyOutOrReceiver = ", data)
}

// Buffer Channel
func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 4)
	//args ke 2= buffer
	fmt.Println("Channel dengan capasitas ", cap(channel), "dengan panjang", len(channel))
	defer close(channel)
	// fmt.Println("Load ke 1", <-channel)
	// channel <- "Kapasitas ke 1"
	// fmt.Println("panjang", len(channel))
	// channel <- "Kapasitas ke 2"
	// fmt.Println("panjang", len(channel))
	// channel <- "Kapasitas ke 3"
	// fmt.Println("panjang", len(channel))
	// channel <- "Kapasitas ke 4"
	// fmt.Println("panjang", len(channel))
	// fmt.Println("Load ke 2", <-channel)
	// fmt.Println("Load ke 3", <-channel)
	// fmt.Println("Load ke 4", <-channel)
	go onlyOutOrReceiverBuffer(channel)
	go onlyInSendingBuffer(channel)
	fmt.Println("Load ke dari luar", <-channel)
	fmt.Println("Finish TestBufferChannel")
	time.Sleep(5 * time.Second) //untuk delay
}

func onlyInSendingBuffer(channel chan<- string) {
	fmt.Println("onlyInSendingBuffer()")
	time.Sleep(2 * time.Second) //delay
	channel <- "Kapasitas ke 1"
	fmt.Println("panjang", len(channel))
	channel <- "Kapasitas ke 2"
	fmt.Println("panjang", len(channel))
	channel <- "Kapasitas ke 3"
	fmt.Println("panjang", len(channel))
	channel <- "Kapasitas ke 4"
	fmt.Println("panjang", len(channel))
}

func onlyOutOrReceiverBuffer(channel <-chan string) {
	fmt.Println("onlyOutOrReceiverBuffer()")
	// fmt.Println("Load ke 1", <-channel)
	// fmt.Println("Load ke 2", <-channel)
	// fmt.Println("Load ke 3", <-channel)
	for data := range channel {
		fmt.Println("Load ke 3", data)
	}
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	fmt.Println("Channel dengan capasitas ", cap(channel), "dengan panjang", len(channel))
	// defer close(channel)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Looping ke " + strconv.Itoa(i)
		}
		close(channel)
	}()
	// go onlyOutOrReceiverBuffer(channel)
	for data := range channel {
		fmt.Println("Load ke 3", data)
	}
	fmt.Println("Finish TestRangeChannel")
	time.Sleep(5 * time.Second) //untuk delay
}

func TestSelectChannel(t *testing.T) {
	channel := make(chan string)
	channel2 := make(chan string)
	fmt.Println("Channel dengan capasitas ", cap(channel), "dengan panjang", len(channel))
	defer close(channel)
	defer close(channel2)

	go takeResponse(channel)
	go takeResponse(channel2)

	// jika tidak mengunakan default for {}, jika tidak ada default bisa gunakan for i i< 4
	// Kenapa 4? karna di takeResponse terdapat 2kali assign data ke channel dan takerespon dipanggil 2x
	// for i := 0; i < 4; i++ {
	counter := 0
	for {

		select {
		case data := <-channel:
			fmt.Println("Data dari chanel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari chanel 2", data)
			counter++
		default:
			fmt.Println("-- Loading data --")
		}
		if counter == 4 {
			break
		}
	}
	// }
	fmt.Println("Finish SelectChannel")
}
