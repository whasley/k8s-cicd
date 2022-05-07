package main

import (
	"net/http"
	"encoding/json"
	"sort"
	"fmt"
	"math"
)

//Struct to store the input request
type Input struct {
	NumList []float64 `json:"numList"`
	Quantifier int	`json:"quantifier,omitempty"`
}

//Function to get the median of float slice
func getMedian(data []float64) float64 {
    dataCopy := make([]float64, len(data))
    copy(dataCopy, data)

    sort.Float64s(dataCopy)

    var median float64
    l := len(dataCopy)
    if l == 0 {
        return 0
    } else if l%2 == 0 {
        median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
    } else {
        median = dataCopy[l/2]
    }

    return median
}

//Function to handle "/min" endpoint and provide response to the user
func calculateMin(w http.ResponseWriter, r *http.Request){
	var input Input

	//Fetching request Parameters
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input request.", http.StatusBadRequest)
		return
	}

	//Validating Input parameters
	if (len(input.NumList) == 0) || (input.Quantifier < 1) {
		http.Error(w, "Incomplete input request parameters.", http.StatusBadRequest)
		return
	} 

	//Calculating minimum number(s) from the list
	sort.Slice(input.NumList, func(i, j int) bool {
		return input.NumList[i] < input.NumList[j]
	})
	if input.Quantifier > len(input.NumList){
		http.Error(w, "Quantifier must be less than equal to the length of list of numbers.", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", input.NumList[:(input.Quantifier)])
	return
}

//Function to handle "/max" endpoint and provide response to the user
func calculateMax(w http.ResponseWriter, r *http.Request){
	var input Input

	//Fetching request Parameters
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input request.", http.StatusBadRequest)
		return
	}

	//Validating Input parameters
	if (len(input.NumList) == 0) || (input.Quantifier < 1) {
		http.Error(w, "Incomplete input request parameters.", http.StatusBadRequest)
		return
	} 

	//Calculating maximum number(s) from the list
	sort.Slice(input.NumList, func(i, j int) bool {
		return input.NumList[i] > input.NumList[j]
	})
	if input.Quantifier > len(input.NumList){
		http.Error(w, "Quantifier must be less than equal to the length of list of numbers.", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", input.NumList[:(input.Quantifier)])
	return
}

//Function to handle "/avg" endpoint and provide response to the user
func calculateAvg(w http.ResponseWriter, r *http.Request){
	var input Input

	//Fetching request Parameters
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input request.", http.StatusBadRequest)
		return
	}

	//Validating Input parameters
	if (len(input.NumList) == 0) {
		http.Error(w, "Incomplete input request parameters.", http.StatusBadRequest)
		return
	} 

	//Calculating average of the numbers
	sum := 0.0
	for _, i := range input.NumList {
		sum += i
	}
	avg := sum/float64(len(input.NumList))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", avg)
	return
}

//Function to handle "/median" endpoint and provide response to the user
func calculateMedian(w http.ResponseWriter, r *http.Request){
	var input Input

	//Fetching request Parameters
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input request.", http.StatusBadRequest)
		return
	}

	//Validating Input parameters
	if (len(input.NumList) == 0) {
		http.Error(w, "Incomplete input request parameters.", http.StatusBadRequest)
		return
	} 
	
	//Calculating median of the numbers
	median := getMedian(input.NumList)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", median)
	return
}

//Function to handle "/percentile" endpoint and provide response to the user
func calculatePercentile(w http.ResponseWriter, r *http.Request){
	var input Input

	//Fetching request Parameters
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input request.", http.StatusBadRequest)
		return
	}

	//Validating Input parameters
	if (len(input.NumList) == 0) || (input.Quantifier <=0) {
		http.Error(w, "Incomplete input request parameters.", http.StatusBadRequest)
		return
	} 

	//Calculating percentile from the list
	sort.Slice(input.NumList, func(i, j int) bool {
		return input.NumList[i] < input.NumList[j]
	})
	if input.Quantifier > 100{
		http.Error(w, "Quantifier must be less than equal to the 100th percentile.", http.StatusBadRequest)
		return
	}
	ordinalRank := int(math.Round((float64(input.Quantifier)/100.0) * float64(len(input.NumList))))	
	w.WriteHeader(http.StatusOK)
	if ordinalRank > 1 {
		fmt.Fprintf(w, "%v", input.NumList[ordinalRank-1])
	} else {
		fmt.Fprintf(w, "%v", input.NumList[0])
	}	
	return
}