program1

package main  
import (  
 "fmt"  
)  
  
func CheckPrime(number int) {  
 isPrime := true  
 if number == 0 || number == 1 {  
  fmt.Printf(" %d is not a  prime number\n", number)  
 } else {  
  for i := 2; i <= number/2; i++ {  
   if number%i == 0 {  
    fmt.Printf(" %d is not a  prime number\n", number)  
    isPrime = false  
    break  
   }  
  }  
  if isPrime == true {  
   fmt.Printf(" %d is a prime number\n", number)  
  }  
 }  
}  
func main() {  
 CheckPrime(0)  // not a prime  
 CheckPrime(1)  // not a prime  
 CheckPrime(3)  // a prime  
 CheckPrime(25) // not a prime  
  
}  

output:
0 is not a  prime number
 1 is not a  prime number
 3 is a prime number
 25 is not a  prime number

Program exited.