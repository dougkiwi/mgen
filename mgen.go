package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("monster.xml")
	if err != nil {
		log.Fatal("file open fail")
	}
	decoder := xml.NewDecoder(file)
	for {
		//fmt.Print(". ")
		t, _ := decoder.Token()
		//log.Print(err)
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "monster" {
				fmt.Println("found one")
			}
		}
	}
	fmt.Printf("hello, world\n")
}
