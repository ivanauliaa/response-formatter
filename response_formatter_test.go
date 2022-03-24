package formatter

import (
	"net/http"
	"reflect"
	"testing"
)

const (
	dummyStringInputData = "Hello, World!"
	dummyErrorInputData  = "Something went wrong."
)

func TestResponseFormatter(t *testing.T) {
	testCases := []struct {
		name         string
		inputStatus  int32
		inputMessage string
		inputData    interface{}
		expected     responseFormat
	}{
		{
			name:         "it should return valid responseFormat object, according to the input",
			inputStatus:  http.StatusCreated,
			inputMessage: successMessage,
			inputData:    dummyStringInputData,
			expected: responseFormat{
				Status:  http.StatusCreated,
				Message: successMessage,
				Data:    dummyStringInputData,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := ResponseFormatter(testCase.inputStatus, testCase.inputMessage, testCase.inputData)

			assertResponseFormat(t, got, testCase.expected)
		})
	}
}

func TestBadRequestResponse(t *testing.T) {
	testCases := []struct {
		name      string
		inputData interface{}
		expected  responseFormat
	}{
		{
			name:      "it should return valid responseFormat object, according to the input",
			inputData: dummyErrorInputData,
			expected: responseFormat{
				Status:  http.StatusBadRequest,
				Message: failMessage,
				Data:    dummyErrorInputData,
			},
		},
	}

	for _, testCase := range testCases {
		got := BadRequestResponse(testCase.inputData)

		assertResponseFormat(t, got, testCase.expected)
	}
}

func TestNotFoundResponse(t *testing.T) {
	testCases := []struct {
		name      string
		inputData interface{}
		expected  responseFormat
	}{
		{
			name:      "it should return valid responseFormat object, according to the input",
			inputData: dummyErrorInputData,
			expected: responseFormat{
				Status:  http.StatusNotFound,
				Message: failMessage,
				Data:    dummyErrorInputData,
			},
		},
	}

	for _, testCase := range testCases {
		got := NotFoundResponse(testCase.inputData)

		assertResponseFormat(t, got, testCase.expected)
	}
}

func TestUnauthorizedResponse(t *testing.T) {
	testCases := []struct {
		name      string
		inputData interface{}
		expected  responseFormat
	}{
		{
			name:      "it should return valid responseFormat object, according to the input",
			inputData: dummyErrorInputData,
			expected: responseFormat{
				Status:  http.StatusUnauthorized,
				Message: failMessage,
				Data:    dummyErrorInputData,
			},
		},
	}

	for _, testCase := range testCases {
		got := UnauthorizedResponse(testCase.inputData)

		assertResponseFormat(t, got, testCase.expected)
	}
}

func TestInternalServerErrorResponse(t *testing.T) {
	testCases := []struct {
		name      string
		inputData interface{}
		expected  responseFormat
	}{
		{
			name:      "it should return valid responseFormat object, according to the input",
			inputData: dummyErrorInputData,
			expected: responseFormat{
				Status:  http.StatusInternalServerError,
				Message: failMessage,
				Data:    dummyErrorInputData,
			},
		},
	}

	for _, testCase := range testCases {
		got := InternalServerErrorResponse(testCase.inputData)

		assertResponseFormat(t, got, testCase.expected)
	}
}

func TestSuccessResponse(t *testing.T) {
	testCases := []struct {
		name      string
		inputData interface{}
		expected  responseFormat
	}{
		{
			name:      "it should return valid responseFormat object, according to the input",
			inputData: dummyStringInputData,
			expected: responseFormat{
				Status:  http.StatusOK,
				Message: successMessage,
				Data:    dummyStringInputData,
			},
		},
	}

	for _, testCase := range testCases {
		got := SuccessResponse(testCase.inputData)

		assertResponseFormat(t, got, testCase.expected)
	}
}

func assertResponseFormat(t testing.TB, got, want responseFormat) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
