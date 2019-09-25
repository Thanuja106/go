package main

import "fmt"
type str struct{
  name string
  email string
}
func main() {
s:=str{"john","john@gmail.com"}
fmt.Println(s)
}
func (ss str) String() string{
  return fmt.Sprintf("%s <%s>",ss.name,ss.email)
}
output:

john <john@gmail.com>//actually it has to come like {john john@gmail.com}
