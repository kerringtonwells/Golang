package main

import (
        "fmt"
        "os/exec"
)

func main() {
        out, _ := exec.Command("sh", "-c", "ls -lhtra"").Output()
        output := string(out)
        fmt.Println(output)
}
