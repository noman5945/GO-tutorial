package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

//Struct types (used to define a set of data of various type)
type gasEngine struct{
	mpg uint
	gallon uint
}

type electricEngine struct{
	mpkwh uint
	kwh uint
}

//method for each struct(these methods extend these structs from OOP pov)
func (e gasEngine) milesLeft()uint{
	return e.gallon*e.mpg
}

func (e electricEngine) milesLeft()uint{
	return e.kwh*e.mpkwh
}
// to make a genralized method that takes any struct type we use interface
type engine interface{
	milesLeft() uint
}

func canMakeIt(e engine,miles uint){
	if miles<=e.milesLeft(){
		fmt.Println("Can make it!")
	}else{
		fmt.Println("Naah need gas/fuel/charge yo")
	}
}

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

	fmt.Println("Type Structs")
	var myEngine gasEngine=gasEngine{mpg: 25,gallon: 20}
	var anotherEngine electricEngine=electricEngine{mpkwh: 25,kwh: 20}
	fmt.Println(myEngine.mpg,myEngine.gallon)
	fmt.Printf("Miles left gas Engine (calling milesleft method of struct):%v\n",myEngine.milesLeft())
	canMakeIt(myEngine,45)
	canMakeIt(anotherEngine,45)
	
	fmt.Println("Pointers")
	var thing1=[5]float32{1,2,3,4,5}
	fmt.Printf("\nMemory location of thing1 array is : %p",&thing1)
	var result_pointer [5]float32=square_with_pointers(&thing1)
	fmt.Printf("\n Result: %v",result_pointer)
	fmt.Printf("\nValue of thing1 array is : %v",thing1)

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

//usage of pointers
func square_with_pointers(thing2 *[5]float32)[5]float32{
	fmt.Printf("\nMemory location of thing2 array is : %p",thing2)
	for i:=range thing2{
		thing2[i]=thing2[i]*thing2[i]
	}
	return *thing2
}