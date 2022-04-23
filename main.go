package main

import (
	data_conversion "github.com/joffref/Projet-MRH/pkg/data-conversion"
	"github.com/joffref/Projet-MRH/pkg/git"
	"github.com/joffref/Projet-MRH/utils"
)

func main() {
	logger := utils.NewLogger()
	git.CloneRepo(logger)
	logger.Debug("Cloned el4000")
	data_conversion.Convert(logger)
	git.RemoveRepo(logger)
}