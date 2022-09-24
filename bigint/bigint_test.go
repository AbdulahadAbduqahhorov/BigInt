package bigint

import (
	
	"testing"
)

type test struct {
	arg1, arg2 Bigint
	expected   string
}

func TestNewInt(t1 *testing.T) {
	addTests := []struct {
		arg      string
		err      error
		expected Bigint
	}{
		{"00000000000000000042", nil, Bigint{Value: "42"}},
		{"+0000000000000042", nil, Bigint{Value: "42"}},
		{"-42", nil, Bigint{Value: "-42"}},
		{"-000000000000420", nil, Bigint{Value: "-420"}},
		{"-0000000000000", nil, Bigint{Value: "0"}},
		{"000000", nil, Bigint{Value: "0"}},
		{"++23", ErrorBadInput, Bigint{Value: "++23"}},
		{"---42", ErrorBadInput, Bigint{Value: "---42"}},
		{"a42", ErrorBadInput, Bigint{Value: "a42"}},
		{"4a2", ErrorBadInput, Bigint{Value: "4a2"}},
		{"--00042", ErrorBadInput, Bigint{Value: "--00042"}},
		
	}
	for _, test := range addTests {

		res, err := NewInt(test.arg)
		if res != test.expected || err != test.err {
			t1.Errorf("got %v but expected %v, input values: %v, %v", res.Value, test.expected, test.arg, test.err)
		}

	}
}


func TestSet(t1 *testing.T) {
	addTests := []struct {
		arg      Bigint
		expected error
	}{
		{Bigint{Value:"123"}, nil},
		{Bigint{Value:"00000000000000000042"}, nil},
		{Bigint{Value:"+0000000000000042"}, nil},
		{Bigint{Value:"-42"}, nil},
		{Bigint{Value:"-000000000000420"}, nil},
		{Bigint{Value:"-0000000000000"}, nil},
		{Bigint{Value:"000000"}, nil},
		{Bigint{Value:"++23"}, ErrorBadInput},
		{Bigint{Value:"---42"}, ErrorBadInput},
		{Bigint{Value:"a42"}, ErrorBadInput},
		{Bigint{Value:"4a2"}, ErrorBadInput},
		{Bigint{Value:"-+00042"}, ErrorBadInput},
		
		
	}
	for _, test := range addTests {

		err := test.arg.Set(test.arg.Value)
		if err != test.expected  {
			t1.Errorf("got %v but expected %v, input values: %v, %v",err,test.expected,test.arg.Value, test.expected)
		}

	}
}


func TestAdd(t1 *testing.T) {
	addTests := []test{
		{Bigint{Value: "42"}, Bigint{Value: "1231"}, "1273"},
		{Bigint{Value: "-42"}, Bigint{Value: "1231"}, "1189"},
		{Bigint{Value: "42"}, Bigint{Value: "-1231"}, "-1189"},
		{Bigint{Value: "-42"}, Bigint{Value: "-1231"}, "-1273"},
		{Bigint{Value: "0"}, Bigint{Value: "0"}, "0"},
		{Bigint{Value: "0"}, Bigint{Value: "1231"}, "1231"},
		{Bigint{Value: "433332"}, Bigint{Value: "1231"}, "434563"},
		{Bigint{Value: "-23423423424"}, Bigint{Value: "42342342"}, "-23381081082"},
	}
	for _, test := range addTests {

		res := Add(test.arg1, test.arg2)
		if res.Value != test.expected {
			t1.Errorf("got %v but expected %v, input values: %v, %v", res.Value, test.expected, test.arg1.Value, test.arg2.Value)
		}

	}
}

func TestSub(t1 *testing.T) {
	addTests := []test{
		{Bigint{Value: "42"}, Bigint{Value: "1231"}, "-1189"},
		{Bigint{Value: "-42"}, Bigint{Value: "1231"}, "-1273"},
		{Bigint{Value: "42"}, Bigint{Value: "-1231"}, "1273"},
		{Bigint{Value: "-42"}, Bigint{Value: "-1231"}, "1189"},
		{Bigint{Value: "0"}, Bigint{Value: "0"}, "0"},
		{Bigint{Value: "0"}, Bigint{Value: "1231"}, "-1231"},
		{Bigint{Value: "-23423423424"}, Bigint{Value: "42342342"}, "-23465765766"},
	}
	for _, test := range addTests {

		res := Sub(test.arg1, test.arg2)
		if res.Value != test.expected {
			t1.Errorf("got %v but expected %v, input values: %v, %v", res.Value, test.expected, test.arg1.Value, test.arg2.Value)
		}

	}
}

