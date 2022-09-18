package bigint

import (
	"errors"
	"strconv"
	"strings"
)

type Bigint struct {
	Value string
}

func removeZeros(num string) string {
	p := 0
	if strings.HasPrefix(num, "-") {
		p = 1
		num = num[1:]
	} else if strings.HasPrefix(num, "+") {
		num = num[1:]
	}

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

var ErrorBadInput = errors.New("bad input, please input only number")

func NewInt(num string) (Bigint, error) {
	err, num := validateNumber(num)
	if err {
		return Bigint{Value: num}, ErrorBadInput
	} else {
		num = removeZeros(num)
		return Bigint{Value: num}, nil

	}
}

func (z *Bigint) Set(num string) error {
	allowed := "1234567890-"
	var err bool

	if strings.HasPrefix(num, "+") {
		num = strings.Replace(num, "+", "", 1)
	}
	if strings.HasPrefix(num, "0") {
		err = true
	}
	arr := strings.Split(num, "")
	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			err = true
		}

	}
	if err {
		return ErrorBadInput
	}
	z.Value = num
	return nil
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Add(a, b Bigint) Bigint {
	string1 := a.Value
	string2 := b.Value
	sum := ""

	if string(string1[0]) != "-" && string(string2[0]) == "-" {
		return Sub(Bigint{Value: string1}, Bigint{Value: string2[1:]})
	}
	if string(string1[0]) == "-" && string(string2[0]) != "-" {
		return Sub(Bigint{Value: string2}, Bigint{Value: string1[1:]})
	}
	var flag bool
	if string(string1[0]) == "-" && string(string2[0]) == "-" {
		string1 = string1[1:]
		string2 = string2[1:]
		flag = true
	}

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
	if flag {
		sum = "-" + sum
	}
	return Bigint{
		Value: sum,
	}
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

func Sub(a, b Bigint) Bigint {

	string1 := a.Value
	string2 := b.Value

	if string(string1[0]) != "-" && string(string2[0]) == "-" {
		return Add(Bigint{Value: string1}, Bigint{Value: string2[1:]})
	}
	if string(string1[0]) == "-" && string(string2[0]) != "-" {
		return Add(Bigint{Value: string1}, Bigint{Value: "-" + string2})
	}

	if string(string1[0]) == "-" && string(string2[0]) == "-" {
		x := string1[1:]
		string1 = string2[1:]
		string2 = x

	}
	compare := compareStrings(string1, string2)

	if compare == -1 {
		new := string2
		string2 = string1
		string1 = new

	} else if compare == 0 {
		return Bigint{
			Value: "0",
		}
	}
	sub := ""
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
	for strings.HasPrefix(sub, "0") {
		sub = sub[1:]
	}

	if compare == -1 {
		sub = "-" + sub

	}

	return Bigint{
		Value: sub,
	}

}

func Multiply(a, b Bigint) Bigint {
	string1 := a.Value
	string2 := b.Value
	if string1 == "0" || string2 == "0" {
		return Bigint{Value: "0"}
	}
	if len(string1) < len(string2) {
		x := string1
		string1 = string2
		string2 = x
	}

	mySlice := make([]string, 0)

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
		mySlice = append(mySlice, mult)
	}
	res := Bigint{Value: "0"}
	for _, v := range mySlice {
		value := Bigint{Value: v}
		res = Add(res, value)
	}
	return res
}

func Mod(a, b Bigint) Bigint {
	// string1 :=a.Value
	// string2 :=b.Value
	// l := len(string2)-1

	// for l<len(string1){
	// 	str1 := string1[:l]

	// }
	return Bigint{Value: ""}
}

func (x *Bigint) Abs() Bigint {
	if x.Value[0] == '-' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	if x.Value[0] == '+' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	return Bigint{
		Value: x.Value,
	}
}
