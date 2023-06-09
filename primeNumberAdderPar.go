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
	for len(add) != 0{
		//fmt.Println("len toks = ", len(tokens))
		sum += <- add
		//fmt.Println("sum = ",sum)
	}
	added <- sum
	//fmt.Println("added :)")
}

func main(){
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