func TestMultiply(t1 *testing.T) {
	addTests := []test{
		{Bigint{Value: "0"}, Bigint{Value: "0"}, "0"},
		{Bigint{Value: "0"}, Bigint{Value: "23"}, "0"},
		{Bigint{Value: "42"}, Bigint{Value: "0"}, "0"},
		{Bigint{Value: "3456789867546352434556574355"}, Bigint{Value: "-123546357648768635735735356745673567567356"}, "-427073797292521218320706066550115597373279990022958574453854084755380"},
		{Bigint{Value: "-2357"}, Bigint{Value: "232"}, "-546824"},
		{Bigint{Value: "-345"}, Bigint{Value: "-12"}, "4140"},
		{Bigint{Value: "23423423424"}, Bigint{Value: "42342342"}, "991802605429819008"},
	}
	for _, test := range addTests {

		res := Multiply(test.arg1, test.arg2)
		if res.Value != test.expected {
			t1.Errorf("got %v but expected %v, input values: %v, %v", res.Value, test.expected, test.arg1.Value, test.arg2.Value)
		}

	}
}

func TestModAndDivision(t1 *testing.T) {
	addTests := []struct {
		arg1, arg2           Bigint
		expected1, expected2 string
	}{
		{Bigint{Value: "0"}, Bigint{Value: "112312"}, "0", "0"},
		{Bigint{Value: "3"}, Bigint{Value: "34234"}, "0", "3"},
		{Bigint{Value: "12314532"}, Bigint{Value: "34"}, "362192", "4"},
		{Bigint{Value: "1462"}, Bigint{Value: "34"}, "43", "0"},
		{Bigint{Value: "45"}, Bigint{Value: "34"}, "1", "11"},
		{Bigint{Value: "21"}, Bigint{Value: "23"}, "0", "21"},

		{Bigint{Value: "-2342352"}, Bigint{Value: "2342"}, "-1001", "1990"},
		{Bigint{Value: "-3"}, Bigint{Value: "5"}, "-1", "2"},
		{Bigint{Value: "-21"}, Bigint{Value: "7"}, "-3", "0"},
		{Bigint{Value: "-34"}, Bigint{Value: "123"}, "-1", "89"},
		{Bigint{Value: "0"}, Bigint{Value: "-123"}, "0", "0"},

		{Bigint{Value: "23423423"}, Bigint{Value: "-234234"}, "-101", "-234211"},
		{Bigint{Value: "21"}, Bigint{Value: "-23"}, "-1", "-2"},
		{Bigint{Value: "1"}, Bigint{Value: "-23"}, "-1", "-22"},
		{Bigint{Value: "0"}, Bigint{Value: "-23"}, "0", "0"},
		{Bigint{Value: "70000000000000010000001000000001"}, Bigint{Value: "-23"}, "-3043478260869565652173956521740", "-19"},
		{Bigint{Value: "21"}, Bigint{Value: "23"}, "0", "21"},

		{Bigint{Value: "-10000000000000000100000034300"}, Bigint{Value: "-111001"}, "90089278474968694876623", "-4677"},
		{Bigint{Value: "-3"}, Bigint{Value: "-3"}, "1", "0"},
		{Bigint{Value: "-3"}, Bigint{Value: "-5"}, "0", "-3"},
		{Bigint{Value: "-72"}, Bigint{Value: "-12"}, "6", "0"},
		{Bigint{Value: "-341"}, Bigint{Value: "-121"}, "2", "-99"},
		{Bigint{Value: "-34100000000000101010111111111000000001"}, Bigint{Value: "-32423424"}, "1051708789299985745185675365778", "-24816129"},
		{Bigint{Value: "-874837385784530458437054835720845374058932452708453407589324504507309824532457209845908694856049586093587435234534594264563463456546"}, Bigint{Value: "-35734057893457030894578245037459834705983457039485730489573409853845089345039459359359345"}, "24481893111409402196867326316544145441501139", "-15971641218588295202337881222903316006223308848483770284208593982410727676378177335662591"},
	}
	for _, test := range addTests {

		quotient,remainder := ModAndDivision(test.arg1, test.arg2)
		if quotient.Value != test.expected1 || remainder.Value!=test.expected2{
			t1.Errorf("got quotient:%v remainder:%v  but expected quotient:%v remainder:%v, input values: %v, %v", quotient.Value, remainder.Value,test.expected1,test.expected2,test.arg1,test.arg2)
		}

	}
}

func TestAbs(t1 *testing.T) {
	addTests := []struct {
		arg,expected      Bigint
		 
	}{
		{Bigint{Value:"123"},Bigint{Value:"123"} },
		{Bigint{Value:"-123"},Bigint{Value:"123"} },
		
		
		
		
	}
	for _, test := range addTests {

		got := test.arg.Abs()
		if got != test.expected  {
			t1.Errorf("got %v but expected %v, input value: %v",got.Value,test.expected.Value,test.arg.Value)
		}

	}
}

