package logger

import (
	"fmt"

	"github.com/MaestroError/html-strings-affixer/config"
)

type Logger struct {
	data map[string][]map[string]string
	log_folder string
}

func (l Logger) Init(c config.Config) {
	l.log_folder = c.Log_folder
	l.data = make(map[string][]map[string]string)
}

func (l *Logger) Log() bool {
	fmt.Println(l)
	return true
}

func (l *Logger) AddNewItem(path string, element map[string]string) {
	l.data[path] = append(l.data[path], element)
}

func (l *Logger) GetLogFolder() string {
	return l.log_folder
}