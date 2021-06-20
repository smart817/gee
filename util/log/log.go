package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

var (
	//Trace   *log.Logger //几乎所有信息
	infolog *log.Logger //重要
	//errorlog *log.Logger //
	herr *log.Logger
	//Warning *log.Logger //
)

func init() {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("无法打开错误文件", err)
	}
	//Trace = log.New(ioutil.Discard, "TRACE:", log.Ldate|log.Ltime|log.Lshortfile)
	infolog = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	//errorlog = log.New(io.MultiWriter(file, os.Stderr), "ERROR:", log.Ldate|log.Ltime|log.Llongfile)
	herr = log.New(io.MultiWriter(file, os.Stderr), "ERROR:", log.Ldate|log.Ltime)
	//Err = errorlog.Println
	Info = infolog.Println

}

//var Err = errorlog.Println
var Info = infolog.Println

func Err(err ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	fmt.Println(ok)
	if err != nil && ok {
		var slice []interface{}
		slice = append(slice, file, line)
		slice = append(slice, err...)
		herr.Println(slice...)
		//Err(err)
	}
}
