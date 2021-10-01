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
	"time"

	"github.com/fxsjy/RF.go/RF/Regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var (
	vehiclestrain []float64
	datetrain     []float64
	vehiclestest  []float64
	datetest      []float64
	forecast      []float64
)

func main() {
	// setup and split dataset
	setupData("Metro_Interstate_Traffic_Volume.csv")

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
	forest := Regression.BuildForest(train_inputs, train_targets, count, len(train_inputs), 1)
	// fmt.Println(forest)

	//testing
	y := []interface{}{vehiclestest[0], vehiclestest[1],
		vehiclestest[2], vehiclestest[3], vehiclestest[4], vehiclestest[5]}

	fmt.Println(y, "predicted: ", forest.Predicate(y), "test: ", vehiclestest[46])

	count2 := len(vehiclestest) - 6
	for i := 0; i < count2; i++ {
		predict := []interface{}{vehiclestest[i], vehiclestest[i+1],
			vehiclestest[i+2], vehiclestest[i+3], vehiclestest[i+4], vehiclestest[i+5]}
		forecast = append(forecast, forest.Predicate(predict))

	}

	p := plot.New()
	p.Title.Text = "Traffic Volume Forecasting"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"Train", makePoints(vehiclestrain, datetrain),
		"Test", makePoints(vehiclestest, datetest),
		"Predict", makePoints(forecast, datetest[6:]))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}
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
		layouts := "2006-01-02 15:04:05"
		t, err := time.Parse(layouts, csvData[i][7])
		if err != nil {
			fmt.Println(err)
		}

		val, _ := strconv.ParseFloat(csvData[i][8], 64)
		//don't split randomly
		if float64(i) < (float64(len(csvData)) * 0.9) {
			vehiclestrain = append(vehiclestrain, val)
			datetrain = append(datetrain, float64(t.Unix()))
		} else {
			vehiclestest = append(vehiclestest, val)
			datetest = append(datetest, float64(t.Unix()))
		}
	}
}

func makePoints(data []float64, date []float64) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = date[i]
		pts[i].Y = data[i]
	}
	return pts
}
