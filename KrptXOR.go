/*
Author: Krptyk
Description: XOR Encryption for Shellcode Obfuscation
Credits: https://github.com/yoda66/GoShellcode - ported from Python to Golang
*/
 
package main
 
import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
)
 
func xorme(buf []byte, k byte) []byte {
    res := make([]byte, len(buf))
    for i, ch := range buf {
        res[i] = ch ^ k
    }
    return res
}
 
func printOutput(buf []byte, t string) {
    var fmtStr string
    var end string
 
    switch t {
    case "go":
        fmtStr = "0x%02x"
        end = ","
        fmt.Println("buf := []byte{")
    case "c#":
        fmtStr = "0x%02x"
        end = ","
        fmt.Println("byte[] buf = {")
    case "py":
        fmtStr = "\\x%02x"
        end = ""
        fmt.Println("buf := []byte{")
    }
 
    for i, ch := range buf[:len(buf)-1] {
        if i > 0 && i%16 == 0 {
            if t == "py" {
                fmt.Print("\\")
            }
            fmt.Println()
        }
        fmt.Printf(fmtStr, ch)
        fmt.Print(end)
    }
 
    fmt.Printf(fmtStr, buf[len(buf)-1])
 
    switch t {
    case "go":
        fmt.Println("}")
    case "c#":
        fmt.Println("};")
    case "py":
        fmt.Println("}")
    }
}
 
func main() {
    t := flag.String("t", "py", "Language type to generate for")
    x := flag.Int("x", 55, "XOR integer (defaults to 55)")
    flag.Parse()
 
    buf, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
        os.Exit(1)
    }
 
    xorbuf := xorme(buf, byte(*x))
    printOutput(xorbuf, *t)
}
