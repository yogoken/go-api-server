package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 1. JSON define
type Input struct {
	In string
}

type Output struct {
	Out string
}

// 2. controller
func jsonHandleFunc(rw http.ResponseWriter, req *http.Request) {
	output := Output{"It's return"}
	// 2.1 JSON return
	defer func() {
		outjson, e := json.Marshal(output)
		if e != nil {
			fmt.Println(e)
		}
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(outjson))
	}()
	// 2.2 type check
	if req.Method != "POST" {
		output.Out = "Not post..."
		return
	}
	// 2.3 request body
	body, e := ioutil.ReadAll(req.Body)
	if e != nil {
		output.Out = e.Error()
		fmt.Println(e.Error())
		return
	}
	// 2.4 JSON parse
	input := Input{}
	e = json.Unmarshal(body, &input)
	if e != nil {
		output.Out = e.Error()
		fmt.Println(e.Error())
		return
	}
	fmt.Printf("%#v\n", input)
}

// 3. main
func main() {
	fs := http.FileServer(http.Dir("static"))
	// 3.1 routes handler
	http.Handle("/", fs)
	http.HandleFunc("/json", jsonHandleFunc)
	err := http.ListenAndServe(":8080", nil)

	// 3.2 error abort
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
