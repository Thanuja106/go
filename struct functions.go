package main

import "fmt"
type ex struct{
  name string
}
func main() {
cls:=ex{"john"}
cls.test()
test1(cls)
}
func (e ex)test(){
  fmt.Println(e.name)//method on struct type
}
func test1(e ex){
  //Function that takes struct type as the parameter
fmt.Println(e.name)
}
/-------------------------------------------------------------------/
output:
john
john
