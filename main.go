package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type point2D struct {
	x int
	y int
}

var (
	r         = regexp.MustCompile(`\((\d*),(\d*)\)`)
	waitGroup = sync.WaitGroup{}
)

const numberOfThreads = 8

func findArea(inputChannel chan string) {
	for pointsString := range inputChannel {
		var points []point2D

		for _, p := range r.FindAllStringSubmatch(pointsString, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])

			points = append(points, point2D{x, y})
		}

		area := 0.0

		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}
		fmt.Println(math.Abs(area) / 2.0)
	}
	waitGroup.Done()
}

func main() {

	dat, _ := ioutil.ReadFile("polygons.txt")
	text := string(dat)

	//making our buffered channel with size of 1000
	inputChannel := make(chan string, 1000)

	//declaring our threads
	for i := 0; i < numberOfThreads; i++ {
		go findArea(inputChannel)
	}

	waitGroup.Add(numberOfThreads)

	for _, line := range strings.Split(text, "\n") {
		inputChannel <- line
	}

	close(inputChannel)
	waitGroup.Wait()
}
