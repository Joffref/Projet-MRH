package data_conversion

import (
	"bufio"
	"bytes"
	"github.com/joffref/Projet-MRH/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
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
	for _, s := range strings.Split(data, " ") {
		if v, _ := regexp.MatchString("2*", s); v == true {
			sanitizedData = append(sanitizedData, s)
		}
	}
	return sanitizedData
}

func pushDataInCSVFile(logger *log.Logger, data []string) {
	logger.Info("pushData")
	file, err := os.OpenFile(utils.Path+"/data/data.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		logger.Error(err)
	}
	defer file.Close()
	for _, s := range data {
		err = file.Sync()
		if err != nil {
			logger.Error(err)
		}
		w := bufio.NewWriter(file)
		_, err := w.WriteString(s)
		if err != nil {
			logger.Error(err)
		}
	}
}
