package main

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"math"
)

const ReadOn = 1
const LiteralValue = 4
const AmountSubpackets = 11
const LengthSubPackets = 15

//go:embed input.txt
var input []byte

func getByteValues(data []byte) []byte {
	byte_values, _ := hex.DecodeString(string(data))
	return byte_values
}

func getValueFromBits(start *int, amount int, data *[]byte) int {
	end_bit := *start + amount
	value := 0
	// value at bit 15 -> start
	for *start < end_bit {
		// Check at which byte the counter is
		byte_count := *start / 8
		// Grab which position you have in the byte
		pos_in_byte := 7 - *start%8
		// Shift the previous values for the next bit
		value = value << 1
		// If the bit is set in the byte, add the value
		if (*data)[byte_count]&(1<<pos_in_byte) != 0 {
			value++
		}
		// Go to the next bit
		*start++
	}
	return value
}

func readLiteral(start *int, data *[]byte) int {
	read_on := getValueFromBits(start, ReadOn, data)
	value := getValueFromBits(start, LiteralValue, data)

	// Read while read_on is not 0
	for read_on == 1 {
		// Shift the values to the next nibble
		value = value << 4
		read_on = getValueFromBits(start, 1, data)
		value += getValueFromBits(start, 4, data)
	}
	return value
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func prod(array []int) int {
	result := 1
	for _, v := range array {
		result *= v
	}
	return result
}

func min(array []int) int {
	result := int(math.MaxInt32)
	for _, v := range array {
		if v < result {
			result = v
		}
	}
	return result
}

func max(array []int) int {
	result := 0
	for _, v := range array {
		if v > result {
			result = v
		}
	}
	return result
}

func readSubPackets(start *int, data *[]byte, type_value int) (int, int) {
	var values []int
	length_type := getValueFromBits(start, 1, data)
	versions_sum := 0
	value := 0

	// Check length type on which way to handle the sub packets
	if length_type == 0 {
		// Get total length that will be read
		length := getValueFromBits(start, LengthSubPackets, data) + *start
		for *start < length {
			version, value := readPacket(start, data)
			versions_sum += version
			values = append(values, value)
		}
	} else {
		amount_subpackets := getValueFromBits(start, AmountSubpackets, data)
		for i := 0; i < amount_subpackets; i++ {
			version, value := readPacket(start, data)
			versions_sum += version
			values = append(values, value)
		}
	}

	switch type_value {
	case 0:
		// Sum
		value = sum(values)
		break
	case 1:
		// Product
		value = prod(values)
		break
	case 2:
		// Min
		value = min(values)
		break
	case 3:
		// Max
		value = max(values)
		break
	case 5:
		// >
		// This operator should have values length 2
		if len(values) == 2 {
			if values[0] > values[1] {
				value = 1
			}
		}
		break
	case 6:
		// <
		// This operator should have values length 2
		if len(values) == 2 {
			if values[0] < values[1] {
				value = 1
			}
		}
		break
	case 7:
		// ==
		// This operator should have values length 2
		if len(values) == 2 {
			if values[0] == values[1] {
				value = 1
			}
		}
		break
	default:
		fmt.Println("Error")
		break
	}

	return versions_sum, value
}

func readPacket(start *int, data *[]byte) (int, int) {
	value := 0
	versions_sum := 0
	version := getValueFromBits(start, 3, data)
	type_value := getValueFromBits(start, 3, data)

	if type_value == 4 {
		// Literal
		value = readLiteral(start, data)
	} else {
		versions_sum, value = readSubPackets(start, data, type_value)
		version += versions_sum
	}
	return version, value
}

func BITS(data []byte) (int, int) {
	pos := 0
	bytes := getByteValues(data)
	ver, val := readPacket(&pos, &bytes)
	return ver, val
}

func main() {
	fmt.Println(BITS(input))
}
