package latihan_unit_tes

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {

	//before excecution
	fmt.Println("Main test is starting")
	m.Run()
	fmt.Println("Main test is finish")
	//after excecution
}
