package main

import "fmt"
type city struct {
  greeting string 
  name string
  run float32
}
func main(){
  cities:=[]city{
    city{
      greeting:"hai",
      name:"a",
      run:12.90,
    },
    city{
      greeting:"bye",
      name:"b",
      run:98.00,
    },
    }
    for _,x:=range cities{
      fmt.Printf("%s %s %f \n",x.greeting,x.name,x.run)

    }
  }
