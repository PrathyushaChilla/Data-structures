program1

package main 
   
import ( 
    "fmt"
    "encoding/json"
) 
   

type Human struct{ 
       
    
    Name string 
    Address string 
    Age int
} 
   

func main() { 
       
    
    var human1 Human 
       
    
    Data := []byte(`{ 
        "Name": "jhansi",   
        "Address": "Banglore", 
        "Age": 23 
    }`) 
       
    
    err := json.Unmarshal(Data, &human1) 
       
    if err != nil { 
           
        
            fmt.Println(err) 
    } 
       
    
    fmt.Println("Struct is:", human1) 
    fmt.Printf("%s lives in %s.\n", human1.Name, human1.Address) 
       
    
    var human2 []Human 
       
   
    Data2 := []byte(` 
    [ 
        {"Name": "supraja", "Address": "hyderabad", "Age": 27}, 
        {"Name": "bavitha", "Address": "chennai", "Age": 18}, 
        {"Name": "prathyusha", "Address": "ban", "Age": 23} 
    ]`) 
       
    
    err2 := json.Unmarshal(Data2, &human2) 
       
        if err2 != nil { 
       
         
            fmt.Println(err2) 
        } 
       
    
    for i := range human2{ 
       
        fmt.Println(human2[i]) 
    } 
} 

output:
Struct is: {jhansi Banglore 23}
jhansi lives in Banglore.
{supraja hyderabad 27}
{bavitha chennai 18}
{prathyusha ban 23}

Program exited.