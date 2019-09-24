package main

import "fmt"

func main() {
a:=[3]int{1,2,3}
change(&a)
fmt.Println(a)
}
func change(p *[3]int){
  //*p[0]!=(*p)[0]
  (*p)[0]=2
   (*p)[1]=3
    (*p)[2]=4                                                     output:[2 3 4]

}
