for i:=0;i<temp;i++{
  for j:=0;j<temp1;j++{
    fmt.Println("enter marks")
    fmt.Scanln(&a[i])
    sum=sum+a[i]
    }
fmt.Println("the sum is",sum)
avg:=sum/temp1
fmt.Println("the avg is",avg)
k[rr]=sum
rr++
c++
}

//fmt.Println("the rank is ",k)
for r:=0;r<c;r++{
  for r1:=r+1;r1<c;r1++{
  if k[r]<=k[r+1]{
       t=k[r]
      k[r]=k[r+1]
       k[r+1]=t 
}
  }
   }
fmt.Println("the rank ",k)
}
