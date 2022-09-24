package bigint

import (
	"errors"
	"strconv"
	"strings"
	
)

var ErrorBadInput = errors.New("bad input, please input only number")
var ErrorZeroDivision = errors.New("ZeroDivisionError: integer division or modulo by zero")
type Bigint struct {
	Value string
}

func removeZeros(num string) string {
	p := 0
	if strings.HasPrefix(num, "-") {
		p = 1
		num = num[1:]
	} 
	num = strings.TrimPrefix(num,"+")

	for strings.HasPrefix(num, "0") && len(num) > 1 {
		num = num[1:]
	}

	if p == 1 && num != "0" {
		num = "-" + num
	}

	return num
}

func validateNumber(num string) (bool, string) {
	allowed := "1234567890"
	var err bool

	start := 0
	if strings.HasPrefix(num, "+") || strings.HasPrefix(num, "-") {
		start = 1
	}

	arr := strings.Split(num[start:], "")
	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			err = true
		}
	}

	return err, num
}

func NewInt(num string) (Bigint, error) {
	err, num := validateNumber(num)
	if err {
		return Bigint{Value: num}, ErrorBadInput
	}
	num = removeZeros(num)
	return Bigint{Value: num}, nil

}

func (z *Bigint) Set(num string) error {
	err, num := validateNumber(num)
	if err {
		return ErrorBadInput
	}
	num = removeZeros(num)
	z.Value = num
	return nil
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Add(a, b Bigint) (total Bigint) {
	string1 := a.Value
	string2 := b.Value
	

	if string(string1[0]) != "-" && string(string2[0]) == "-" {
		total.Value =  sub(string1,string2[1:])
		return 
	}
	if string(string1[0]) == "-" && string(string2[0]) != "-" {
		total.Value =  sub(string2,string1[1:])
		return
	}
	if string(string1[0]) == "-" && string(string2[0]) == "-" {
		total.Value= "-"+add(string1[1:],string2[1:])
		return 
	}
	total.Value = add(string1, string2)
	return 
}


func add(string1,string2 string)string{
	sum :=""
	reminder := 0
	for i := 0; i < max(len(string1), len(string2)); i++ {
		number1, number2 := 0, 0

		if i < len(string1) {
			number1 = int(string1[len(string1)-1-i] - '0')
		}

		if i < len(string2) {
			number2 = int(string2[len(string2)-1-i] - '0')
		}
		summa := number1 + number2 + reminder
		sum = strconv.Itoa(summa%10) + sum
		reminder = summa / 10
	}

	if reminder > 0 {
		sum = strconv.Itoa(reminder) + sum
	}
	return sum

}

func compareStrings(string1, string2 string) int {

	res := 0
	if len(string1) > len(string2) {
		res = 1
	} else if len(string1) < len(string2) {
		res = -1
	} else {
		for i := 0; i < len(string1); i++ {
			if string1[i] > string2[i] {
				res = 1
				break

			} else if string1[i] < string2[i] {
				res = -1
				break
			}
		}
	}
	return res
}

func Sub(a, b Bigint) (subResult Bigint ){
	
	string1 := a.Value
	string2 := b.Value

	if string(string1[0]) != "-" && string(string2[0]) == "-" {
		subResult.Value=add(string1,string2[1:])
		return 
	}
	if string(string1[0]) == "-" && string(string2[0]) != "-" {
		subResult.Value ="-"+add(string1[1:],string2)
		return 
	}

	if string(string1[0]) == "-" && string(string2[0]) == "-" {
		subResult.Value = sub(string2[1:],string1[1:])
		return
	}

	subResult.Value = sub(string1,string2)
	return
}

func sub(string1, string2 string) (sub string){
	compare := compareStrings(string1, string2)

	if compare == -1 {
		string1,string2 = string2,string1

	} else if compare == 0 {
		return "0"
	}
	
	
	carry := 0
	for i := 0; i < max(len(string1), len(string2)); i++ {
		number1, number2 := 0, 0

		if i < len(string1) {
			number1 = int(string1[len(string1)-1-i] - '0')
		}

		if i < len(string2) {
			number2 = int(string2[len(string2)-1-i] - '0')
		}

		subtract := number1 - number2 - carry

		if subtract < 0 {
			subtract += 10
			carry = 1
		} else {
			carry = 0
		}
		sub = strconv.Itoa(subtract) + sub

	}
	

	if compare == -1 {
		sub = "-" + sub

	}
	sub = removeZeros(sub)
	return 

}

func Multiply(a, b Bigint) Bigint {
	string1 := a.Value
	string2 := b.Value

	if string1 == "0" || string2 == "0" {
		return Bigint{Value: "0"}
	}
	
	var flag bool
	if string(string1[0]) == "-" && string(string2[0]) != "-" {
		flag = true
		string1 = string1[1:]
	}
	if string(string1[0]) != "-" && string(string2[0]) == "-" {
		flag = true
		string2 = string2[1:]
	}

	if string(string1[0]) == "-" && string(string2[0]) == "-" {
		string1 = string1[1:]
		string2 = string2[1:]
	}

	if len(string1) < len(string2) {
		string1,string2 = string2,string1
	}


	result := "0"
	for i := len(string2) - 1; i > -1; i-- {
		reminder := 0
		mult := ""

		for j := len(string1) - 1; j > -1; j-- {
			sum := int(string1[j]-'0')*int(string2[i]-'0') + reminder
			reminder = sum / 10
			mult = strconv.Itoa(sum%10) + mult

		}
		if reminder != 0 {
			mult = strconv.Itoa(reminder) + mult
		}
		for k := len(string2) - 1 - i; k > 0; k-- {
			mult += "0"
		}
		result = add(result,mult)

	}
	

	if flag {
		result = "-" + result
	}
	return Bigint{Value:result}
}

func ModAndDivision(a, b Bigint) ( Bigint, Bigint) {
	
	string1 := a.Value
	string2 := b.Value
	quotient:=""
	remainder:= ""
	
	if string2 == "0" {
		panic(ErrorZeroDivision)
	}
	if string1 == "0" {
		return Bigint{Value:"0"},Bigint{Value:"0"}
	}



	if string(string1[0]) == "-" && string(string2[0]) != "-" {
		quotient,remainder= mod(string1[1:], string2)
		
		if remainder!="0"{
			remainder = sub(string2, remainder)
			quotient = add(quotient,"1")
		}
		return Bigint{Value:"-"+quotient},Bigint{Value:remainder}
	} 


	if string(string1[0]) != "-" && string(string2[0]) == "-" {

		quotient, remainder = mod(string1, string2[1:])
		
		if remainder!="0"{
			remainder = "-" + sub(string2[1:],remainder)
			quotient = add(quotient,"1")
		}
		return Bigint{Value:"-"+quotient},Bigint{Value:remainder}
	}  
	if string(string1[0]) == "-" && string(string2[0]) == "-" {
		quotient, remainder = mod(string1[1:], string2[1:])
		if remainder!="0"{
			remainder = "-" + remainder
		}
		return Bigint{Value:quotient},Bigint{Value:remainder}

	} 

	quotient,remainder = mod(string1, string2)

	

	return Bigint{Value:quotient},Bigint{Value:remainder}
}

func mod(string1, string2 string) (string, string) {
	
	if compareStrings(string1, string2) == -1 {
		return "0",string1
	}

	l := len(string2)
	dividend := string1[:l]
	divisor:= string2
	quotient := ""
	remainder:= ""
	
	count := 0	
	for {
		if compareStrings(dividend, divisor) >= 0 {
			dividend = sub(dividend,divisor)
			count++
		} else if l >= len(string1) {
			break
		} else {

			c := strconv.Itoa(count)
			quotient = quotient + c
			count = 0
			
			dividend = dividend + string(string1[l])
			dividend = removeZeros(dividend)
			l++
		}
	}
	quotient += strconv.Itoa(count)
	quotient,remainder = removeZeros(quotient),removeZeros(dividend)
	return quotient,remainder
}

func (x Bigint) Abs() Bigint {
	if x.Value[0] == '-' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	
	return Bigint{
			Value: x.Value,
		}
		

}


