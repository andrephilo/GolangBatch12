package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
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

		// _, err = file.Write(jsonData)

		// if err != nil {
		// 	return
		// }

		// err = file.Sync()
		// if err != nil {
		// 	return
		// }

		time.Sleep(time.Second * 15)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	for {
		fmt.Println("Page index called")

		var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
		if err != nil {
			return
		}
		defer file.Close()

		data, _ := ioutil.ReadAll(file)

		var raw RAW

		json.Unmarshal(data, &raw)

		fmt.Fprintln(w, "Status")

		var statusWater = raw.Status.Water
		var statusWind = raw.Status.Wind

		switch {
		case (statusWater < 5):
			fmt.Fprintln(w, "Water: Aman")
		case (statusWater >= 6) && (statusWater <= 8):
			fmt.Fprintln(w, "Water: Siaga")
		case (statusWater > 8):
			fmt.Fprintln(w, "Water: Bahaya")
		}

		switch {
		case (statusWind < 6):
			fmt.Fprintln(w, "Wind: Aman")
		case (statusWind >= 7) && (statusWind <= 15):
			fmt.Fprintln(w, "Wind: Siaga")
		case (statusWind > 15):
			fmt.Fprintln(w, "Wind: Bahaya")
		}

		// var t, _ = template.ParseFiles("template.html")
		// if err != nil {
		// 	return
		// }
		// var result = map[string]interface{}{
		// 	"Water": statusWater,
		// 	"Wind":  statusWind,
		// }
		// t.Execute(w, result)

		time.Sleep(time.Second * 15)
	}
}

func init() {
	go write()
}

func main() {
	fmt.Println("Listen to address :8080")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
