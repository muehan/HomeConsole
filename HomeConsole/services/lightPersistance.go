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

	arrayOfLights := make([]models.Light, len(config.Lights)+1)
	maxindex := 1
	for i, c := range config.Lights {
		if maxindex < c.ID {
			maxindex = c.ID
		}
		arrayOfLights[i] = c
	}
	maxindex++

	newLight := models.Light{}
	newLight.ID = maxindex
	newLight.Name = post.Name
	newLight.URL = post.URL
	arrayOfLights[len(config.Lights)] = newLight

	config.Lights = arrayOfLights

	b, err := xml.Marshal(config)

	if err != nil {
		fmt.Println("error Marshal xml to byte array")
		fmt.Println(err)
	}

	ioutil.WriteFile("homeConfig.xml", b, 0777)
}
