package main

import (
    "fmt"
    "os"
)

func (e *Eval) printStack() {
    
    if e.Stack.IsEmpty() {
	   fmt.Printf("0 stack is empty\n")///////////////
       return
   }
    e.Stack.ShowStack()
    
 }

func (e *Eval) add() {
    
    q := e.Stack.CheckStack()
    
    
    if !q {
       a := e.Stack.Pop()
       b := e.Stack.Pop()
       e.Stack.Push(a + b)
    } else {
        fmt.Println("Stack < 2 elements, no add : "  , q)
        
    }
}

func (e *Eval) sub() {
    a := e.Stack.Pop()
    b := e.Stack.Pop()
    e.Stack.Push(b - a)
}

func (e *Eval) mul() {
    a := e.Stack.Pop()
    b := e.Stack.Pop()
    e.Stack.Push(a * b)
}

func (e *Eval) div() {
    a := e.Stack.Pop()
    b := e.Stack.Pop()
    e.Stack.Push(b / a)
}

func (e *Eval) dup() {
    a := e.Stack.Pop()
    e.Stack.Push(a)
    e.Stack.Push(a)
}

func (e *Eval) сlearStack() {
   e.Stack.ClearStack()
} 

func (e *Eval) print() {
       
   if e.Stack.IsEmpty() {
	   fmt.Printf("0 stack is empty\n")
       return
   }
    
    a := e.Stack.Pop()
    fmt.Printf("%f\n", a)
}

func (e *Eval) startDefinition() {
     e.compiling = true   
}

func (e *Eval) ShowDic() {
    fmt.Println("----Facts----")
  
    for i, w := range e.Dictionary {
      
      fmt.Print("id: ", i)
     if w.Message == ""  {
            
        fmt.Println("   name:", w.Name)
        continue
    } else {
        fmt.Print("   name: ", w.Name)
         fmt.Println("   message:", w.Message)    
      }
         
  }
    
}

func (e *Eval) emit() {
   a := e.Stack.Pop()
   fmt.Printf("%c", rune(a))
   
}

func (e *Eval) do() {
    
}

func (e *Eval) loop() {
   cur := e.Stack.Pop()
   max := e.Stack.Pop()
   
   cur++
   
   e.Stack.Push(max)
   e.Stack.Push(cur)
    
}

func exit() {
 os.Exit(0)
    
}
