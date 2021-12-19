package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

var lock = sync.RWMutex{}

type MyJson struct {
	MyId    string
	Xml     string
	MyPDF   map[string]PDF `json:"my_pdf"`
	RefLink []string       `json:"ref_link"`
}

type PDF struct {
	PDFname      string `json:"pdf_name"`
	EncodedValue string `json:"encoded_value"`
}

func addElementsToPDF(x string, y string, mappdf map[string]PDF) {

	w, c := func(x string, y string) (string, string) {

		a := base64.StdEncoding.EncodeToString([]byte(x))
		b := base64.StdEncoding.EncodeToString([]byte(y))

		return a, b
	}(x, y)

	lock.Lock()
	defer lock.Unlock()
	pdfino := PDF{"test1", w}
	pdfino2 := PDF{"test2", c}

	mappdf["web1"] = pdfino
	mappdf["web2"] = pdfino2
}

func main() {

	var reflink []string
	var pdftest = make(map[string]PDF)

	plainStr := "Hello World."
	secondEncodedString := "Saluti da Roma"

	//	encodedStr, encodedStr2 := encodefiles(plainStr, secondEncodedString)

	addElementsToPDF(plainStr, secondEncodedString, pdftest)

	reflink = append(reflink, "b", "c")

	r := MyJson{MyId: "Myid", Xml: "myxml", MyPDF: pdftest, RefLink: reflink}

	p, err := json.MarshalIndent(&r, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", p)
}
