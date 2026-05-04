package intermediate

import "fmt"

func main() {
	looper(1)
	fac :=factorial(5)
	sum:= sumOfDigits(16)
	fmt.Println("factorial is :",fac)
	fmt.Println("Sum is :",sum)
}

func looper(x int) int {
	if x < 5 {
		fmt.Println(x)
		x++
		looper(x)
	}
	return x
}

func factorial(x int) int{
	if (x == 0) {
		return 1
	}
	
	return x*factorial(x-1)
	
}

func sumOfDigits( n int) int{

	if(n < 10){
		return n
	}else{
		return n%10 + sumOfDigits(n/10)
	}
}
