// Here is the process for Random Forest: 1. You still have to transform your data
// 2. You still have to test for stationarity
// 3. You have to think about creating a bunch of useful features like season, time of day, t-1, t-7, t-14,
// split weeks, holidays, features that go into all machine learning models
// 4. Set up cross validation (train, test)
// 5. Optimize with gridsearch or kfold
// 6. Pick parameters, then run a model
// 7. Look at results

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var datatrain [][]float64
var vehiclestrain []float64
var datetrain []string
var vehiclestest []float64
var datetest []string
var inputTrain [][]float64
var outputTrain []float64

func main() {
	// setup and split dataset
	setupData("trafficall1.csv")

	// // transform the time series data into supervised learning
	// train := series_to_supervised(vehiclestrain)
	// fmt.Println("banyak output: ", len(train))

	// // split into input and output columns
	count := len(vehiclestrain) - 6
	train_inputs := make([][]interface{}, count)
	train_targets := make([]float64, count)

	for i := 0; i < count; i++ {
		train_inputs[i] = []interface{}{vehiclestrain[i], vehiclestrain[i+1],
			vehiclestrain[i+2], vehiclestrain[i+3], vehiclestrain[i+4], vehiclestrain[i+5]}
		train_targets[i] = vehiclestrain[i+6]
	}
	forest := BuildForest(train_inputs, train_targets, count, len(train_inputs), 1)
	// fmt.Println(forest)
	y := []interface{}{vehiclestest[0], vehiclestest[1],
		vehiclestest[2], vehiclestest[3], vehiclestest[4], vehiclestest[5]}

	fmt.Println(y, "predicted: ", forest.Predicate(y), "test: ", vehiclestest[6])

}

func setupData(file string) {
	// rand.Seed(time.Now().UTC().UnixNano())
	f, err := os.Open(file)
	if err != nil {
		return
	}
	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()
	for i := 1; i < len(csvData); i++ {
		// 	if csvData[i][4] != "Wednesday" {
		// 		continue
		// 	}
		val, _ := strconv.ParseFloat(csvData[i][2], 64)
		//don't split randomly
		if float64(i) < (float64(len(csvData)) * 0.9) {
			vehiclestrain = append(vehiclestrain, val)
			datetrain = append(datetrain, csvData[i][0])
		} else {
			vehiclestest = append(vehiclestest, val)
			datetest = append(datetest, csvData[i][0])
		}
	}
}
