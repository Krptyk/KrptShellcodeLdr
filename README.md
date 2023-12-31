# DLL Shellcode Loader & Shellcode Encryption

This project demonstrates the creation of a DLL shellcode loader in Golang tailored for Windows environments, aiming to bypass specific security measures, particularly Windows Defender  Moreover, it explores a technique to encrypt shellcode using XOR encryption, making it more elusive to AV static detection.

## UPDATES:
04/01/2024 - Detected and signatured by defender

0## 4/01/2024 - Avoiding Detection by Defender: Code Update Recommendations

To ensure your code does not get flagged by Windows Defender, consider the following updates:

1. **Allocate Memory with Read/Write Permissions**: Start by allocating memory that has read/write access.
2. **Write Shellcode into Allocated Memory**: After memory allocation, write your shellcode into this memory space.
3. **Change Memory Permissions with VirtualProtect API**: Once the shellcode is in place, use the VirtualProtect API to modify the memory permissions to read/execute.
4. **Implement a Timer Check**:
    - **Start Time Comparison**: Add a check to compare the start of the sleep timer to the current system time.
    - **End Time Verification**: After the sleep duration, compare the current system time again to ensure that the full time has passed.
    - **Exit on Fast Forward**: If the full time has not elapsed (which could indicate a scanner fast-forwarding through sleep timers), then exit the program. This step helps in evading scanners that skip through sleep timers.

        

# Full walkthrough
### <a href="https://krptyk.com/2023/09/20/encrypting-shellcode-to-evade-av/" target="_blank">Encrypting Shellcode to Evade AV</a>

### <a href="https://krptyk.com/2023/09/20/dll-shellcode-loader-bypass-defender/" target="_blank">DLL Shellcode Loader in Go</a>


## Disclaimer

The tutorials and code shared through this project are for educational purposes only. Misuse of this information and code for malicious activities is unlawful and unethical. The creators take no responsibility for any misuse.

## Getting Started

### Environment Setup

- Development Environment: Kali Linux
- Target Environment: Windows with AMD64 architecture

Compile command for building the loader:


    CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o bypassdll.dll KrptGoLoader.go


## Dependencies

For the DLL Shellcode Loader:

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

For the XOR Encryption of Shellcode:

    package main
    import (
      "flag"
      "fmt"
      "io/ioutil"
      "os"
    )

## Creating the DLL Shellcode Loader

  Compilation Environment: Use Kali Linux as the development environment. The compile command utilizes the mingw-w64 GCC compiler targeting a Windows OS with an AMD64 architecture. It is built as a shared library with output as bypassdll.dll¹.

  Importing Necessary Packages: Import the necessary packages and libraries to facilitate the loader’s functionalities. The Windows and standard library headers are imported in the C space, while in Go, formats, system calls, time, and unsafe packages are imported to work with low-level memory and system services¹.

## Encrypting the Shellcode

  XOR Encryption for Shellcode Obfuscation: XOR encryption is a simple and effective technique for encrypting shellcode, aiding in the evasion of AV static detection².

  Creating the Go Code: The Go program is designed to take raw binary data as input, perform an XOR operation on it, and produce encrypted shellcode as output².

## Usage

  ### KrptXOR.go
    msfvenom -p windows/x64/exec CMD=calc -f raw 2>/dev/null | go run GoXOR.go -t go -x 31 >> shellcode.txt

  ### KrptLoader.go
    CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o bypassdll.dll bypassdll.go
    C:\>Rundll32.exe .\KrptLoader.dll,ExecuteShellcode
    
