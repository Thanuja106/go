package main

import "fmt"
type vertex struct{
 x int
 y int
}
func main() {
v:=vertex{1,2}
p:=&v
 fmt.Println(p)
p.x=23//we can write this *p.X or p.x
fmt.Println(v)
}
OUTPUT:
&{23,2}
{23,2}
