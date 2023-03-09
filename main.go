package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"golang.org/x/crypto/bcrypt"
)
type person struct{
	First string
	Last string
	Age int
}
type ByAge []person

func (a ByAge)Len() int{ return len(a)}
func (a ByAge)Swap(i,j int) {a[i],a[j] = a[j],a[i]}
func (a ByAge)Less(i,j int) bool {return a[j].Age>a[i].Age /* TODO: Sort By Name||Age*/ }



func main() {


	bcript()
}
//TODO: JSON//
func json_(){
	s:= `[{"First":"Johan","Last":"Chacon","Age":32},{"First":"Fle","Last":"Bass", "Age":51}]`
	bs := []byte(s)

	p:=[]person{}
	err := json.Unmarshal(bs,&p)	
	if err!=nil{
		fmt.Println("This is the error: ",err)
	}
	fmt.Println(p)

	for i, val := range p{
		fmt.Println("This is the Index:",i)
		fmt.Println("This is the value:",val)
	}
}
// TODO: Sort// 
func sort_By_Name(){
	p1 := person{
		First:"Person1",
		Last:"Lp1",
		Age: 23,
	}
	p2 := person{
		First: "Person2",
		Last: "Lp2",
		Age:98,
	}
	p3:= person{
		First: "Person3",
		Last: "Lp3",
		Age:44,
	}
	people := []person{
		p1,p2,p3,
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
func _sort_(){
	xi := []int{9,1,2,3,41,5,1,412,5,6,}
	// xs := []string{"Johan","Josue","Fernando","David","Pedro"}
	fmt.Println(xi)
	sort.Slice(xi, func(i, j int) bool {
		return xi[i]<xi[j]
	})
	fmt.Println(xi)
}
func sort_int(){
	xi:= []int{9,1,2,3,41,5,1,412,5,6,}
	//order slice 
	sort.Ints(xi)
	fmt.Println(xi)
	//reverse slice
	sort.Sort(sort.Reverse(sort.IntSlice(xi)))
	fmt.Println(xi)
}

///////////////TODO: BCRIPT//////////////////////
func  bcript(){
	s:= "Password1234"
	fmt.Println("This is the original Password", s)
	bs, err := bcrypt.GenerateFromPassword( []byte(s),	bcrypt.MinCost)
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(bs)
	//The password is Bcrypt
	p:= "Password123"

	err= bcrypt.CompareHashAndPassword([]byte(bs),[]byte(p))
	if err!= nil{
		fmt.Println("The password is not the same:",err)
	}else{
		fmt.Println("You are loged in")
	}



}
