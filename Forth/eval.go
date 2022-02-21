package main

import ( 
   "fmt"
   "strings"
   "strconv"
   
   )

type Word struct {

  Name      string
  Message   string
  Function  func()
  Id        []float64

}

type Eval struct {

 Stack       Stack 
 
 Dictionary  []Word
 Math        []Math 
  
 tmp         Word
 doOpen      int
 ifOffset    int
 compiling   bool
}

type Math struct {
  Name string
  Function  func()
}

func NewEval() *Eval {

 e := &Eval{}
 
 e.Math = []Math {
        {Name: "+", Function: e.add},                        
        {Name: "-", Function: e.sub},                        
        {Name: "*", Function: e.mul},                        
        {Name: "/", Function: e.div},    
 }
 e.Dictionary = []Word {
        {Name: ".", Function: e.print},  
        {Name: ".s", Function: e.printStack},
        {Name: ":", Function: e.startDefinition},
        {Name: "cls",  Function: e.сlearStack},// чистит стек полностью
        {Name: "dup",  Function: e.dup},//возможно перенсти в Math
        {Name: "show", Function: e.ShowDic},//выводит список слов
        {Name: "emit", Function: e.emit},         
        {Name: "do",   Function: e.do},            
        {Name: "loop", Function: e.loop},
        {Name: "exit", Function: exit},
         
 }
   return e
}

func (e *Eval) Eval(str []string) {
     ////////////////////////////////
   
     for _, tok := range str {

         tok = strings.TrimSpace(tok)
                 
         if tok == "" {
             continue
         }
         
         if tok == ":" {
              e.startDefinition()  
              continue
         }
        
         if e.compiling {
                                    
            if e.tmp.Name == "" {
                             
                id := e.findWord(tok)
                
                if id != -1 {
                   fmt.Printf("word %s defined :", tok)
                   return
                }
                
                e.tmp.Name = tok
                continue
            }
          //////////Проверяет на сообщение и записывает его
          // Возможно добавить или изменить на ."
            if tok == ".'" {
                             
                s := strings.Join(str, " ")
              
                i := strings.Index(s, ".'")
                    fmt.Println("index: ",i)
                            
                ib := strings.Index(s, ";")
                    fmt.Println("index: ",ib)
                
                a := s[(i+2):(ib)]
                    
                e.tmp.Message = a
                 
                    fmt.Println("Str Word :", e.tmp.Message)
            }
              
            if tok == ";"  {
               
               e.Dictionary = append(e.Dictionary, e.tmp)
               e.tmp.Name = ""
               e.tmp.Message  = ""
               e.tmp.Id = []float64{}
               e.compiling = false
               continue 
            }
            
           
          
            id := e.findWord(tok)

            if id >= 0 {
              
                e.tmp.Id = append(e.tmp.Id, float64(id))
                    
                if tok == "do" {
                 
                    e.doOpen = len(e.tmp.Id) - 1
                }
                
                if tok == "loop" {
                 
                    e.tmp.Id = append(e.tmp.Id, -2)
                    e.tmp.Id = append(e.tmp.Id, float64(e.doOpen))
                }
                
                if tok == "if" {
                 
                    e.tmp.Id = append(e.tmp.Id, -3)
                    e.tmp.Id = append(e.tmp.Id, 99)
                    
                    e.ifOffset = len(e.tmp.Id)
                }
                
                if tok == "then" {
                 
                    e.tmp.Id[e.ifOffset-1] = float64(len(e.tmp.Id)-1)
                }
                
                
            } else { // Если нет в словаре ,то ищем в Math.Если нет, то предполагаем что число
              
            /////////////////////////////////
               id_math := e.findMath(tok) 
              
               if id_math >= 0 {
                   
                fmt.Println("in math : ", id_math )
                e.tmp.Id = append(e.tmp.Id, float64(id_math))  
               }
           ////////////////////////////////// 
                
                e.tmp.Id = append(e.tmp.Id, -1)
                val, err := strconv.ParseFloat(tok, 64)
                
                if err != nil {
                   
                   fmt.Printf("Don't number  %s: %s\n", tok, err.Error())
                   continue
                }
                e.tmp.Id = append(e.tmp.Id, val)
            }

            continue
       }
     ////// Конец compiling /////////////////             
         
        handled := false
        
        for id_word, word := range e.Dictionary {
          
            if tok == word.Name {
               
                e.evalWord(id_word)
                handled = true
           }
        }
 ///////////////////////////////////////////
        for id_math, word_math := range e.Math {
        
            if tok == word_math.Name {
               
                    // fmt.Println("in math", tok)
                e.evalMath(id_math)
                handled = true
            }
        }
 
 ///////////////////////////////////////////
        if !handled {
             
          i, err := strconv.ParseFloat(tok, 64)
      
           if err != nil {
            
               fmt.Printf("not in the dictionary or spec.symbol: %s, error :%s\n", tok, err.Error())
               continue
           }
       
           e.Stack.Push(i)
         }
    }
}
///////////// Функции поиска ///////////////
func (e *Eval) findWord(name string) int {
    
    for index, word := range e.Dictionary {
    
        if name == word.Name {
           return index   
        }
           
    }
    return -1
}

func (e *Eval) findMath(name string) int {
    
    for index, word := range e.Math {
    
        if name == word.Name {
           return index   
        }
           
    }
    return -1
}

//////////// Функции вычисления /////////////
func (e *Eval) evalMath(index int) {
    word_math := e.Math[index]
  
  if word_math.Function != nil {
     
        word_math.Function()    
        return  
   } 
}

func (e *Eval) evalWord(index int) {
    
   word := e.Dictionary[index]
    
   if word.Function != nil {
     
       word.Function()    
       return  
   } 
       addNum := false
       
       jump := false
       condjump := false
       inst := 0
         
         for inst < len(word.Id) {
         
         opcode := word.Id[inst]
            
             if addNum {
               
                 e.Stack.Push(opcode)
                 addNum = false
             } else if jump {
                
                cur := e.Stack.Pop()
                max := e.Stack.Pop()
                
                if max > cur {
                 
                    e.Stack.Push(max)
                    e.Stack.Push(cur)
                    
                    inst = int(opcode)
                    inst --
                }
                
                jump = false
                 
               } else if condjump {
                 
                   val := e.Stack.Pop()
                   if val == 0 {
                    
                       inst = int(opcode)
                       inst -- 
                   }
                   condjump = false
                   
               } else {
                   
                    if opcode == -1 {  
                      
                         addNum = true 
                     } else if opcode == -2 {
                           jump = true
                 
                       } 
                       
                       else if opcode == -3 {
                            condjump = true                           
                           
                       } else {          
                            e.evalWord(int(opcode))
                         }
                  }
                
          inst++
       }
}
