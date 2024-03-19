package main

import (
	"fmt"
  "github.com/gin-gonic/gin"
)
type bill struct{
	name string;
	items map[string]float64;
	tip float64;
}
 
 func newBill (name string) bill{
	b:=bill{
		name: name,
		items: map[string]float64{},

	}
	return b
 }
func main(){
	name :=newBill("Desmond")

	fmt.Println(name)

}


