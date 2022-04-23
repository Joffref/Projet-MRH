package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/joffref/Projet-MRH/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func CloneRepo(logger *log.Logger) {
	workingDir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("Working dir: ", workingDir)
	logger.Info("Cloning repo")
	path := workingDir + utils.Path
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:      utils.URI,
		Progress: os.Stdout,
	})
	if err != nil {
		logger.Fatal(err)
	}
}

func RemoveRepo(logger *log.Logger) {
	workingDir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("Working dir: ", workingDir)
	logger.Info("Removing repo")
	path := workingDir + utils.Path
	err = os.RemoveAll(path)
	if err != nil {
		logger.Fatal(err)
	}
}
