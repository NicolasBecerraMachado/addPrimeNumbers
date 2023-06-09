package main

import(
	"os"
	"strconv"
	"fmt"
	"strings"
	"math"
	//"sync"
	"time"
)

var tokens = make(chan struct{}, 5000000)
var added = make(chan int, 1)
var add = make(chan int, 5000000)

func isPrime(n int){
	stop := int(math.Sqrt(float64(n)))
	for i := 2 ; i < stop+1; i++{
		if n%i == 0{
			<- tokens
			return
		}
	}
	add <- n
	<- tokens
}


func adder(){
	sum := 0
	for len(add) != 0 || len(tokens) != 0{
		sum += <- add
	}
	added <- sum
}

// func to calculate and print execution time
func exeTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}

func main(){
	defer exeTime("main")()
	n,err := strconv.Atoi(strings.Split(os.Args[1],":")[1])
	if err != nil{
		os.Exit(1)
	}

	if n < 2{
		return
	}

	for i := 2; i < n; i++{
		tokens <- struct{}{}
		go isPrime(i)
	}
	
	//fmt.Println("we did it boyz")

	go adder()

	sum := <- added

	fmt.Println(sum)
}