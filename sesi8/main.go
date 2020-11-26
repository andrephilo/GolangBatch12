package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

var path = "data.json"

type RAW struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

func write() {
	for {
		var file, err = os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			return
		}
		defer file.Close()

		dataRAW := RAW{}
		dataRAW.Status.Water = rand.Intn(20)
		dataRAW.Status.Wind = rand.Intn(20)
		jsonData, _ := json.Marshal(dataRAW)
		err = ioutil.WriteFile("data.json", jsonData, 0644)
		time.Sleep(time.Second * 15)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Page index called")

	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	var raw RAW

	json.Unmarshal(data, &raw)

	var StatusWater = raw.Status.Water
	var StatusWind = raw.Status.Wind

	type Status struct {
		StatusWater string
		StatusWind  string
	}

	const tpl = `
		<html>
			<meta http-equiv="refresh" content="15" />
			<body>
				Status Wind : {{.StatusWater}}
				Status Water : {{.StatusWind}}
			</body>
		</html>
	`
	status := Status{}
	switch {
	case (StatusWater < 5):
		status.StatusWater = "Aman"
	case (StatusWater >= 6) && (StatusWater <= 8):
		status.StatusWater = "Siaga"
	case (StatusWater > 8):
		status.StatusWater = "Bahaya"
	}

	switch {
	case (StatusWind < 6):
		status.StatusWind = "Aman"
	case (StatusWind >= 7) && (StatusWind <= 15):
		status.StatusWind = "Siaga"
	case (StatusWind > 15):
		status.StatusWind = "Bahaya"
	}

	tmplt, err := template.New("index").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}
	tmplt.Execute(w, status)
}

func init() {
	go write()
}

func main() {
	fmt.Println("Listen to address :8080")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
