package tools

import (
	"testing"
)

func TestConvertByte2String(t *testing.T) {
	// Happy path tests
	t.Run("HappyPath_UTF8", func(t *testing.T) {
		input := []byte("Hello, World!")
		expected := "Hello, World!"
		result := ConvertByte2String(input, UTF8)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("HappyPath_GB18030", func(t *testing.T) {
		input := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21}
		expected := "Hello, World!"
		result := ConvertByte2String(input, GB18030)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	// Edge case tests
	t.Run("EdgeCase_EmptyInput_UTF8", func(t *testing.T) {
		input := []byte{}
		expected := ""
		result := ConvertByte2String(input, UTF8)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("EdgeCase_EmptyInput_GB18030", func(t *testing.T) {
		input := []byte{}
		expected := ""
		result := ConvertByte2String(input, GB18030)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("EdgeCase_NilInput_UTF8", func(t *testing.T) {
		var input []byte = nil
		expected := ""
		result := ConvertByte2String(input, UTF8)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("EdgeCase_NilInput_GB18030", func(t *testing.T) {
		var input []byte = nil
		expected := ""
		result := ConvertByte2String(input, GB18030)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("EdgeCase_InvalidGB18030", func(t *testing.T) {
		input := []byte{0xFF, 0xFF} // Invalid GB18030 bytes
		expected := "\uFFFD\uFFFD"  // Expected replacement characters
		result := ConvertByte2String(input, GB18030)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})
}
func TestRandomPassword(t *testing.T) {
	// Test case for length 0
	t.Run("LengthZero", func(t *testing.T) {
		result := RandomPassword(0)
		if result != "" {
			t.Errorf("Expected empty string, got %s", result)
		}
	})

	// Test case for length 1
	t.Run("LengthOne", func(t *testing.T) {
		result := RandomPassword(1)
		if len(result) != 1 {
			t.Errorf("Expected string of length 1, got %s with length %d", result, len(result))
		}
	})

	// Test case for length 10
	t.Run("LengthTen", func(t *testing.T) {
		result := RandomPassword(10)
		if len(result) != 10 {
			t.Errorf("Expected string of length 10, got %s with length %d", result, len(result))
		}
	})

	// Test case for length 100
	t.Run("LengthHundred", func(t *testing.T) {
		result := RandomPassword(100)
		if len(result) != 100 {
			t.Errorf("Expected string of length 100, got %s with length %d", result, len(result))
		}
	})

	// Test case for negative length
	t.Run("NegativeLength", func(t *testing.T) {
		result := RandomPassword(-1)
		if result != "" {
			t.Errorf("Expected empty string, got %s", result)
		}
	})

	// Test case for deterministic output
	// t.Run("DeterministicOutput", func(t *testing.T) {
	// 	rand.Seed(1) // Set a fixed seed for deterministic output
	// 	result1 := RandomPassword(10)
	// 	rand.Seed(1) // Reset the seed to ensure the same output
	// 	result2 := RandomPassword(10)
	// 	if result1 != result2 {
	// 		t.Errorf("Expected deterministic output, got different results: %s and %s", result1, result2)
	// 	}
	// })
}
