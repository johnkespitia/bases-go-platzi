// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

package main

import (
  "fmt"
  "strings"
)

func main() {
  param := "[([]){}]"
  
  //valid2 := "()" //true
  
  
  //invalid := "((]"
  if checkString(param) {
    fmt.Printf("%s Valid",param)
  }else{
    fmt.Printf("%s Invalid",param)
  }
  
}

func checkString(x string) bool {
  splittedString := strings.Split(x, "")
  var validateSlice []string
    for _, char := range splittedString {
    if char == "(" {
      validateSlice = append(validateSlice, char)
    } else if char == "{" {
      validateSlice = append(validateSlice, char)
    } else if char == "[" {
      validateSlice = append(validateSlice, char)
    } else if len(validateSlice) > 0{
      if char == ")" && validateSlice[len(validateSlice)-1] == "(" {
        validateSlice = validateSlice[:len(validateSlice)-1]
      } else if char == "]" && validateSlice[len(validateSlice)-1] == "[" {
        validateSlice = validateSlice[:len(validateSlice)-1]
      } else if char == "}" && validateSlice[len(validateSlice)-1] == "{" {
        validateSlice = validateSlice[:len(validateSlice)-1]
      }else{
        validateSlice = append(validateSlice, char)
      }
    }else{
      validateSlice = append(validateSlice, char)
    }
  }
  return len(validateSlice)==0
}

