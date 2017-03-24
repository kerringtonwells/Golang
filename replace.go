 package main
 
 import (
     "fmt"
     "strings"
     "os/exec"
 )
 
func main() {
  var replacer = strings.NewReplacer(" ", "\", \"")
  str := "ls -lhtr"
  str = replacer.Replace(str)
}
