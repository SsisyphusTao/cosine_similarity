package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// parse command-line arguments
	file1 := flag.String("file1", "", "Path to the first tensor file")
	file2 := flag.String("file2", "", "Path to the second tensor file")
	flag.Parse()

	// read and parse the tensor files
	tensor1 := readTensor(*file1)
	tensor2 := readTensor(*file2)

	// calculate and print the cosine similarity
	cosineSimilarity := calculateCosineSimilarity(tensor1, tensor2)
	fmt.Printf("Cosine Similarity: %f\n", cosineSimilarity)
}

func printSlice(slice []float64) {
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}

func readTensor(filePath string) []float64 {
	// open the file, read it line by line, split each line on some delimiter,
	// and convert each field to a float64
	// return the result as a slice of float64s

	// open the file
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read line
	fileScanner := bufio.NewScanner(file)
	tensor := []float64{}
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		// convert line to float64
		v, err := strconv.ParseFloat(fileScanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		tensor = append(tensor, v)
	}
	printSlice(tensor)

	return tensor
}

func calcuateDotProduct(tensor1, tensor2 []float64) float64 {
	// calculate the dot product of two tensors
	// this involves multiplying each element of one tensor by the corresponding element of the other tensor,
	// and then summing the products
	sum := 0.0
	for i, v := range tensor1 {
		sum += v * tensor2[i]
	}
	return sum
}

func calculateMagnitude(tensor []float64) float64 {
	// calculate the magnitude of a tensor
	// this involves calculating the square root of the sum of the squares of each element
	sum := 0.0
	for _, v := range tensor {
		sum += math.Pow(v, 2)
	}
	return math.Sqrt(sum)
}

func calculateCosineSimilarity(tensor1, tensor2 []float64) float64 {
	// calculate the cosine similarity of the two tensors
	// this involves calculating the dot product of the tensors,
	// as well as the magnitude of each tensor
	printSlice(tensor1)
	printSlice(tensor2)
	dotProduct := calcuateDotProduct(tensor1, tensor2)
	magnitude1 := calculateMagnitude(tensor1)
	magnitude2 := calculateMagnitude(tensor2)
	fmt.Printf("dotProduct: %f\n", dotProduct)
	fmt.Printf("magnitude1: %f\n", magnitude1)
	fmt.Printf("magnitude2: %f\n", magnitude2)
	return dotProduct / (magnitude1 * magnitude2)
}
