package main

import "fmt"

func main() {
a:=[3]int{1,2,3}
change(&a)
fmt.Println(a)
}
func change(p *[3]int){
  //*p[0]!=(*p)[0]
  (*p)[0]=2||p[0]=2
   (*p)[1]=3||p[1]=3
    (*p)[2]=4||p[2]=4 are the same                                                     output:[2 3 4]

}

/--------------------------------------------------------------------------------------------------------------------
package main

import "fmt"

func main() {
a:=[6]int{1,2,3,4,5,6}
var p [4]*int
for i:=0;i<len(p);i++{
  p[i]=&a[i]
}
for i:=0;i<len(p);i++{
  fmt.Println(*p[i],"\n")
}
}

OUTPUT:
1
2
3
4
