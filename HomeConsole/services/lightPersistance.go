package services

import (
	"HomeConsole/HomeConsole/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/// GetLights returns all Lights as List of Light Models from lights.config
func GetLights() []models.Light {

	xmlFile, err := os.Open("homeConfig.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer xmlFile.Close()

	b, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		fmt.Println("Error Unmarshal file input")
	}

	var c models.Config
	xml.Unmarshal(b, &c)

	return c.Lights
}

func AddLight(post models.Post) {

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

	var config models.Config
	xml.Unmarshal(bytes, &config)

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

	config.Lights = sliceOfLights

	b, err := xml.Marshal(config)

	if err != nil {
		fmt.Println("error Marshal xml to byte array")
		fmt.Println(err)
	}

	ioutil.WriteFile("homeConfig.xml", b, 0777)
}
