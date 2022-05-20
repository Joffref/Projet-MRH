package plotting

import (
	"bytes"
	"encoding/csv"
	"github.com/Arafatk/glot"
	"github.com/joffref/Projet-MRH/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func readData() ([][]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(pwd + "/data/data.csv")
	if err != nil {
		return nil, err
	}
	return csv.NewReader(bytes.NewReader(data)).ReadAll()
}

func Graph(logger *log.Logger, file string) {
	var dates []float64
	var values []float64
	data, err := readData()
	for _, row := range data {
		for _, col := range row {
			date, err := time.Parse(utils.DateFormat, strings.Replace(strings.Split(col, ";")[0], "'", "", -1))
			if err != nil {
				logger.Error(err)
			}
			dates = append(dates, float64(date.Unix()))
			value, err := strconv.ParseFloat(strings.Replace(strings.Split(col, ";")[1], "'", "", -1), 64)
			if err != nil {
				logger.Error(err)
			}
			values = append(values, value)
		}
	}
	if err != nil {
		logger.Fatal(err)
	}
	var dataSet [][]float64
	dataSet = append(dataSet, dates)
	dataSet = append(dataSet, values)
	plot, _ := glot.NewPlot(2, true, true)
	plot.AddPointGroup("data", "points", dataSet)
	plot.SavePlot("figures/" + file + ".png")
}
