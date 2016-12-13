package services

import "HomeConsole/HomeConsole/models"

/// GetLights returns all Lights as List of Light Models from lights.config
func GetLights() []*models.Light {
	light := new(models.Light)
	light.Name = "Büro"
	light.URL = "http://192.168.1.1/api/1"

	light2 := new(models.Light)
	light2.Name = "Wohnzimmer"
	light2.URL = "http://192.168.1.1/api/2"

	lights := []*models.Light{}
	lights = append(lights, light)
	lights = append(lights, light2)

	return lights
}
