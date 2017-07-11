package main

import (
	"fmt"
	"errors"
	"reflect"
	"strings"
)

type Qiexi interface {
	shout()
}

type Person struct {
	firstName string "name Tag"
}

func (P Person) shout(){
	fmt.Println(P.firstName)
}

const (
	a = iota * 50
	b
	c = `This is a raw string \n`
)



func main() {
	A:=Person{"Jason"}
	var B Qiexi
	B=A
	B.shout()
	if t,ok:=B.(Person);ok{
		fmt.Printf("%T\n%T\n",t,B)
	}

	defer func() {
		if someth:=recover();someth!=nil{
			//fmt.Println("haha",someth)
			if strings.Index(someth.(string),"new")>-1{
				fmt.Println("this has a New")
			}
			fmt.Println(reflect.TypeOf(someth))
		}
	}()

	CusErr := errors.New("I am a new error")
	panic("the Error is: " + CusErr.Error())
	fmt.Println(reflect.TypeOf(CusErr))
	fmt.Printf("%T\n",CusErr)

	TestSli:=make([]int,2)
	TestSli[2]=5

	fmt.Println("this part is ok")

	//C:='A'
	//ty:=reflect.ValueOf(&C)
	//conten:=ty.Elem()
	//fmt.Println(ty.CanSet(),conten.CanSet())
	//
	//fmt.Println(conten)

}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		case rune:
			fmt.Printf("Param #%d is a Rune\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}



