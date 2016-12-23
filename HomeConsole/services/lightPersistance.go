package services

import (
	"HomeConsole/HomeConsole/helper"
	"HomeConsole/HomeConsole/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/// GetLights returns all Lights as List of Light Models from lights.config
func GetLights() []models.Light {
	config := loadConfiguration()
	return config.Lights
}

func AddLight(post models.Post) {

	// parse xml to config model
	config := loadConfiguration()

	// find next higest ID and put all array item into the slice
	// slice is needed to append the new light at the end
	sliceOfLights := make([]models.Light, len(config.Lights)+1)
	maxindex := 1
	for i, c := range config.Lights {
		if maxindex < c.ID {
			maxindex = c.ID
		}
		sliceOfLights[i] = c
	}
	maxindex++

	// map post model into new light and set to last item of the slice
	newLight := models.Light{}
	newLight.ID = maxindex
	newLight.Name = post.Name
	newLight.URL = post.URL
	sliceOfLights[len(config.Lights)] = newLight

	saveConfiguration(sliceOfLights)
}

func ChangeLight(postmodel models.Post) {

	lights := GetLights()

	searchID := helper.StringToInt(postmodel.ID)

	for i, light := range lights {

		if light.ID == searchID {
			lights[i].Name = postmodel.Name
			lights[i].URL = postmodel.URL
		}
	}

	saveConfiguration(lights)
}

func GetLight(id int) (light models.Light) {

	lights := GetLights()

	for _, l := range lights {
		if l.ID == id {
			light = l
		}
	}

	return light
}

func loadConfiguration() (config models.Config) {
	// parse xml to config model
	xmlFile, err := os.Open("homeConfig.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer xmlFile.Close()

	bytes, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		fmt.Println("Error Unmarshal file input")
	}

	xml.Unmarshal(bytes, &config)

	return config
}

func saveConfiguration(lights []models.Light) {
	var config models.Config

	config.Lights = lights

	b, err := xml.Marshal(config)

	if err != nil {
		fmt.Println("error Marshal xml to byte array")
		fmt.Println(err)
	}

	ioutil.WriteFile("homeConfig.xml", b, 0777)
}
