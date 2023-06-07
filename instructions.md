Q) how to know (confirm) the linux system version ? 
A) An the command prompt send the command: 
ssdm2 @ ssdm2-X99-F8 ~/github.com/guidorafael/taxas (main)
└$ hostnamectl
 Static hostname: ssdm2-X99-F8
       Icon name: computer-desktop
         Chassis: desktop
      Machine ID: 26d423dc96cf49f6abd4ddb1c86260bc
         Boot ID: 5dcad0dacf5b4fb597df88f1cb647524
Operating System: Ubuntu 22.04 LTS                
          Kernel: Linux 5.19.0-43-generic
    Architecture: x86-64
 Hardware Vendor: HUANANZHI
  Hardware Model: X99-F8
ssdm2 @ ssdm2-X99-F8 ~/github.com/guidorafael/taxas (main)
└$ 

// Tutorial //  
https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04

How To Build Go Executables for Multiple Platforms on Ubuntu 20.04
Published on March 9, 2022 · Updated on March 10, 2022.
(acesso em 1º de junho de 2023)

By Madison Scott-Clary and Marko Mudrinić

Step 1 — Creating a simple Go program
/////////////////////////////////////
└$ mkdir hello
└$ cd hello

When importing packages, you have to manage the dependencies through the code’s own module. 
You can do this by creating a go.mod file with the go mod init command:

└$ go mod init hello         (e ele "esta" na pasta hello)

Next, create a Hello, World! Go file in your preferred text editor:
Add the following text into your hello.go file:

package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

Test your code to check that it prints the Hello, World! greeting:
└$ go run .       // só tem este na pasta ... // only have "hello.go" in the directory

Output:
Hello, World!

Step 2 — Building an Executable
///////////////////////////////
The go run command ran the code for your “Hello, World!” program, 
but you may want to build your program into a binary to run anywhere 
on your system, not just from the source. The go build command builds executables.

└$ go build hello

