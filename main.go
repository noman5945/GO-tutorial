package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main(){
	//Data types
	fmt.Println("Hello go")
	var intNum int16=32767;
	intNum=intNum+1;
	fmt.Println(intNum)

	var newNum int=64;
	var secNum int=20;
	var result int=newNum%secNum
	fmt.Println(result)

	fmt.Println(len("GGWP dhds")) // len gives the length of bits in the string. Problamatic for special chars
	fmt.Println(utf8.RuneCountInString("GGWP dhds")) // len gives the actual length of the string.

	//rune typo
	var myRune rune='a'
	fmt.Println(myRune)
	
	myVar:="text0"
	fmt.Println(myVar)

	var1,var2:=1,2
	fmt.Println(var1,var2)
	
	//constants
	const pi float32=3.1416
	fmt.Println(pi)
	
	//functions
	var result2,reminder,err =intDivision(5,0)
	if err!=nil{
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("Result and Reminder are: %v and %v\n",result2,reminder)
	//Data Structures
	fmt.Println("Arrays")
	arrayDataStruct()
	
	fmt.Println("Slice")
	sliceDataStruct()
	
	fmt.Println("Maps")
	mapDataStruct()
}

// functions
func intDivision(numerator int,denominator int)(int ,int ,error){
	var err error
	if denominator==0{
		err=errors.New("cannot divide by zero")
		return 0,0,err
	}
	var result int= numerator/denominator
	var reminder int = numerator%denominator
	return result,reminder,err
}

// Arrays
func arrayDataStruct(){
	var intArr [3]int32
	intArr[1]=123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
}

//Slices
func sliceDataStruct(){
	var intSlice []int32=[]int32{8,9}
	fmt.Println(intSlice)
	
	var intSlice2 []int32=[]int32{10,15}
	intSlice=append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32=make([]int32, 3,10)
	fmt.Println(intSlice3)

}

//Maps
func mapDataStruct(){
	//This variable m is a map of string keys to int values
	var m=make(map[string]int) ;
	m["route"]=66;
	fmt.Println(m["route"]);
	//This statement retrieves the value stored under the key "route" and assigns it to a new variable i:
	var i int =m["route"];
	fmt.Println(i)
	//To initialize a map with some data, use a map literal:
	var commits = map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	//iterate the map
	for key,value:=range commits{
		fmt.Println("key:",key ,"value:",value)
	}
	//Initiate empty Map
	var emptyMap=map[string]int{};
	fmt.Println(emptyMap)
}