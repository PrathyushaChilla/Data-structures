package main

import(
 "fmt"
)
func main() {
m := make(map[string]int)
m["Orange"] = 100
m["Straberry"] = 150
m["Mango"] = 300
for key, val := range m {
fmt.Print("[ ",key," -> ", val," ]")
}
fmt.Println("Orange price:", m["Orange"])
delete(m, "Orange")
fmt.Println("orange price:", m["Orange"])
v, ok := m["Orange"]
fmt.Println("orange price:", v, "Present:", ok)
v2, ok2 := m["Straberry"]
fmt.Println("Straberry price:", v2, "Present:", ok2)
}

output:
[ Mango -> 300 ][ Orange -> 100 ][ Straberry -> 150 ]Orange price: 100
orange price: 0
orange price: 0 Present: false
Straberry price: 150 Present: true

Program exited.