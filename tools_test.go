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
