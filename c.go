package main

import(
	"fmt"
	"strings"
	"strconv"
)

func main(){
	var input string
	var op, x, y string
	var arabNumber = true
	fmt.Println("Введите строку")
	fmt.Scan(&input)
	strings.ToLower(input)
	
	t, op := validInput(input)
	if !t {
		fmt.Println("Недопустимый ввод.")
	}	
	if validArab(x)!=validArab(y){
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return 
	}
	x,y=splitNumber(input, op)
	arabNumber=validArab(y)	

	if arabNumber{		
		a, _:=strconv.Atoi(x)
		b, _:=strconv.Atoi(y)
		if !validNumber(a,b){
			fmt.Println("Числа должны быть в заданом диапазоне от 1 до 10 включительно.")
			return 
		}
		
		fmt.Println(calc(op, a, b))
		return
	}else{
		a:=romeToArab(x)
		b:=romeToArab(y)
		if !validNumber(a,b){
			fmt.Println("Числа должны быть в заданом диапазоне от 1 до 10 включительно.")
			return 
		}
		res:=calc(op, a, b)
		if res<1{
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		}else{
			fmt.Println(arabToRome(res))
		}
	}
	
}

func romeToArab(rom string) int{
		rome_number:=map[string]int{
		"i":1, "ii":2, "iii":3, "iv":4,
		"v":5, "vi":6,"vii":7,"viii":8,
		"ix":9, "x":10,
		"l":50,
		"c":100,
	}

	return rome_number[rom]
}

func validInput(str string) (bool, string){

	if strings.Count(str, "+")==1{
		return true, "+"
	}
	if strings.Count(str, "-")==1{
		return true, "-"
	}
	if strings.Count(str, "/")==1{
		return true, "/"
	}
	if strings.Count(str, "*")==1{
		return true, "*"
	}
	return false, "err"
}

func calc(op string, x,y int) int{

	switch op {
	case "+":
		return x+y
	case "-":
		return x-y
	case "*":
		return x*y
	case "/":
		return x/y
	default:
		fmt.Println("Something went wrong:(")
		return 0
	}
}

func splitNumber(s, op string) (string, string){
	m := strings.Split(s, op)
	a := strings.Trim(m[0], " ")
	b := strings.Trim(m[1], " ")

	return a,b
}

func validNumber(x,y int) bool{
	if (x > 10 || x < 1) || (y > 10 || y <1){
		return false
	}
	
	return true
}

func validArab(x string) bool{
	if strings.Contains(x, "i") || strings.Contains(x, "v") || strings.Contains(x, "x"){
		return false
	}

	return true
}

func arabToRome(num int) string{
	var res string
	var arr = []int{100,50,10,9,8,7,6,5,4,3,2,1}
	rome_number:=map[int]string{
		1:"i", 2:"ii", 3:"iii", 4:"iv",
		5:"v", 6:"vi", 7:"vii", 8:"viii",
		9:"ix", 10:"x",
		50:"l",	100:"c",
	}
	for _, i := range arr{
		t:=num/i
		if t>0{
			num-=i
		}
		for j:=0; j<t; j++{
			res+=rome_number[i]
		}
	}

	return strings.ToUpper(res)
}