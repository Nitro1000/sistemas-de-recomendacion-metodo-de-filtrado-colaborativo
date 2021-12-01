package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func remove(s []float64, index int) []float64 {
	return append(s[:index], s[index+1:]...)
}
func Pearson(person1, person2 []float64, emptyElement Pair) (result float64) {
	// person1Aux, person2Aux := person1, person2
	person1 = remove(person1, emptyElement.posJ)
	person2 = remove(person2, emptyElement.posJ)

	sum := 0.0
	for _, v := range person1 {
		sum += v
	}
	averageP1 := sum / float64(len(person1))
	sum = 0
	for _, v := range person2 {
		sum += v
	}
	averageP2 := sum / float64(len(person2))

	num := 0.0
	for i := 0; i < len(person1); i++ {
		num += ((person1[i] - averageP1) * (person2[i] - averageP2))
	}
	aux1 := 0.0
	aux2 := 0.0
	for i := 0; i < len(person1); i++ {
		aux1 += math.Pow(float64(person1[i]-averageP1), 2)
		aux2 += math.Pow(float64(person2[i]-averageP2), 2)
	}
	den := math.Sqrt(aux1) * math.Sqrt(aux2)
	result = float64(num) / den
	fmt.Println(person1, person2)
	return
}

func Coseno(person1, person2 []float64, emptyElement Pair) (result float64) {
	person1 = remove(person1, emptyElement.posJ)
	person2 = remove(person2, emptyElement.posJ)

	num := 0.0
	for i := 0; i < len(person1); i++ {
		num += (person1[i] * person2[i])
	}
	aux1 := 0.0
	aux2 := 0.0
	for i := 0; i < len(person1); i++ {
		aux1 += math.Pow(float64(person1[i]), 2)
		aux2 += math.Pow(float64(person2[i]), 2)
	}
	den := math.Sqrt(aux1) * math.Sqrt(aux2)
	result = float64(num) / den
	fmt.Println(person1, person2)
	return
}

func Euclide(person1, person2 []float64, emptyElement Pair) (result float64) {
	person1 = remove(person1, emptyElement.posJ)
	person2 = remove(person2, emptyElement.posJ)

	result = 0.0
	for i := 0; i < len(person1); i++ {
		result += math.Pow(float64(person1[i])-float64(person2[i]), 2)
	}
	result = math.Sqrt(result)
	fmt.Println(person1, person2)
	return
}

func SimplePrediction(matrix [][]float64, metricValues []float64, neighborSelect []int, emptyElement Pair) (result float64) {
	num, den := 0.0, 0.0

	for i := 0; i < len(neighborSelect); i++ {
		num += metricValues[neighborSelect[i]] * matrix[neighborSelect[i]][emptyElement.posJ]
		den += math.Abs(metricValues[neighborSelect[i]])
	}
	result = (num / den)
	return
}

func MiddlePrediction(matrix [][]float64, metricValues []float64, neighborSelect []int, emptyElement Pair) (result float64) {
	num, den := 0.0, 0.0
	sum := 1.0
	for _, v := range matrix[emptyElement.posI] {
		sum += v
	}
	averageUser := sum / float64(len(matrix[emptyElement.posI])-1)
	averageUsers := 0.0
	for i := 0; i < len(neighborSelect); i++ {
		sum = 0.0
		for _, v := range matrix[neighborSelect[i]] {
			sum += v
		}
		averageUsers = sum / float64(len(matrix[neighborSelect[i]]))
		num += metricValues[neighborSelect[i]] * (matrix[neighborSelect[i]][emptyElement.posJ] - averageUsers)
		den += math.Abs(metricValues[neighborSelect[i]])
	}
	result = averageUser + (num / den)
	return
}

type Pair struct {
	posI int
	posJ int
}

