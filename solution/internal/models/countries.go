package models

type Countries struct {
	Name   string `json:"name"`
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
	Region string `json:"region"`
}
