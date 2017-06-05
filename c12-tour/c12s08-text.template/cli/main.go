package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name      string
	Age       int
	Team      string
	TeamColor string
}
type Page struct {
	Members []*Person
}

func main() {
	fmt.Println("Using package 'text/tempalte'")

	genWelcomeLetters("welcome")
	genMembershipList()
}
func genMembershipList() {
	t := template.Must(template.ParseFiles("members.html"))

	data := getData()
	page := Page{data}

	var f *os.File
	var err error

	f, err = os.Create("membership.list.html")
	if nil != err {
		log.Fatalln(err)
	}
	t.Execute(f, page)
	f.Close()
}

func genWelcomeLetters(outfile string) {
	t := template.Must(template.ParseFiles("welcome.tmpl"))

	data := getData()

	var f *os.File
	var err error
	var fn string

	for i, p := range data {
		fn = fmt.Sprintf("%s.%d.txt", outfile, i)
		f, err = os.Create(fn)
		if nil != err {
			log.Println(err)
			continue // skip record for which we can't create file
		}
		t.Execute(f, p)
		fmt.Fprintln(f)
		f.Close()
	}
}

func getData() (data []*Person) {
	var p *Person
	p = &Person{"Jane Doe", 21, "A", "RED"}
	data = append(data, p)
	p = &Person{"Mary-Ann Sue", 84, "B", "BLUE"}
	data = append(data, p)
	p = &Person{"Jack Smith", 43, "C", "GREEN"}
	data = append(data, p)
	p = &Person{"Ann Wise", 36, "D", "YELLOW"}
	data = append(data, p)
	p = &Person{"Sam Jones", 41, "A", "RED"}
	data = append(data, p)
	return
}
