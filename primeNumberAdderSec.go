package main

import(
	"os"
	"strconv"
	"fmt"
	"strings"
	"math"
	"time"
)

func isPrime(n int) (int){
	if n == 2{
		return 2
	}else{
		stop := int(math.Sqrt(float64(n)))
		for i := 2 ; i < stop+1; i++{
			if n%i == 0{
				return 0
			}
		}
	}
	return n
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
		fmt.Println("invalid arg")
		os.Exit(1)
	}

	if n < 2{
		fmt.Println(0)
		return
	}

	sum := 0

	for i := 2; i < n; i++{
		sum += isPrime(i)
	}

	fmt.Println(sum)
}