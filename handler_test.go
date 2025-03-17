package lab2

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedError  string
		converterMock  func(string) (string, error)
	}{
		{
			name:           "Valid expression",
			input:          "3 4 +",
			expectedOutput: "(3 + 4)",
			expectedError:  "",
			converterMock: func(input string) (string, error) {
				if input == "3 4 +" {
					return "(3 + 4)", nil
				}
				return "", errors.New("unexpected input for mock")
			},
		},
		{
			name:           "Invalid expression - converter error",
			input:          "3 +",
			expectedOutput: "",
			expectedError:  "expression conversion error: invalid expression: not enough operands",
			converterMock: func(input string) (string, error) {
				if input == "3 +" {
					return "", errors.New("invalid expression: not enough operands")
				}
				return "", errors.New("unexpected input for mock")
			},
		},
		{
			name:           "Empty input",
			input:          "",
			expectedOutput: "",
			expectedError:  "empty input expression",
			converterMock: func(input string) (string, error) {
				return "", errors.New("this should not be called")
			},
		},
		{
			name:           "Whitespace input",
			input:          "   ",
			expectedOutput: "",
			expectedError:  "empty input expression",
			converterMock: func(input string) (string, error) {
				return "", errors.New("this should not be called")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inputReader := strings.NewReader(test.input)
			outputBuffer := &bytes.Buffer{}

			handler := ComputeHandler{
				Input:     inputReader,
				Output:    outputBuffer,
				Converter: test.converterMock,
			}

			err := handler.Compute()

			if test.expectedError != "" {
				assert.Error(t, err, "Expected an error but got nil")
				assert.EqualError(t, err, test.expectedError, "Error message mismatch")
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.Equal(t, test.expectedOutput, outputBuffer.String(), "Output mismatch")
			}

			if test.expectedError != "" {
				assert.Empty(t, outputBuffer.String(), "Output should be empty when an error occurs")
			}
		})
	}
}
