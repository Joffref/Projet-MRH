package main

import (
	"fmt"
	data_conversion "github.com/joffref/Projet-MRH/pkg/data-conversion"
	"github.com/joffref/Projet-MRH/pkg/git"
	"github.com/joffref/Projet-MRH/pkg/plotting"
	"github.com/joffref/Projet-MRH/utils"
	"os"
)

func main() {
	logger := utils.NewLogger()
	git.CloneRepo(logger)
	logger.Debug("Cloned el4000")
	wd, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	entries, err := os.ReadDir(wd + "/data/")
	if err != nil {
		logger.Fatal(err)
	}
	for _, file := range entries {
		data_conversion.Convert(logger, file.Name())
		plotting.Graph(logger, file.Name())
		data_conversion.RemoveData(logger)
	}
	fmt.Println(entries[0].Name())
	git.RemoveRepo(logger)
}
