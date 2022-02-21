package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
 

 
)

func main() {

   // Должен написать сообщение подключен/не подключен монитор( желательно в цвете )

    reader := bufio.NewReader(os.Stdin)
    forth := NewEval()
    
    for {
      fmt.Printf("> ")

      text, err := reader.ReadString('\n')
    
         
      if err != nil {
        fmt.Printf("Fehler beim Lesen des Eingangs : %s\n", err.Error())
        return
       }
   
      text = strings.TrimSpace(text)
      forth.Eval(strings.Split(text, " "))
          
     }
    
}




