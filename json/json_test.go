package json

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// Test data structures
type testStruct struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Active  bool   `json:"active"`
	Numbers []int  `json:"numbers"`
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
		wantErr  bool
	}{
		{
			name:     "simple struct",
			input:    testStruct{Name: "John", Age: 30, Active: true, Numbers: []int{1, 2, 3}},
			expected: `{"name":"John","age":30,"active":true,"numbers":[1,2,3]}`,
			wantErr:  false,
		},
		{
			name:     "empty struct",
			input:    testStruct{},
			expected: `{"name":"","age":0,"active":false,"numbers":null}`,
			wantErr:  false,
		},
		{
			name:     "nil slice",
			input:    map[string]interface{}{"data": nil},
			expected: `{"data":null}`,
			wantErr:  false,
		},
		{
			name:     "string",
			input:    "hello world",
			expected: `"hello world"`,
			wantErr:  false,
		},
		{
			name:     "number",
			input:    42,
			expected: `42`,
			wantErr:  false,
		},
		{
			name:     "boolean",
			input:    true,
			expected: `true`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Marshal(tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if string(result) != tt.expected {
				t.Errorf("Marshal() = %q, expected %q", string(result), tt.expected)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   interface{}
		expected interface{}
		wantErr  bool
	}{
		{
			name:     "simple struct",
			input:    `{"name":"John","age":30,"active":true,"numbers":[1,2,3]}`,
			target:   &testStruct{},
			expected: &testStruct{Name: "John", Age: 30, Active: true, Numbers: []int{1, 2, 3}},
			wantErr:  false,
		},
		{
			name:     "empty struct",
			input:    `{"name":"","age":0,"active":false,"numbers":null}`,
			target:   &testStruct{},
			expected: &testStruct{Name: "", Age: 0, Active: false, Numbers: nil},
			wantErr:  false,
		},
		{
			name:     "string",
			input:    `"hello world"`,
			target:   new(string),
			expected: func() *string { s := "hello world"; return &s }(),
			wantErr:  false,
		},
		{
			name:     "number",
			input:    `42`,
			target:   new(int),
			expected: func() *int { i := 42; return &i }(),
			wantErr:  false,
		},
		{
			name:     "invalid json",
			input:    `{"invalid": json}`,
			target:   &testStruct{},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Unmarshal([]byte(tt.input), tt.target)
			
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if !reflect.DeepEqual(tt.target, tt.expected) {
				t.Errorf("Unmarshal() = %+v, expected %+v", tt.target, tt.expected)
			}
		})
	}
}

func TestMarshalString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
		wantErr  bool
	}{
		{
			name:     "simple struct",
			input:    testStruct{Name: "Alice", Age: 25},
			expected: `{"name":"Alice","age":25,"active":false,"numbers":null}`,
			wantErr:  false,
		},
		{
			name:     "string value",
			input:    "test string",
			expected: `"test string"`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MarshalString(tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("MarshalString() = %q, expected %q", result, tt.expected)
			}
		})
	}
}

func TestUnmarshalString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   interface{}
		expected interface{}
		wantErr  bool
	}{
		{
			name:     "simple struct",
			input:    `{"name":"Bob","age":35}`,
			target:   &testStruct{},
			expected: &testStruct{Name: "Bob", Age: 35},
			wantErr:  false,
		},
		{
			name:     "invalid json",
			input:    `{invalid}`,
			target:   &testStruct{},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UnmarshalString(tt.input, tt.target)
			
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if !reflect.DeepEqual(tt.target, tt.expected) {
				t.Errorf("UnmarshalString() = %+v, expected %+v", tt.target, tt.expected)
			}
		})
	}
}

func TestNewEncoder(t *testing.T) {
	buf := &bytes.Buffer{}
	encoder := NewEncoder(buf)
	
	if encoder == nil {
		t.Error("NewEncoder() returned nil")
	}
	
	// Test encoding
	testData := testStruct{Name: "Test", Age: 20}
	err := encoder.Encode(testData)
	if err != nil {
		t.Errorf("Encoder.Encode() error: %v", err)
	}
	
	// Verify output
	output := buf.String()
	if !strings.Contains(output, `"name":"Test"`) {
		t.Errorf("Encoded output doesn't contain expected data: %s", output)
	}
}