func main() {
	nameFile := flag.String("name", "tabla.txt", "Nombre del fichero.")
	metric := flag.String("metric", "CP", "Métrica elegida. Los posibles valores son:\n1. CP (Correlación de Pearson).\n2. DC (Distancia coseno).\n3. DE (Distancia Euclídea).")
	neighbors := flag.Int("neighbors", 3, "Número de vecinos considerado.")
	prediction := flag.String("prediction", "DM", "Tipo de predicción:\n1. PS (Predicción simple).\n2. DM (Diferencia con la media).")
	flag.Parse()
	fmt.Println("Nombre del fichero: ", *nameFile)
	fmt.Println("Métrica elegida: ", *metric)
	fmt.Println("Número de vecinos considerado: ", *neighbors)
	fmt.Println("Tipo de predicción: ", *prediction)

	file, err := os.Open(*nameFile)

	if err != nil {
		fmt.Printf("Error abriendo archivo %s: %v", *nameFile, err)
		log.Fatal(err)
	}
	defer file.Close()

	bufer := make([]byte, 1)
	var matrixAux [][]string
	var rowAux []string
	for {
		byte, err := file.Read(bufer)
		char := string(bufer[:byte])
		if char == "-" {
			char = "-1"
		}
		if char != " " && char != "\n" {
			rowAux = append(rowAux, char)
		}
		if char == "\n" {
			matrixAux = append(matrixAux, rowAux)
			rowAux = nil
			// fmt.Println("Leido un endline")
		}
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error leyendo contenido: %v", err)
				log.Fatal(err)
			}
			matrixAux = append(matrixAux, rowAux)
			break
		}
		// fmt.Println("Leído este fragmento: ", char)
	}

	// fmt.Println(matrixAux)

	var matrix [][]float64
	var row []float64
	var emptyElement Pair

	for i := 0; i < len(matrixAux); i++ {
		for j := 0; j < len(matrixAux[i]); j++ {
			if matrixAux[i][j] == "-1" {
				emptyElement.posI = i
				emptyElement.posJ = j
			}
			if matrixAux[i][j] != "" {
				element, err := strconv.ParseFloat(matrixAux[i][j], 64)
				if err != nil {
					log.Fatal(err)
				}
				row = append(row, element)
			}

		}
		matrix = append(matrix, row)
		row = nil
	}

	var metricValues []float64

	for i, v := range matrix {
		if i != emptyElement.posI {
			switch *metric {
			case "CP":
				metricValues = append(metricValues, Pearson(matrix[emptyElement.posI], v, emptyElement))
			case "DC":
				metricValues = append(metricValues, Coseno(matrix[emptyElement.posI], v, emptyElement))
			case "DE":
				metricValues = append(metricValues, Euclide(matrix[emptyElement.posI], v, emptyElement))
			default:
				log.Fatal("ERROR: unknown metric")
			}
		} else {
			metricValues = append(metricValues, (-1 * math.MaxFloat64))
		}
	}
	fmt.Println(metricValues)

	var neighborSelect []int
	var tempMetricValues = make([]float64, len(metricValues))
	copy(tempMetricValues, metricValues)
	for i := 0; i < *neighbors; i++ {
		index := indexOfMaxNumber(tempMetricValues)
		neighborSelect = append(neighborSelect, index)
		tempMetricValues[index] = (-1 * math.MaxFloat64)
	}

	predict := 0.0

	switch *prediction {
	case "PS":
		predict = SimplePrediction(matrix, metricValues, neighborSelect, emptyElement)
	case "DM":
		predict = MiddlePrediction(matrix, metricValues, neighborSelect, emptyElement)
	default:
		log.Fatal("ERROR: unknown prediction")
	}

	fmt.Println(predict)

	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		fmt.Print(matrixAux[i][j])
	// 	}
	// 	fmt.Println()
	// }

}

func indexOfMaxNumber(array []float64) (result int) {
	aux := array[0]
	for i, v := range array {
		if v > aux {
			aux = v
			result = i
		}
	}
	return
}
