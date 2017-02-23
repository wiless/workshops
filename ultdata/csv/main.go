// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to read in a CSV, create three filtered datasets, and
// save those datasets to three separate files.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Pull in the CSV file.
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	fmt.Print(irisDF)
	// Filter the dataset into three separate dataframes,
	// each corresponding to one of the Iris species.
	output := irisDF.Filter(dataframe.F{"species", "==", "Iris-setosa"})
	fw, er := os.Create("../data/setosa.csv")
	if er != nil {
		log.Panic("Error creating file ", er)
	}
	output.WriteCSV(fw)

	output = irisDF.Filter(dataframe.F{"species", "==", "Iris-virginica"})
	fw, er = os.Create("../data/virginica.csv")
	if er != nil {
		log.Panic("Error creating file ", er)
	}
	output.WriteCSV(fw)

	output = irisDF.Filter(dataframe.F{"species", "==", "Iris-versicolor"})
	fw, er = os.Create("../data/versicolor.csv")
	if er != nil {
		log.Panic("Error creating file ", er)
	}
	output.WriteCSV(fw)
	fmt.Print(output)

	// Save each of the species dataframe to a file.
}
