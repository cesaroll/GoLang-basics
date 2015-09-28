package main

//import "fmt"

import(
    "fmt"
    "math"
    "math/rand"
    "life"
)

func main(){

    //Hello World
    fmt.Printf("Hello World! \n")


    //Function
    test()


    //Math
    mymath()
    
    //life
    mylife()
    
}


func test() {
    fmt.Printf("This is a friggin test")
}

func mymath() {
    fmt.Println("Pi equals ", math.Pi)
    rand.Seed(123456789)
    fmt.Println("random number: ", rand.Intn(10))
}

func mylife() {
    fmt.Println("Life equals: ", life.Life, " Half Life equals: ", life.Halflife())
}