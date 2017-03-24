
     var replacer = strings.NewReplacer(" ", "\", \"")
     str := "ls -lhtr"
     str = replacer.Replace(str)
