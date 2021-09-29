package main

// //ubah data
func series_to_supervised(data []float64) [][]float64 {
	// fmt.Println("banyak input: ", len(data))
	count := len(data) - 6
	for i := 0; i < count; i++ {
		// board := []float64{data[i : i+5]}

		datatrain = append(datatrain, data[i:i+6])
	}
	return datatrain
}
