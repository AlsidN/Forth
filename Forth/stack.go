package main

import (
   "fmt"
)

type Stack []float64

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

func (s *Stack) CheckStack() bool {
    
 return len(*s) < 2  
}

func (s *Stack) Push(x float64) {
   *s = append(*s, x)
   
}

func (s *Stack) ShowStack()  {
  
      for  n := 0; n <= len(*s)-1; n++ {
       
        q := (*s)[n]
        fmt.Printf("%v ", q)
      }
  
   fmt.Println("<-Top ok")
       
}

func (s *Stack) Pop() float64 {

    i := len(*s) - 1
    x := (*s)[i]
    *s = (*s)[:i]

   return x
}

func (s *Stack) ClearStack() {
  
   i := (len(*s) - 1)- (len(*s) - 1)
  *s = (*s)[:i]

  fmt.Println("Clear ok, stack is empty")///////////////  
    
}
