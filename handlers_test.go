package main

import (
    "net/http"
    "net/http/httptest"
	"testing"
	"bytes"
	"reflect"
)

//Function to test calculateMin method
func TestCalculateMin(t *testing.T) {
	//Case 1: Happy Scenario
	var input1 = []byte(`{"numList":[1,5,3], "quantifier": 2}`)
	req := httptest.NewRequest(http.MethodPost, "/min", bytes.NewBuffer(input1))
	res := httptest.NewRecorder()

	calculateMin(res, req)
	if reflect.DeepEqual(res.Code, http.StatusOK) {
		t.Log("Test Case:1 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	//Case 2: Incomplete Parameters
	var input2 = []byte(`{"numList":[], "quantifier": 2}`)
	req = httptest.NewRequest(http.MethodPost, "/min", bytes.NewBuffer(input2))
	res = httptest.NewRecorder()
	expectedOutput := "Incomplete input request parameters.\n"
	calculateMin(res, req)
	actualOutput := res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:2 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:2 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

	//Case 3: Invalid input request
	var input3 = []byte(`{"numList":["test"], "quantifier": 2}`)
	req = httptest.NewRequest(http.MethodPost, "/min", bytes.NewBuffer(input3))
	res = httptest.NewRecorder()
	expectedOutput = "Invalid input request.\n"
	calculateMin(res, req)
	actualOutput = res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:3 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:3 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

	//Case 4: Quantifier greater than list of numbers
	var input4 = []byte(`{"numList":[1], "quantifier": 2}`)
	req = httptest.NewRequest(http.MethodPost, "/min", bytes.NewBuffer(input4))
	res = httptest.NewRecorder()
	expectedOutput = "Quantifier must be less than equal to the length of list of numbers.\n"
	calculateMin(res, req)
	actualOutput = res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:4 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:4 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

}

//Function to test calculateMax method
func TestCalculateMax(t *testing.T) {
	//Case 1: Happy Scenario
	var input1 = []byte(`{"numList":[1,5,3], "quantifier": 2}`)
	req := httptest.NewRequest(http.MethodPost, "/max", bytes.NewBuffer(input1))
	res := httptest.NewRecorder()

	calculateMax(res, req)
	if reflect.DeepEqual(res.Code, http.StatusOK) {
		t.Log("Test Case:1 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	//Case 2: Incomplete Parameters
	var input2 = []byte(`{"numList":[], "quantifier": 2}`)
	req = httptest.NewRequest(http.MethodPost, "/max", bytes.NewBuffer(input2))
	res = httptest.NewRecorder()
	expectedOutput := "Incomplete input request parameters.\n"
	calculateMax(res, req)
	actualOutput := res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:2 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:2 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

	//Case 3: Invalid input request
	var input3 = []byte(`{"numList":["test"], "quantifier": 2}`)
	req = httptest.NewRequest(http.MethodPost, "/max", bytes.NewBuffer(input3))
	res = httptest.NewRecorder()
	expectedOutput = "Invalid input request.\n"
	calculateMax(res, req)
	actualOutput = res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:3 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:3 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

	//Case 4: Quantifier greater than list of numbers
	var input4 = []byte(`{"numList":[1], "quantifier": 2}`)
	req = httptest.NewRequest(http.MethodPost, "/max", bytes.NewBuffer(input4))
	res = httptest.NewRecorder()
	expectedOutput = "Quantifier must be less than equal to the length of list of numbers.\n"
	calculateMax(res, req)
	actualOutput = res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:4 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:4 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

}

//Function to test calculateAvg method
func TestCalculateAvg(t *testing.T) {
	//Case 1: Happy Scenario
	var input = []byte(`{"numList":[1,5,3]}`)
	req := httptest.NewRequest(http.MethodPost, "/avg", bytes.NewBuffer(input))
	res := httptest.NewRecorder()

	calculateAvg(res, req)
	if reflect.DeepEqual(res.Code, http.StatusOK) {
		t.Log("Test Case:1 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	//Case 2: Incomplete Parameters
	var input2 = []byte(`{"numList":[]}`)
	req = httptest.NewRequest(http.MethodPost, "/avg", bytes.NewBuffer(input2))
	res = httptest.NewRecorder()
	expectedOutput := "Incomplete input request parameters.\n"
	calculateAvg(res, req)
	actualOutput := res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:2 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:2 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

	//Case 3: Invalid input request
	var input3 = []byte(`{"numList":["test"]}`)
	req = httptest.NewRequest(http.MethodPost, "/avg", bytes.NewBuffer(input3))
	res = httptest.NewRecorder()
	expectedOutput = "Invalid input request.\n"
	calculateAvg(res, req)
	actualOutput = res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:3 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:3 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

}

//Function to test calculateMedian method
func TestCalculateMedian(t *testing.T) {
	//Case 1.1: Happy Scenario In Odd numbers list
	var input = []byte(`{"numList":[1,5,3]}`)
	req := httptest.NewRequest(http.MethodPost, "/median", bytes.NewBuffer(input))
	res := httptest.NewRecorder()

	calculateMedian(res, req)
	if reflect.DeepEqual(res.Code, http.StatusOK) {
		t.Log("Test Case:1.1 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	//Case 1.2: Happy Scenario In Even numbers list
	var input1 = []byte(`{"numList":[1,5]}`)
	req = httptest.NewRequest(http.MethodPost, "/median", bytes.NewBuffer(input1))
	res = httptest.NewRecorder()

	calculateMedian(res, req)
	if reflect.DeepEqual(res.Code, http.StatusOK) {
		t.Log("Test Case:1.2 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	//Case 2: Incomplete Parameters
	var input2 = []byte(`{"numList":[]}`)
	req = httptest.NewRequest(http.MethodPost, "/Median", bytes.NewBuffer(input2))
	res = httptest.NewRecorder()
	expectedOutput := "Incomplete input request parameters.\n"
	calculateMedian(res, req)
	actualOutput := res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:2 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:2 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

	//Case 3: Invalid input request
	var input3 = []byte(`{"numList":["test"]}`)
	req = httptest.NewRequest(http.MethodPost, "/median", bytes.NewBuffer(input3))
	res = httptest.NewRecorder()
	expectedOutput = "Invalid input request.\n"
	calculateMedian(res, req)
	actualOutput = res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:3 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:3 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

}

//Function to test calculatePercentile method
func TestCalculatePercentile(t *testing.T) {
	//Case 1: Happy Scenario
	var input1 = []byte(`{"numList":[1,5,3], "quantifier": 2}`)
	req := httptest.NewRequest(http.MethodPost, "/percentile", bytes.NewBuffer(input1))
	res := httptest.NewRecorder()

	calculatePercentile(res, req)
	if reflect.DeepEqual(res.Code, http.StatusOK) {
		t.Log("Test Case:1 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	//Case 2: Incomplete Parameters
	var input2 = []byte(`{"numList":[], "quantifier": 2}`)
	req = httptest.NewRequest(http.MethodPost, "/percentile", bytes.NewBuffer(input2))
	res = httptest.NewRecorder()
	expectedOutput := "Incomplete input request parameters.\n"
	calculatePercentile(res, req)
	actualOutput := res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:2 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:2 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

	//Case 3: Invalid input request
	var input3 = []byte(`{"numList":["test"], "quantifier": 2}`)
	req = httptest.NewRequest(http.MethodPost, "/percentile", bytes.NewBuffer(input3))
	res = httptest.NewRecorder()
	expectedOutput = "Invalid input request.\n"
	calculatePercentile(res, req)
	actualOutput = res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:3 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:3 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

	//Case 4: Quantifier greater than 100th percentile
	var input4 = []byte(`{"numList":[1], "quantifier": 101}`)
	req = httptest.NewRequest(http.MethodPost, "/percentile", bytes.NewBuffer(input4))
	res = httptest.NewRecorder()
	expectedOutput = "Quantifier must be less than equal to the 100th percentile.\n"
	calculatePercentile(res, req)
	actualOutput = res.Body.String()
	if reflect.DeepEqual(res.Code, http.StatusBadRequest) {
		t.Log("Test Case:4 StatusCode Passed ")
	} else {
		t.Errorf("Unexpected result, got status %d but wanted %d", res.Code, http.StatusOK)
	}

	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:4 Response Body Passed")
	} else {
		t.Errorf("Unexpected result, got %v but wanted %v", actualOutput, expectedOutput)
	}

}