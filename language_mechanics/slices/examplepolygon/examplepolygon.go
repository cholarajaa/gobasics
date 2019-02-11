package main

import (
	"fmt"
)

// Polygon defines a set of points that complete a ring around a geographic area
type Polygon [][2]float64

// PolygonRings defines a MongoDB Structure for storing multiple polygon rings
type PolygonRings struct {
	Type        string    `bson:"type"`
	Coordinates []Polygon `bson:"coordinates"`
}

// Represents a marine station and its polygons
type MarineStation struct {
	StationId string       `bson:"station_id"`
	Polygons  PolygonRings `bson:"polygons"`
}

func main() {

	// Create an empty slice to store the polygon rings
	// for the different marine stations
	marineStations := []MarineStation{}

	// Create a marine station for AMZ123
	marineStation := MarineStation{
		StationId: "AMZ123",
		Polygons: PolygonRings{
			Type:        "Polygon",
			Coordinates: []Polygon{},
		},
	}

	// Create the points for the first polygon ring
	point1 := [2]float64{-79.7291190729999, 26.9729398600001}
	point2 := [2]float64{-80.0799532019999, 26.9692689500001}
	point3 := [2]float64{-80.0803627959999, 26.970533371}
	point4 := [2]float64{-80.0810508729999, 26.975004196}
	point5 := [2]float64{-79.7291190729999, 26.9729398600001}

	// Create a polygon for this ring
	polygon := Polygon{point1, point2, point3, point4, point5}

	// Add the polygon to the slice of polygon coordinates
	marineStation.Polygons.Coordinates = append(marineStation.Polygons.Coordinates, polygon)

	// Create the points for the second polygon ring
	point1 = [2]float64{-80.4370117189999, 27.7877197270001}
	point2 = [2]float64{-80.4376220699999, 27.7885131840001}
	point3 = [2]float64{-80.4384155269999, 27.7885131840001}
	point4 = [2]float64{-80.4370117189999, 27.7877197270001}

	// Create a polygon for this ring
	polygon = Polygon{point1, point2, point3, point4}

	// Add the polygon to the slice of polygon coordinates
	marineStation.Polygons.Coordinates = append(marineStation.Polygons.Coordinates, polygon)

	// Add the marine station
	marineStations = append(marineStations, marineStation)

	Display(marineStations)
}

func Display(marineStations []MarineStation) {

	for _, marineStation := range marineStations {

		fmt.Printf("\nStation: %s\n", marineStation.StationId)

		for index, rings := range marineStation.Polygons.Coordinates {

			fmt.Printf("Ring: %d\n", index)

			for _, coordinate := range rings {

				fmt.Printf("Point: %f,%f\n", coordinate[0], coordinate[1])
			}
		}
	}
}
