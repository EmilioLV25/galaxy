package main

import (
	"fmt"
	"math"
)

// Respuesta es el formato json
type Respuesta struct {
	Dia   int    `json:"dia"`
	Clima string `json:"clima"`
}

// Planet es la estructura de los planetas del sistema
type Planet struct {
	Distance float64 //500 1000 2000 km
	Degree   float64
	Sense    float64 // -1; -3; 5
}

/*
	newDay mueve los grados en sentido horario o antihorario a los planetas del sistema
*/
func newDay(planets []Planet) {
	for i := 0; i < len(planets); i++ {
		if planets[i].Sense < 0 && planets[i].Degree <= 0 {
			planets[i].Degree = 360
		} else {
			if planets[i].Sense > 0 && planets[i].Degree >= 360 {
				planets[i].Degree = 0
			}
		}
		planets[i].Degree += planets[i].Sense
	}
}

func setTernas(planets []Planet, ternas *[][]float64) {
	myTernas := *ternas
	for k, v := range planets {
		myTernas[k][0] = math.Trunc(v.Distance * math.Cos((v.Degree*math.Pi)/180)) // X
		myTernas[k][1] = math.Trunc(v.Distance * math.Sin((v.Degree*math.Pi)/180)) // Y
	}
}

func whatHappen(ternas [][]float64) string {
	// ALINEADOS
	if getArea(ternas[0], ternas[1], ternas[2]) <= 66200 { // 3 PLANETAS ALINEADOS...
		if getArea(ternas[0], ternas[1], []float64{0, 0}) <= 66200 && getArea(ternas[0], ternas[2], []float64{0, 0}) <= 66200 { // CON EL SOL...
			return "sequia"
		}
		// SIN EL SOL
		return "condiciones optimas"
	}

	// TRIANGULO
	sumAreas := getArea([]float64{0, 0}, ternas[1], ternas[2]) + getArea(ternas[0], []float64{0, 0}, ternas[2]) + getArea(ternas[0], ternas[1], []float64{0, 0})
	areaTriangulo := getArea(ternas[0], ternas[1], ternas[2])
	if sumAreas <= areaTriangulo {
		if areaTriangulo == 1435791 {
			return "lluvias intensas"
		}

		return "lluvias"
	}

	return "normal"
}

// Devuelve el area de un triangulo utilizando el determinante.
func getArea(a []float64, b []float64, c []float64) float64 {
	return math.Trunc(math.Abs(a[0]*b[1]+a[1]*c[0]+b[0]*c[1]-b[1]*c[0]-c[1]*a[0]-a[1]*b[0]) / 2)
}

func generarDatos(planets []Planet) []Respuesta {
	var sequia, lluvias, cOptimas int = 0, 0, 0
	var maxLluvias []int
	var lastClima string

	var arrayDays []Respuesta

	ternas := make([][]float64, len(planets))
	for i := 0; i < len(ternas); i++ {
		ternas[i] = make([]float64, 2)
	}

	for day := 0; day < 3650; day++ {
		// GENERAR TERNAS
		setTernas(planets, &ternas)

		// WHAT HAPPEN ON THIS DAY
		info := whatHappen(ternas)

		// PERIODOS
		if lastClima != info {
			if lastClima == "sequia" {
				sequia++
			} else {
				if lastClima == "lluvias" {
					lluvias++
				} else {
					if lastClima == "condiciones optimas" {
						cOptimas++
					} else {
						if lastClima == "lluvias intensas" {
							maxLluvias = append(maxLluvias, day)
						}
					}
				}
			}
			lastClima = info
		}

		// SAVE INFO
		arrayDays = append(arrayDays, Respuesta{Dia: day + 1, Clima: info})

		// NEW DAY
		newDay(planets)
	}
	fmt.Println("sequia: ", sequia)
	fmt.Println("lluvia: ", lluvias)
	fmt.Println("cond. optimas: ", cOptimas)
	fmt.Println("lluvias intensas: ", maxLluvias)

	return arrayDays
}

func setPlanet(Distance float64, Sense float64) Planet {
	planet := Planet{
		Distance: Distance,
		Degree:   0,
		Sense:    Sense,
	}

	return planet
}

func run() []Respuesta {
	planets := []Planet{
		setPlanet(500, -1), setPlanet(2000, -3), setPlanet(1000, 5),
	}

	return generarDatos(planets)
}
