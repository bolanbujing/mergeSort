package main
import (
  "fmt"
  "math"
)

func mergeTwoList(left []int, right []int, cl chan int) {
  fmt.Printf("before left:%v\n",left)
  fmt.Printf("before right:%v\n",right)
  i:=0
  j:=0
  ch := make([]int, 0)
  for i < len(left) && j < len(right) {
    if left[i] < right[j] {
      ch = append(ch, left[i])
      i++
    }else {
      ch = append(ch, right[j])
      j++;
    }
  }
  if i < len(left) {
    ch = append(ch, left[i:len(left)]...)
  }
  if j < len(right) {
    ch = append(ch, right[j:len(right)]...)
  }
  fmt.Printf("ch:%v\n",ch)
  left = append(left[:0], ch[0:len(left)]...)
  fmt.Printf("after left:%v\n",left)
  right = append(right[:0], ch[len(left):len(ch)]...)
  fmt.Printf("after right:%v\n",right)
  cl<-1
}

func mergeSort(slice []int, length int, rge int, m int) {
  fmt.Printf("rge:%v , %d\n", rge, m)
  cl := make(chan int)
  i:=0
  left:=i
  right:=left+rge
  for left < length {
    if right >= length {
      return
    } else if right+rge > length {
       go mergeTwoList(slice[left:right], slice[right:length], cl)
    } else {
       go mergeTwoList(slice[left:right], slice[right:right+rge], cl)
    }
    left = right+rge
    right = left+rge
  }
  n:=m
  for n>0 {
    <-cl  
    fmt.Printf("n:%d\n", n)
    n--
  }
}

func main() {
  //slice := []int{23,5,-12,6,23,7889,2,6,22,1}
  slice := []int{23,5,-12,6,23,7889,2,6,22,1,4}
  fmt.Println(slice)
  var i float64 = 0
  rge:=1
  k := 0
  m := 0
  for rge < len(slice) {
    k = len(slice)%(2*rge)
    m = len(slice)/(2*rge)
    if k <= rge {
      mergeSort(slice, len(slice), rge, m)
    } else {
      mergeSort(slice, len(slice), rge, m+1)
    }
    i++
    rge = int(math.Pow(2,i)) 
    fmt.Println(slice)
  }
  fmt.Printf("%v\n",slice)
}

