package logger

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/MaestroError/html-strings-affixer/config"
)

type Logger struct {
	data map[string][]map[string]string
	log_folder string
	folder_to_scan string
}

func (l *Logger) Init(c config.Config) {
	l.log_folder = c.Log_folder
	l.folder_to_scan = c.Folder_to_scan
	l.data = make(map[string][]map[string]string)
}

func (l *Logger) Log() bool {
	if (l.log_folder != "" || l.data !=nil) {
		l.saveFile()
		return true
	} else {
		return false
	}
}

func (l *Logger) AddNewItem(path string, element map[string]string) {
	l.data[path] = append(l.data[path], element)
}

func (l *Logger) GetLogFolder() string {
	return l.log_folder
}

func (l *Logger) GetLogData() map[string][]map[string]string {
	return l.data
}

func (l *Logger) saveFile() {
	jsonString, err := json.Marshal(l.data)
	if err != nil {
		panic(err)
	}
	
	scanFolder := strings.ReplaceAll(l.folder_to_scan, "\\", "-")
	scanFolder = strings.ReplaceAll(l.folder_to_scan, "/", "-")

	path := l.log_folder + "/" + scanFolder
	l.createDirIfNotExists(path)
	path = path + "/" + time.Now().Format("20060102150405") + ".json"
	// write file
	fileErr := ioutil.WriteFile(path, []byte(jsonString), 0)
	if fileErr != nil {
		panic(fileErr)
	}
}


func (l *Logger) createDirIfNotExists(path string) {
	exists, err := l.exists(path)
	if err != nil {
		panic(err)
	}
	if exists == false {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func (l *Logger) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}