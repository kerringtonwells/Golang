package main

import (
        "fmt"
        "os/exec"
)

func main() {
        out, _ := exec.Command("sh", "-c", "date +\"%Y-%m-%d %H:%M:%S %Z\"").Output()
        fmt.Printf("%s", out)
}
