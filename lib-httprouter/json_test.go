package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncodeJson(t *testing.T) {
	LogJson("ugik")
	LogJson(123)
	LogJson(true)
	LogJson([]string{"Ugik", "Fullstack", "Developer"})
}

func TestObjectOne(t *testing.T) {
	var dataObject = ObjectOne{1, "Ugik", "Ini object json"}
	LogJson(dataObject)
}

func TestDecode(t *testing.T) {
	resultJson := `{"Id":1,"Name":"Ugik","Description":"Ini object json"}`
	jsonBytes := []byte(resultJson)

	// mengunakan pointer karna akan dirubah langsung di unmarshal mengunakan memory bukan di return
	object := ObjectOne{}
	json.Unmarshal(jsonBytes, &object)
	fmt.Println(object)
	fmt.Println(object.Id)
	fmt.Println(object.Name)
	fmt.Println(object.Description)
}

type ObjectTwo struct {
	Name    string
	Age     int
	Married bool
	Tags    []string
	Address Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestArrayJson(t *testing.T) {
	objectTwo := ObjectTwo{
		Name:    "Ugik Dev",
		Age:     28,
		Married: true,
		Tags:    []string{"Gaming", "Music", "Coding"},
		Address: Address{"Totokarto", "Indonesia", "33215"},
	}
	LogJson(objectTwo)
}

func TestDecodeObjectJson(t *testing.T) {
	responseJson := `{"Name":"Ugik Dev","Age":28,"Married":true,"Tags":["Gaming","Music","Coding"],"Address":{"Street":"Totokarto","Country":"Indonesia","PostalCode":"33215"}}`
	jsonBytes := []byte(responseJson)

	// mengunakan pointer karna akan dirubah langsung di unmarshal mengunakan memory bukan di return
	object := ObjectTwo{}
	json.Unmarshal(jsonBytes, &object)
	fmt.Println(object)
	fmt.Println(object.Name)
	fmt.Println(object.Age)
	fmt.Println(object.Married)
	fmt.Println(object.Tags)
	fmt.Println(object.Address.Street)
	fmt.Println(object.Address.Country)
	fmt.Println(object.Address.PostalCode)
}

func TestDecodeArray(t *testing.T) {
	responseJson := `[{"Street":"Totokarto","Country":"Indonesia","PostalCode":"33215"},{"Street":"Totokarto","Country":"Indonesia","PostalCode":"33215"},{"Street":"Totokarto","Country":"Indonesia","PostalCode":"33215"}]`
	jsonBytes := []byte(responseJson)

	address := &[]Address{}
	json.Unmarshal(jsonBytes, address)

	for index, addres := range *address {
		fmt.Println(index, addres)
	}

	// testing kembalikan array ke respon json array
	responArrayJson, _ := json.Marshal(address)
	fmt.Println("Balikan lagi ", string(responArrayJson))
}

type Product struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestCaseSensitife(t *testing.T) {
	// snake_case = variable huruf kecil semua dengan menambahkan tag
	product := Product{"Fazea", "exaample.com"}
	responseResult, _ := json.Marshal(product)
	fmt.Println("Menjadi Snake Case ", string(responseResult))

	// ketika membaca dari jsonstring tidak msalah huruf besar kecil
	jsonProduct := `{"name":"Fazea","image_url":"exaample.com"}`
	productBalikan := &Product{}
	json.Unmarshal([]byte(jsonProduct), productBalikan)
	fmt.Println("Balikan dari snake case ", productBalikan)

	// pada image url tidak terbaca
	jsonProduct2 := `{"nAme":"Fazea","ImageUrl":"exaample.com"}`
	productBalikan2 := &Product{}
	json.Unmarshal([]byte(jsonProduct2), productBalikan2)
	fmt.Println("Balikan dari snake case 2", productBalikan2)

}

func TestMapJson(t *testing.T) {

	responseJson := `{"NAme":"Ugik Dev","Age":28,"Married":true,"Tags":["Gaming","Music","Coding"],"Address":{"Street":"Totokarto","Country":"Indonesia","PostalCode":"33215"}}`
	byteJson := []byte(responseJson)

	// jika mengunakan map bisa saja, tapi tidak terhubung lagi dengan kontrak struct
	var result map[string]interface{}
	json.Unmarshal(byteJson, &result)
	fmt.Println("Result", result)

	dataMap := map[string]interface{}{
		"Name":    "Yuvita Tri",
		"Age":     26,
		"Hobbies": []string{"Masak", "Livestreming"},
	}

	bytesMap, _ := json.Marshal(dataMap)

	fmt.Println("Respon Data Map", string(bytesMap))
}

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("example.json")
	decoder := json.NewDecoder(reader)
	objectTwo := &ObjectTwo{}
	decoder.Decode(objectTwo)
	fmt.Println("Data Example.json ", objectTwo)

	// testing ke2

	// reader2, _ := os.Open("openapi.yml")
	reader2, _ := os.Open("example.json")
	decoder2 := json.NewDecoder(reader2)
	var dataMap map[string]interface{}

	decoder2.Decode(&dataMap)
	fmt.Println("Data MAP ", dataMap)
}

func TestStreamEncoder(t *testing.T) {
	// responseJson := `{"NAme":"Ugik Dev","Age":28,"Married":true,"Tags":["Gaming","Music","Coding"],"Address":{"Street":"Totokarto","Country":"Indonesia","PostalCode":"33215"}}`

	dataMap := map[string]interface{}{
		"Name":    "Yuvita Tri Rejeki",
		"Age":     26,
		"Hobbies": []string{"Masak", "Livestreming"},
	}

	write, _ := os.Create("example_result.json")
	decoder := json.NewEncoder(write)
	decoder.Encode(dataMap)

}
