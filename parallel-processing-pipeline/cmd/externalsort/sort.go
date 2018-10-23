package main

import (
	"bufio"
	"fmt"
	"os"
	"parallel-processing-pipeline/pipline"
)

func main() {
	p := createPipeline("small.in", 400, 4)
	writeToFile(p, "small.out")
	printFile("small.out")
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipline.ReaderSource(file, -1)
	for v := range p {
		fmt.Println(v)
	}
}
func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipline.WriterSink(writer, p)
}
func createPipeline(fileName string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount

	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipline.InMemorySort(source))

	}
	return pipline.MergeN(sortResults...)
}
