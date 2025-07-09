package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strings"
)

func calculation(a int, b int, c string)int{
	ans := 0
	switch c {
		case "+": {
			ans =  a+b
			}
		case "-":{
			ans =  a-b
		}
		case "*":{
			ans =  a*b
		}
		case "/":{
			ans =  a/b
		}
	}
	return ans
}

func even_odd(a int) string {
	if (a%2) == 0 {
		return "the number is even"
	}else{
		return "the number is odd"
	}
}

func first_10_square()[]float64{
	lis := []float64{}
	for i:=1;i<=10;i++{
		pow := math.Pow(float64(i),float64(2))
		lis = append(lis,pow)
	}
	return lis
}

func sum(a int, b int) int {
	return a+b
}
func isEven(n int) bool {
	if (n%2) == 0{
		return true
	}else{
		return false
	}
}

func swap(a, b int) (int, int){
	a, b = b, a
	return a,b
}

func sum_average_reverse(arr []int)(int, int, []int){
	var sum int
	n := len(arr)
	//sum
	for _, val := range arr{
		sum += val
	}
	//avg
	avg := sum/len(arr)
	//reverse
	for i:=0;i<n/2;i++{
		arr[i], arr[n-1-i] =arr[n-1-i], arr[i] 
	}

	return (sum), (avg), (arr)
}

func count_words(c string) map[string]int {
	words := strings.Fields(c)
	m := make(map[string]int)

	for _, val := range words {
		m[val] += 1
	}
	return m
}

func main(){
	var a int
	var b int
	var c string
	var text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the a")
	fmt.Scanln(&a)
	fmt.Println("enter the b")
	fmt.Scanln(&b)
	fmt.Println("enter the operator")
	fmt.Scanln(&c)
	fmt.Println("the return is:",calculation(a,b,c))
	fmt.Println(even_odd(a))
	fmt.Println(first_10_square())
	fmt.Println(sum(a,b))
	fmt.Println(isEven(a))
	fmt.Println(swap(a, b))
	var lis  = []int{2,4,8,6}
	fmt.Println(sum_average_reverse(lis))
	fmt.Println("enter the statement")
	text, _ = reader.ReadString('\n')
	fmt.Println("the list\n", count_words(text))
}

