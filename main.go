package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		wd := os.Args[1]
		chopdir(wd)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		logf(err)
	}
	files, err := ioutil.ReadDir(wd)
	if err != nil {
		logf(err)
	}
	for _, file := range files {
		if !existsDir(filepath.Join(wd, file.Name())) {
			continue
		}
		chopdir(filepath.Join(wd, file.Name()))
	}
}

func chopdir(wd string) {
	if !existsDir(wd) {
		logf("ディレクトリは存在しません。")
	}

	if !isAloneDir(wd) {
		return
	}

	pwd, err := filepath.Abs(filepath.Dir(wd))
	if err != nil {
		logf(err)
	}
	wd, err = filepath.Abs(wd)
	if err != nil {
		logf(err)
	}

	newdir := chopMove(pwd, wd)
	os.Remove(wd)
	os.Rename(newdir, wd)
}

func existsDir(d string) bool {
	f, err := os.Stat(d)
	if os.IsNotExist(err) {
		return false
	}
	if !f.IsDir() {
		return false
	}
	return true
}

func chopMove(pwd string, wd string) string {
	if isAloneDir(wd) {
		return chopMove(pwd, child(wd))
	}
	todir := getMovetoDir(pwd, wd)
	err := os.Rename(wd, todir)
	if err != nil {
		logf(err)
	}
	return todir
}

func getMovetoDir(pwd string, wd string) string {
	todir := filepath.Join(pwd, filepath.Base(wd))
	for i := 0; existsDir(todir); i++ {
		todir = todir + strconv.Itoa(i)
	}
	return todir

}

func isAloneDir(d string) bool {
	if !existsDir(d) {
		return false
	}

	files, err := ioutil.ReadDir(d)
	if err != nil {
		fmt.Println(runtime.Caller(1))
		logf(err)
	}
	if len(files) != 1 {
		return false
	}
	if existsDir(filepath.Join(d, files[0].Name())) {
		return true
	}
	return false
}

func child(s string) string {
	files, err := ioutil.ReadDir(s)
	if err != nil {
		logf(err)
	}
	return filepath.Join(s, files[0].Name())
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
func getFilename(ext string) string {
	exec, _ := os.Executable()
	return filepath.Join(filepath.Dir(exec), getFileNameWithoutExt(exec)+ext)
}

func logg(m interface{}) {
	f, err := os.OpenFile(getFilename(".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Cannot open log file: " + err.Error())
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stderr))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(m)
}
func logf(m interface{}) {
	logg(m)
	os.Exit(1)
}
