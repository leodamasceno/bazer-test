package main

import (
        "fmt"
    "os/exec"
    "bytes"
)

func main() {
    cmd := exec.Command("terraform", "validate", "-no-color")
var out bytes.Buffer
var stderr bytes.Buffer
cmd.Stdout = &out
cmd.Stderr = &stderr
err := cmd.Run()
if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    return
}
fmt.Println("Result: " + out.String())
}
