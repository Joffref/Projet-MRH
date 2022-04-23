package data_conversion

import (
	"bytes"
	"github.com/joffref/Projet-MRH/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
)

func Convert(logger *log.Logger) {
	log.Info("Convert")
	data := batchConvert(logger)
	pushDataInCSVFile(logger, data)
}

func batchConvert(logger *log.Logger) []string {
	logger.Info("BatchConvert")
	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	}
	entries, err := ioutil.ReadDir(path + "/data/")
	if err != nil {
		logger.Error(err)
	}
	var data bytes.Buffer
	var dataStringFormat []string
	for _, entry := range entries {
		cmd := exec.Command(
			"python3",
			path+utils.Path+"/el4000.py",
			"-p",
			"csv",
			"-d",
			"';'",
			"-v",
			"--data-only",
			path+"/data/"+entry.Name(),
		)
		cmd.Stdout = io.MultiWriter(logger.Writer(), &data)
		cmd.Stderr = logger.Writer()
		dataStringFormat = append(dataStringFormat, sanitizeData(logger, data.String())...)
		err = cmd.Run()
		if err != nil {
			logger.Error(err)
		}
	}
	return dataStringFormat
}

func sanitizeData(logger *log.Logger, data string) []string {
	logger.Info("sanitizeData")
	var sanitizedData []string
	for _, s := range strings.Split(data, "\n") {
		if v, _ := regexp.MatchString("^2", s); v == true {
			sanitizedData = append(sanitizedData, s)
		}
	}
	return sanitizedData
}

func pushDataInCSVFile(logger *log.Logger, data []string) {
	logger.Info("pushData")
	logger.Info("data: ", reflect.TypeOf(data))
	pwd, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	}
	file, err := os.OpenFile(pwd+"/data/data.csv", os.O_RDWR|os.O_CREATE, 0655)
	if err != nil {
		logger.Error(err)
	}
	defer file.Close()
	for _, s := range data {
		_, err = file.WriteString(s + "\n")
		if err != nil {
			logger.Error(err)
		}
	}
}

func RemoveData(logger *log.Logger) {
	logger.Info("RemoveData")
	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	}
	err = os.Remove(path + "/data/data.csv")
	if err != nil {
		logger.Error(err)
	}
}
