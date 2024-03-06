package main

import (
	"fmt"
	// "sort"
	// "strings"
)

func main(){
	//string
	// var name1 string = "Dezy";
	// var name2 string = "Chris";
	// name3:= "Baddest boy"
	// name3 = "Godwin"
	// var num4 int= 25
	// fmt.Println(name1, name2, name3, num4)
	// fmt.Print(name1 + " ")
	// fmt.Println(name1)

	//arrays
// 	var name [4]string = [4]string{"Desmond", "Brian","Javis", "Randolf"};
//    name [0] = "Paul"
// 	fmt.Println(name)

// 	//Slices
// 	var count = []int{1,2,3,4,5};
// 	count [1] = 3
// 	var newCount = append(count, 9)

// 	fmt.Println(count, newCount)







	var nameOf [4]string = [4]string{"desmond", "randolf","Brian", "Agogo"}
	nameOf[0] = "Immanuel"
	// family:= "My name is desmond and from Kumba"

	// fmt.Print(strings.Contains(family, "ond"));
	// fmt.Print(strings.Replace(family, "desmond", "Brian", 1));
	// fmt.Print(strings.Index(family, "ond"));
	// fmt.Print(strings.Split(family, ","));

	// age:=[]int{43, 45, 97, 20, 12, 90,48}
	// sort.Ints(age);
	// fmt.Println(age);
	// index := sort.SearchInts(age, 12)
	// fmt.Println(index)

	//Loop

	ourName := []string{"desmond", "randolf","Brian", "Agogo"}
// for i:=0; i<len(ourName); i++ {
// 	fmt.Println("hello, your name is", ourName[i])
// }


//.......
// for index, value := range ourName{
// 	fmt.Println("hello, the value at index %v is %v", index, value)
// }

//.......
//  for _, value := range ourName{
//  	fmt.Printf("hello, my name is %v \n", value)
//  }

  // Break and Contnue;
  for index, value:= range ourName{
	if index== 1 {
		fmt.Println("Continue at pos", index)
		continue
	}
	if index == 2{
		fmt.Println("Break at this pos", index)
		break
	}
	fmt.Printf("th value at pos %v is %v \n", index, value)
	
  };
}

// func seeLoops(){

// }
