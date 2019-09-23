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


/-------------------------------------------------------------------------------------------/
package main//another way of struct in arrays

import "fmt"

type ex struct{
  name string
  rank int
}
func main() {
  var a [] ex
  cls1:=ex{"aa",1}
  cls2:=ex{"bb",2}
  a=append(a,cls1)
  a=append(a,cls2)
  fmt.Println(a[0].name,a[0].rank,"\n",a[1])

}
/-----------------------------------------------------------------------------------------------/
var opts = []struct {//another way
    shortnm      byte
    longnm, help string
    needArg      bool
}{
    {'a', "multiple", "Usage for a", false},
    {
        shortnm: 'b',
        longnm:  "b-option",
        needArg: false,
        help:    "Usage for b",
    },
}


