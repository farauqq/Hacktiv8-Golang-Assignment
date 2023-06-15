package service

import (
	"assignment3/entity"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

func AutoReload() {
	for {
		water := RandomNumberWater()
		wind := RandomNumberWind()

		num := entity.StatusWaterWind{}
		num.Status.Water = water
		num.Status.Wind = wind

		jsonData, err := json.Marshal(num)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = ioutil.WriteFile("data.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(time.Second * 5)
	}
}

func ReloadWeb(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var status entity.StatusWaterWind

	err = json.Unmarshal(jsonData, &status)
	if err != nil {
		log.Fatal(err.Error())
	}

	water := status.Status.Water
	wind := status.Status.Wind

	var (
		statusWater string
		statusWind  string
	)

	statusWater = KondisiWater(water)
	statusWind = KondisiWind(wind)

	var waterAman = false
	var waterSiaga = false
	var waterBahaya = false
	if statusWater == "Aman" {
		waterAman = true
	} else if statusWater == "Siaga" {
		waterSiaga = true
	} else if statusWater == "Bahaya" {
		waterBahaya = true
	}

	var windAman = false
	var windSiaga = false
	var windBahaya = false
	if statusWind == "Aman" {
		windAman = true
	} else if statusWind == "Siaga" {
		windSiaga = true
	} else if statusWind == "Bahaya" {
		windBahaya = true
	}

	data := map[string]interface{}{
		"statusWater": statusWater,
		"statusWind":  statusWind,
		"water":       water,
		"wind":        wind,
		"waterAman":   waterAman,
		"waterSiaga":  waterSiaga,
		"waterBahaya": waterBahaya,
		"windAman":    windAman,
		"windSiaga":   windSiaga,
		"windBahaya":  windBahaya,
	}

	template, err := template.ParseFiles("./template/template.html")
	if err != nil {
		log.Fatal(err.Error())
	}

	template.Execute(w, data)

}
