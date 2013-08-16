package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Monster struct {
	Name string `xml:"name"`
	Descriptor string `xml:"descriptor"`
	Alignment string `xml:"alignment"`
	Challenge_rating string `xml:"challenge_rating"`
}

type Monsters struct {
	Monster []Monster `xml:"monster"`
}


func main() {
	//monsters := make([]Monster,500)
	//monsters := new(Monsters)
	//monsters := new(Monsters)
	monsters := Monsters{make([]Monster,500)}
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
			if se.Name.Local == "monsters" {
				fmt.Println("found one")
				err := decoder.DecodeElement(&monsters, &se)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	//m := new(Monster)
	for  i, m := range monsters.Monster {
		fmt.Printf("%d:\t%s\t%s\t%s\n",i,m.Name,m.Challenge_rating,m.Alignment)
	}
	fmt.Printf("hello, world\n")
}
