package main

import(
	"os"
	"strconv"
	"fmt"
	"strings"
	"math"
)

func isPrime(n int) (int){
	if n == 2{
		return 2
	}else{
		stop := int(math.Sqrt(float64(n)))
		//fmt.Println("for ",n," stop is = ", stop)
		for i := 2 ; i < stop+1; i++{
			//fmt.Println("mod = ",n%i)
			if n%i == 0{
				//fmt.Println("no sum")
				return 0
			}
		}
	}
	return n
}

func main(){
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