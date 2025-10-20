package logger

import (
	"log"
	"os"
)

var std = log.New(os.Stdout, "[cyber-av] ", log.LstdFlags|log.Lmsgprefix)

func Infof(format string, v ...any)  { std.Printf("INFO: "+format, v...) }
func Errorf(format string, v ...any) { std.Printf("ERROR: "+format, v...) }
