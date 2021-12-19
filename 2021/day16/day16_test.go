package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseLiteralValue(t *testing.T) {
	test := map[string][]int{"D2FE28": {2021}}
	for s := range test {
		input := getByteValues([]byte(s))
		start := 6
		value := readLiteral(&start, &input)
		if value != (test[s])[0] {
			fmt.Println(value, test[s])
			t.Fail()
		}
	}
}

func TestGetValueFromBits(t *testing.T) {
	test := "0036"
	input := getByteValues([]byte(test))
	val := 0
	start := 0
	val = getValueFromBits(&start, 15, &input)
	if val != 27 {
		t.Fail()
	}
}

func TestGetByteValues(t *testing.T) {
	test := map[string][]byte{"D2FE28": {210, 254, 40}}
	for s := range test {
		output := getByteValues([]byte(s))
		if !reflect.DeepEqual(output, test[s]) {
			t.Fail()
		}
	}
}

func TestReadPacket(t *testing.T) {
	input := getByteValues([]byte("9C0141080250320F1802104A08"))
	fmt.Println(input)
	start := 0
	ver, val := readPacket(&start, &input)
	fmt.Println(ver, val)
	if ver != 20 || val != 1 {
		t.Fail()
	}
}
