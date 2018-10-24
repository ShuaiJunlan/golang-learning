package main

import (
	"bufio"
	"fmt"
	"os"
	"parallel-processing-pipeline/pipline"
)

func main() {
	const filename = "large.in"
	const n = 300000000

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	p := pipline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipline.WriterSink(
		writer, p)
	writer.Flush() //if doesn't use this statement, may result in the size of 'large.in' less than 800000000 byte

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipline.ReaderSource(
		bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
	//file.Close()
}

func mergetDemo() {
	p := pipline.Merge(
		pipline.InMemorySort(
			pipline.ArraySource(3, 2, 6, 7, 4)),
		pipline.InMemorySort(
			pipline.ArraySource(6, 7, 812, 5, 56, 786)))
	//for true {
	//	if num, ok := <- p; ok{
	//		fmt.Println(num)
	//	}else {
	//		break
	//	}
	//
	//}
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}
