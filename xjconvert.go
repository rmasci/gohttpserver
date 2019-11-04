package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func xjconverter(w http.ResponseWriter, r *http.Request) {
	var convert, format, data string
	fmt.Fprintln(w, "")
	if r.Method == "POST" {

		r.ParseForm()
		for key, val := range r.Form {
		switch key {
			case "data":
				data = strings.Join(val, "")
				//fmt.Fprintf(w,"Data: %v\n", data)
			case "convert":
				convert = strings.Join(val, "")
				//fmt.Fprintf(w, "Convert: %v\n",convert)
			case "format":
				format = strings.Join(val, "")
				//fmt.Fprintf(w,"Format: %v\n",format)
			}
		}
		
		if convert == "" || format == "" || data == "" {
			fmt.Fprintf(w, "Convert: %v\nFormat: %v\nData: %v\n", convert, format, data)
			return
		}
		fmt.Fprintf(w, "%v", xjconverterPost(convert, format, data))
		
	} else {
		fmt.Fprintf(w, "Hello There!")
	}
}

func xjconverterPost(convert, format, data string) string {
	var out bytes.Buffer
	var cmdFunction string
	var cmdPretty string
	switch convert {
	case "1":
		cmdFunction = "x2j"
	case "2":
		cmdFunction = "j2x"
	case "3":
		cmdFunction = "format"
	}

	switch format {
	case "1":
		if cmdFunction != "format" {
			cmdPretty = "-p"
		}
	case "2":
		cmdPretty = "-u"
	}

	cmd := exec.Command("/usr/local/bin/xjconvert", cmdFunction, cmdPretty)
	cmd.Stdin = bytes.NewBuffer([]byte(data))
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}
	return string(out.Bytes())
}
