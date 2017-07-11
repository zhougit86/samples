package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName string
}

type PersonSlice struct {
	Person []*Person
}

func (p *PersonSlice) Len() int {
	return len(p.Person)
}

func (p *PersonSlice) Less(i,j int) bool {
	return p.Person[i].LastName<p.Person[j].LastName
}

func (p *PersonSlice) Swap(i,j int) {
	p.Person[i],p.Person[j]=p.Person[j],p.Person[i]
}

func main()  {
	Tom := Person{"William","Tom"}
	Jerry := Person{"William","Jerry"}
	pList:=&PersonSlice{[]*Person{&Tom,&Jerry}}
	sort.Sort(pList)
	fmt.Println(pList)

}

