package main

import "fmt"

func main() {
sum(1,2)
sum(1,2,3)
a:=[]int{3,4,5,67,7}
sum(a...)
}
func sum(aa...int){//varadic functions
  tot:=0
  fmt.Print(aa," ")
  for _,a:=range aa{
  tot=tot+a

  }
    fmt.Println(tot)
}
OUTPUT:
[1 2] 3
[1 2 3] 6
[3 4 5 67 7] 86
/-----------------------------------------------------------------------------
package main
import "fmt"
func main() {
  add(1,2,3,4)
  }
func add(num...int){
  t:=0
  for _,n:=range num{
    t=t+n
}
fmt.Println(t)
}
OUTPUT
10
