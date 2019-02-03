package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name      string     `xml:"name"`
	Mobile    string     `xml:"mobile"`
	Locations []Location `xml:"location"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	var person Person
	xmlString := getDummyXml()
	xml.Unmarshal([]byte(xmlString), &person)
	fmt.Println(person)
	//fmt.Println("name : ", person.Name, ", mobile : ", person.Mobile, ", location : ", person.Locations)
	for _, Location := range person.Locations {
		fmt.Printf("\n%s", Location)
	}
}

func getDummyXml() string {
	return "<person><name>praveen</name><mobile>7977726385</mobile><location><loc>thane</loc></location><location><loc>mulund</loc></location></person>"
}
