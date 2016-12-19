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

func SetLights(lights *[]models.Light) {

	fmt.Println("Save new lights")
	fmt.Println(lights)

	b, err := xml.Marshal(lights)

	if err != nil {
		fmt.Println("error dunring xml parsing")
	}

	ioutil.WriteFile("homeConfig.xml", b, 0777)
}
