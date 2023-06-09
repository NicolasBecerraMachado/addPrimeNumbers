package main

import(
	"os"
	"strconv"
	"fmt"
	"strings"
	"math"
	//"sync"
	//"time"
)

var tokens = make(chan struct{}, 5000000)

func isPrime(n int, add chan int){
	stop := int(math.Sqrt(float64(n)))
	//fmt.Println("for n = ", n, " stop = ", stop)
	for i := 2 ; i < stop+1; i++{
		if n%i == 0{
			//fmt.Println("about to evict tokens")
			<- tokens
			//fmt.Println("DONE")
			return
		}
	}
	add <- n
	//fmt.Println("about to evict tokens")
	<- tokens
	//fmt.Println("DONE")
}


func adder(add chan int)(int){
	sum := 0
	for len(add) != 0{
		sum += <- add
	}
	return sum
}

func main(){

	add := make(chan int, 5000000)

	n,err := strconv.Atoi(strings.Split(os.Args[1],":")[1])
	if err != nil{
		//fmt.Println("invalid arg")
		os.Exit(1)
	}

	if n < 2{
		//fmt.Println(0)
		return
	}

	for i := 2; i < n; i++{
		tokens <- struct{}{}
		//fmt.Println("tok send")
		go isPrime(i, add)
	}
	
	//fmt.Println("tokens = " , len(tokens))
	//fmt.Println("addL = ", len(add))

	for len(tokens) != 0 {
		//fmt.Println("waiting for tokens, len = " ,len(tokens))
		//time.Sleep(1 * time.Second)
	}

	sum := adder(add)

	fmt.Println(sum)
}