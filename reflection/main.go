package main

import (
	"fmt"
	"reflect"
)

//Sometime we need to know the type of a value or variable at runtime
//Example we want to pass different type of structs to a method e.g order and employee
//Two different structs order and employee
type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

//func to return the type of interface and the value of the interface
/*
func createQuery(q interface{}) {
	//concrete type of interface passed as an argument = main.order
	t := reflect.TypeOf(q)
	//value of the interface passed as argument = 456 56
	v := reflect.ValueOf(q)
	//specific kind of the type of interface passed as value = struct
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)

	//reflect.Numfield() gets number of fields in a struct
	//reflect.Field(i int) gets refect.Value of ith field

	//The methods Int and String help extract the reflect.
	//Value as an int64 and string respectively.
*/

/*
//example using struct kind, NumField and Field(i)
func createQuery(q interface{}) {
	//check whether the Kind of q is a struct
	//NumField method works only on struct
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
*/

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		//get type of interface in this case struct
		//and get name of that struct in first case order
		t := reflect.TypeOf(q).Name()
		//create insert adding order as table name
		query := fmt.Sprintf("insert into %s values(", t)
		//fmt.Printf("Initial query =  %s \r\n", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			//switch on type of value in struct field int, string..
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
					//fmt.Printf("Query format is: %s \r\n", query)
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
					//if second value add , then value
					//fmt.Printf("Alternate query format is: %s \r\n", query)
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	//handle unsupported types
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)

}
