package main

//Whatpulse bevat de data die we via de API krijgen
type Whatpulse struct {
	Clicks   int `json:"clicks"`
	Download int `json:"download"`
	Keys     int `json:"keys"`
	Upload   int `json:"upload"`
	Uptime   int `json:"uptime"`
}
