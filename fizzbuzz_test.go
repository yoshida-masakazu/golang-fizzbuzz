package fizzbuzz

import (
	"testing"
)

func Test_input_0_expected_EmptyCharacters(t *testing.T) {
	input := 0
	result := FizzBuzz(input)
	if result != "" {
		t.Errorf("expected '' but got '%s'", result)
	}
}

func Test_input_1_expected_1(t *testing.T) {
	input := 1
	result := FizzBuzz(input)
	if result != "1" {
		t.Errorf("expected '1' but got '%s'", result)
	}
}

func Test_input_2_expected_2(t *testing.T) {
	input := 2
	result := FizzBuzz(input)
	if result != "2" {
		t.Errorf("expected '2' but got '%s'", result)
	}
}

func Test_input_3_expected_Fizz(t *testing.T) {
	input := 3
	result := FizzBuzz(input)
	if result != "Fizz" {
		t.Errorf("expected 'Fizz' but got '%s'", result)
	}
}

func Test_input_4_expected_4(t *testing.T) {
	input := 4
	result := FizzBuzz(input)
	if result != "4" {
		t.Errorf("expected '4' but got '%s'", result)
	}
}

func Test_input_5_expected_Buzz(t *testing.T) {
	input := 5
	result := FizzBuzz(input)
	if result != "Buzz" {
		t.Errorf("expected 'Buzz' but got '%s'", result)
	}
}

func Test_input_6_expected_Fizz(t *testing.T) {
	input := 6
	result := FizzBuzz(input)
	if result != "Fizz" {
		t.Errorf("expected 'Fizz' but got '%s'", result)
	}
}

func Test_input_7_expected_7(t *testing.T) {
	input := 7
	result := FizzBuzz(input)
	if result != "7" {
		t.Errorf("expected '7' but got '%s'", result)
	}
}

func Test_input_8_expected_8(t *testing.T) {
	input := 8
	result := FizzBuzz(input)
	if result != "8" {
		t.Errorf("expected '8' but got '%s'", result)
	}
}

func Test_input_9_expected_Fizz(t *testing.T) {
	input := 9
	result := FizzBuzz(input)
	if result != "Fizz" {
		t.Errorf("expected 'Fizz' but got '%s'", result)
	}
}

func Test_input_10_expected_Buzz(t *testing.T) {
	input := 10
	result := FizzBuzz(input)
	if result != "Buzz" {
		t.Errorf("expected 'Buzz' but got '%s'", result)
	}
}

func Test_input_11_expected_11(t *testing.T) {
	input := 11
	result := FizzBuzz(input)
	if result != "11" {
		t.Errorf("expected '11' but got '%s'", result)
	}
}

func Test_input_12_expected_Fizz(t *testing.T) {
	input := 12
	result := FizzBuzz(input)
	if result != "Fizz" {
		t.Errorf("expected 'Fizz' but got '%s'", result)
	}
}

func Test_input_13_expected_13(t *testing.T) {
	input := 13
	result := FizzBuzz(input)
	if result != "13" {
		t.Errorf("expected '13' but got '%s'", result)
	}
}

func Test_input_14_expected_14(t *testing.T) {
	input := 14
	result := FizzBuzz(input)
	if result != "14" {
		t.Errorf("expected '14' but got '%s'", result)
	}
}

func Test_input_15_expected_FizzBuzz(t *testing.T) {
	input := 15
	result := FizzBuzz(input)
	if result != "FizzBuzz" {
		t.Errorf("expected 'FizzBuzz' but got '%s'", result)
	}
}