func TestNewDecoder(t *testing.T) {
	input := `{"name":"Test","age":25}`
	reader := strings.NewReader(input)
	decoder := NewDecoder(reader)
	
	if decoder == nil {
		t.Error("NewDecoder() returned nil")
	}
	
	// Test decoding
	var result testStruct
	err := decoder.Decode(&result)
	if err != nil {
		t.Errorf("Decoder.Decode() error: %v", err)
	}
	
	// Verify result
	expected := testStruct{Name: "Test", Age: 25}
	if result.Name != expected.Name || result.Age != expected.Age {
		t.Errorf("Decoded result = %+v, expected %+v", result, expected)
	}
}

func TestMustMarshal(t *testing.T) {
	// Test normal case
	data := testStruct{Name: "Test", Age: 30}
	result := MustMarshal(data)
	
	if len(result) == 0 {
		t.Error("MustMarshal() returned empty result")
	}
	
	// Verify it contains expected data
	if !strings.Contains(string(result), `"name":"Test"`) {
		t.Errorf("MustMarshal() result doesn't contain expected data: %s", string(result))
	}
}

func TestMustMarshalPanic(t *testing.T) {
	// Test that invalid data causes panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustMarshal() should have panicked with invalid data")
		}
	}()
	
	// Use a channel which cannot be marshaled to JSON
	invalidData := make(chan int)
	MustMarshal(invalidData)
}

func TestMustMarshalString(t *testing.T) {
	// Test normal case
	data := testStruct{Name: "Test", Age: 30}
	result := MustMarshalString(data)
	
	if result == "" {
		t.Error("MustMarshalString() returned empty result")
	}
	
	// Verify it contains expected data
	if !strings.Contains(result, `"name":"Test"`) {
		t.Errorf("MustMarshalString() result doesn't contain expected data: %s", result)
	}
}

func TestMustMarshalStringPanic(t *testing.T) {
	// Test that invalid data causes panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustMarshalString() should have panicked with invalid data")
		}
	}()
	
	// Use a channel which cannot be marshaled to JSON
	invalidData := make(chan int)
	MustMarshalString(invalidData)
}

func TestIndent(t *testing.T) {
	// Test JSON indentation
	input := `{"name":"John","age":30}`
	var buf bytes.Buffer
	
	err := Indent(&buf, []byte(input), "", "  ")
	if err != nil {
		t.Errorf("Indent() error: %v", err)
	}
	
	result := buf.String()
	expected := "{\n  \"name\": \"John\",\n  \"age\": 30\n}"
	
	if result != expected {
		t.Errorf("Indent() result = %q, expected %q", result, expected)
	}
}

func TestIndentWithPrefix(t *testing.T) {
	// Test JSON indentation with prefix
	input := `{"key":"value"}`
	var buf bytes.Buffer
	
	err := Indent(&buf, []byte(input), "> ", "  ")
	if err != nil {
		t.Errorf("Indent() error: %v", err)
	}
	
	result := buf.String()
	if !strings.Contains(result, "> ") {
		t.Errorf("Indent() result doesn't contain prefix: %q", result)
	}
}

func TestIndentInvalidJSON(t *testing.T) {
	// Test with invalid JSON
	input := `{invalid json}`
	var buf bytes.Buffer
	
	err := Indent(&buf, []byte(input), "", "  ")
	if err == nil {
		t.Error("Indent() should have returned error for invalid JSON")
	}
}

// File operation tests
func TestMarshalToFile(t *testing.T) {
	// Create temp file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test.json")
	
	// Test data
	data := testStruct{Name: "FileTest", Age: 40, Active: true}
	
	// Marshal to file
	err := MarshalToFile(filename, data)
	if err != nil {
		t.Errorf("MarshalToFile() error: %v", err)
	}
	
	// Verify file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("MarshalToFile() did not create file")
	}
	
	// Read and verify content
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read created file: %v", err)
	}
	
	if !strings.Contains(string(content), `"name":"FileTest"`) {
		t.Errorf("File content doesn't contain expected data: %s", string(content))
	}
}

func TestUnmarshalFromFile(t *testing.T) {
	// Create temp file with JSON content
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test.json")
	content := `{"name":"FileTest","age":40,"active":true,"numbers":[1,2,3]}`
	
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// Unmarshal from file
	var result testStruct
	err = UnmarshalFromFile(filename, &result)
	if err != nil {
		t.Errorf("UnmarshalFromFile() error: %v", err)
	}
	
	// Verify result
	expected := testStruct{Name: "FileTest", Age: 40, Active: true, Numbers: []int{1, 2, 3}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("UnmarshalFromFile() result = %+v, expected %+v", result, expected)
	}
}

