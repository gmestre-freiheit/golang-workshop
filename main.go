package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// I just convert json to golang on web
// https://transform.tools/json-to-go
type PlanetsResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Planet    `json:"results"`
}

type Planet struct {
	Name           string    `json:"name"`
	RotationPeriod string    `json:"rotation_period"`
	OrbitalPeriod  string    `json:"orbital_period"`
	Diameter       string    `json:"diameter"`
	Climate        string    `json:"climate"`
	Gravity        string    `json:"gravity"`
	Terrain        string    `json:"terrain"`
	SurfaceWater   string    `json:"surface_water"`
	Population     string    `json:"population"`
	Residents      []string  `json:"residents"`
	Films          []string  `json:"films"`
	Starships      []string  `json:"starships"`
	Created        time.Time `json:"created"`
	Edited         time.Time `json:"edited"`
	URL            string    `json:"url"`
}

type Resident struct {
	Name      string        `json:"name"`
	Height    string        `json:"height"`
	Mass      string        `json:"mass"`
	HairColor string        `json:"hair_color"`
	SkinColor string        `json:"skin_color"`
	EyeColor  string        `json:"eye_color"`
	BirthYear string        `json:"birth_year"`
	Gender    string        `json:"gender"`
	Homeworld string        `json:"homeworld"`
	Films     []string      `json:"films"`
	Species   []interface{} `json:"species"`
	Vehicles  []interface{} `json:"vehicles"`
	Starships []string      `json:"starships"`
	Created   time.Time     `json:"created"`
	Edited    time.Time     `json:"edited"`
	URL       string        `json:"url"`
}

type PlanetwithResidents struct {
	PlanetName     string   `json:"name"`
	ResidentsNames []string `json:"residents"`
}

func getPlanetsList() ([]Planet, error) {
	var planets []Planet
	planetsRequestURI := "https://swapi.dev/api/planets/"
	client := &http.Client{}
	for planetsRequestURI != "" {
		req, err := http.NewRequest("GET", planetsRequestURI, nil)
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var planet *PlanetsResponse
		err = json.Unmarshal(body, &planet)
		if err != nil {
			return nil, err
		}
		planetsRequestURI = planet.Next
		planets = append(planets, planet.Results...)
	}
	return planets, nil
}

func main() {
	planets, err := getPlanetsList()
	if err != nil {
		panic(fmt.Errorf("unable to get the planets list %w", err))
	}

	for i, planet := range planets {
		fmt.Printf("Planet %v is named %v\n", i+1, planet.Name)
	}
}
