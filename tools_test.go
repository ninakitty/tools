package tools

import (
	"strings"
	"testing"

	"golang.org/x/exp/rand"
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
	// Test for length 0
	result := RandomPassword(0)
	if result != "" {
		t.Errorf("Expected empty string for length 0, but got %s", result)
	}

	// Test for negative length
	result = RandomPassword(-5)
	if result != "" {
		t.Errorf("Expected empty string for negative length, but got %s", result)
	}

	// Test for positive length
	length := 10
	result = RandomPassword(length)
	if len(result) != length {
		t.Errorf("Expected password of length %d, but got %d", length, len(result))
	}

	// Test for all characters in the generated password are from the allowed set
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	result = RandomPassword(length)
	for _, char := range result {
		if !strings.Contains(letters, string(char)) {
			t.Errorf("Generated password contains invalid character: %c", char)
		}
	}

	// Test for randomness
	rand.Seed(1) // Setting seed for reproducibility
	password1 := RandomPassword(length)
	rand.Seed(1)
	password2 := RandomPassword(length)
	if password1 != password2 {
		t.Errorf("Expected same passwords due to same seed, but got %s and %s", password1, password2)
	}

	// Test for different seeds
	rand.Seed(2)
	password2 = RandomPassword(length)
	if password1 == password2 {
		t.Errorf("Expected different passwords due to different seeds, but got %s and %s", password1, password2)
	}
}