func TestUnmarshalFromFileNotFound(t *testing.T) {
	// Test with non-existent file
	var result testStruct
	err := UnmarshalFromFile("/non/existent/file.json", &result)
	if err == nil {
		t.Error("UnmarshalFromFile() should have returned error for non-existent file")
	}
}

func TestMustMarshalToFile(t *testing.T) {
	// Create temp file
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "must_test.json")
	
	// Test data
	data := testStruct{Name: "MustFileTest", Age: 35}
	
	// This should not panic
	MustMarshalToFile(filename, data)
	
	// Verify file exists and has correct content
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read created file: %v", err)
	}
	
	if !strings.Contains(string(content), `"name":"MustFileTest"`) {
		t.Errorf("File content doesn't contain expected data: %s", string(content))
	}
}

func TestMustMarshalToFilePanic(t *testing.T) {
	// Test that invalid filename causes panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustMarshalToFile() should have panicked with invalid filename")
		}
	}()
	
	// Use invalid filename (directory that doesn't exist)
	MustMarshalToFile("/invalid/path/file.json", testStruct{})
}

func TestMustUnmarshalFromFile(t *testing.T) {
	// Create temp file with JSON content
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "must_test.json")
	content := `{"name":"MustFileTest","age":45}`
	
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// This should not panic
	var result testStruct
	MustUnmarshalFromFile(filename, &result)
	
	// Verify result
	if result.Name != "MustFileTest" || result.Age != 45 {
		t.Errorf("MustUnmarshalFromFile() result = %+v, expected name=MustFileTest, age=45", result)
	}
}

func TestMustUnmarshalFromFilePanic(t *testing.T) {
	// Test that non-existent file causes panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustUnmarshalFromFile() should have panicked with non-existent file")
		}
	}()
	
	var result testStruct
	MustUnmarshalFromFile("/non/existent/file.json", &result)
}

// Benchmark tests
func BenchmarkMarshal(b *testing.B) {
	data := testStruct{Name: "Benchmark", Age: 30, Active: true, Numbers: []int{1, 2, 3, 4, 5}}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Marshal(data)
		if err != nil {
			b.Fatalf("Marshal error: %v", err)
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	input := `{"name":"Benchmark","age":30,"active":true,"numbers":[1,2,3,4,5]}`
	data := []byte(input)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result testStruct
		err := Unmarshal(data, &result)
		if err != nil {
			b.Fatalf("Unmarshal error: %v", err)
		}
	}
}

func BenchmarkMarshalString(b *testing.B) {
	data := testStruct{Name: "Benchmark", Age: 30}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := MarshalString(data)
		if err != nil {
			b.Fatalf("MarshalString error: %v", err)
		}
	}
}

func BenchmarkUnmarshalString(b *testing.B) {
	input := `{"name":"Benchmark","age":30}`
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result testStruct
		err := UnmarshalString(input, &result)
		if err != nil {
			b.Fatalf("UnmarshalString error: %v", err)
		}
	}
}

// Additional edge case tests
func TestEdgeCases(t *testing.T) {
	t.Run("empty byte slice", func(t *testing.T) {
		var result interface{}
		err := Unmarshal([]byte("null"), &result)
		if err != nil {
			t.Errorf("Unmarshal empty byte slice error: %v", err)
		}
	})
	
	t.Run("empty string", func(t *testing.T) {
		result, err := MarshalString("")
		if err != nil {
			t.Errorf("MarshalString empty string error: %v", err)
		}
		if result != `""` {
			t.Errorf("MarshalString empty string = %q, expected %q", result, `""`)
		}
	})
	
	t.Run("encoder with nil writer", func(t *testing.T) {
		// This should not panic, though it may not be very useful
		encoder := NewEncoder(io.Discard)
		if encoder == nil {
			t.Error("NewEncoder with discard writer returned nil")
		}
	})
}

// Test type compatibility
func TestTypeCompatibility(t *testing.T) {
	// Test that our encoder/decoder interfaces work with expected types
	buf := &bytes.Buffer{}
	encoder := NewEncoder(buf)
	
	// Test encoding various types
	testCases := []interface{}{
		42,
		"string",
		true,
		[]int{1, 2, 3},
		map[string]interface{}{"key": "value"},
		testStruct{Name: "test"},
	}
	
	for _, tc := range testCases {
		err := encoder.Encode(tc)
		if err != nil {
			t.Errorf("Encoder failed for type %T: %v", tc, err)
		}
	}
}