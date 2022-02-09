package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"os"
	"strconv"
)



func main() {
	if len(os.Args) < 2 {
		panic("Need to provide graph size")
	}
	arg1, _ := strconv.Atoi(os.Args[1])
	var maxNode int64 = int64(arg1)
	var outDegree int64 = 2
	var i, from, to int64
	rand.Seed(time.Now().UnixNano())
	file, err := os.Create("gen.json")
    if err != nil{
        fmt.Println(err)
    }
	defer file.Close()

	// nodes
	file.WriteString("{\n    \"nodes\": [\n")
	for i = 1; i < maxNode; i++ {
		file.WriteString("        \"" + strconv.FormatInt(i, 10) + "\",\n")
	}
	file.WriteString("        \"" + strconv.FormatInt(maxNode, 10) + "\"\n    ],\n")

	// edges
	file.WriteString("    \"edges\": [\n")
	edges := make([][]int64, 0)
	for from = 1; from <= maxNode; from++ {
		probability := float64(outDegree) / float64(maxNode - from + 1)
		probability = math.Min(probability, 0.75)
		connected := false
		for to = from+1; to <= maxNode; to++ {
			if rand.Float64() < probability {
				edges = append(edges, []int64{from, to})
				connected = true
			}
		}
		if !connected {
			if from != maxNode {
				edges = append(edges, []int64{from, from+1})
			}
		}
	}
	for i = 0; i < int64(len(edges)-1); i++ {
		file.WriteString("        [\"" + strconv.FormatInt(int64(edges[i][0]), 10) + 
		                 "\",\"" + strconv.FormatInt(int64(edges[i][1]), 10) + "\"],\n")
	}
	file.WriteString("        [\"" + strconv.FormatInt(int64(edges[i][0]), 10) + 
					 "\",\"" + strconv.FormatInt(int64(edges[i][1]), 10) + "\"]\n    ],\n")

	// bidirect
	file.WriteString("    \"bidirect\": [\n")
	bidirect := make([][]int64, 0)
	probability := float64(outDegree) / float64(maxNode)
	probability = math.Min(probability, 0.4)
	for from = 1; from <= maxNode; from++ {
		for to = from+1; to <= maxNode; to++ {
			if rand.Float64() < probability {
				bidirect = append(bidirect, []int64{from, to})
			}
		}
	}
	for i = 0; i < int64(len(bidirect)-1); i++ {
		file.WriteString("        [\"" + strconv.FormatInt(int64(bidirect[i][0]), 10) + 
		                 "\",\"" + strconv.FormatInt(int64(bidirect[i][1]), 10) + "\"],\n")
	}
	file.WriteString("        [\"" + strconv.FormatInt(int64(bidirect[i][0]), 10) + 
					 "\",\"" + strconv.FormatInt(int64(bidirect[i][1]), 10) + "\"]\n    ],\n")

	// settings
	file.WriteString("    \"y\": [\"")
	y := rand.Intn(99) + 1
	file.WriteString(strconv.FormatInt(int64(y), 10) + "\"],\n")
	file.WriteString("    \"x\": [\"")
	x := rand.Intn(99) + 1
	for x == y {
		x = rand.Intn(99) + 1
	}
	file.WriteString(strconv.FormatInt(int64(x), 10) + "\"],\n")
	file.WriteString("    \"z\": []\n}")
	fmt.Println("generated a graph with", maxNode, "nodes!")
}