package main

import (
	"github.com/L-F-Z/cee/graph"
	"github.com/L-F-Z/cee/identify"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	var g *graph.Graph = graph.NewGraph()
	if len(os.Args) < 2 {
		panic("Need to provide a json file for constructing causal graph")
	}
	str, _ := ioutil.ReadFile(os.Args[1])
	var dat map[string]interface{}
    if err := json.Unmarshal(str, &dat); err != nil {
        panic(err)
    }
	fmt.Println("NODES")
	nodes := dat["nodes"].([]interface{})
	for _, e := range nodes {
		node := e.(string)
		g.AddNode(node, true)
		fmt.Println(node)
	}
	fmt.Println("EDGES")
	edges := dat["edges"].([]interface{})
	for _, e := range edges {
		edge := e.([]interface{})
		g.AddEdge(edge[0].(string), edge[1].(string))
		fmt.Println(edge[0].(string), edge[1].(string))
	}
	fmt.Println("BIDIRECT")
	bidirect := dat["bidirect"].([]interface{})
	for _, e := range bidirect {
		edge := e.([]interface{})
		g.AddBidirect(edge[0].(string), edge[1].(string))
		fmt.Println(edge[0].(string), edge[1].(string))
	}
	x := make([]int64, 0)
	nodes = dat["x"].([]interface{})
	for _, e := range nodes {
		x = append(x, g.NodeID(e.(string)))
	}
	fmt.Println("X = ", x)
	y := make([]int64, 0)
	nodes = dat["y"].([]interface{})
	for _, e := range nodes {
		y = append(y, g.NodeID(e.(string)))
	}
	fmt.Println("Y = ", y)
	z := make([]int64, 0)
	nodes = dat["z"].([]interface{})
	for _, e := range nodes {
		z = append(z, g.NodeID(e.(string)))
	}
	fmt.Println("Z = ", z)
	p := identify.Identify(y, x, z, g)
	fmt.Println("\n------Result------")
	fmt.Println(p)
}

/*
func main() {
	var g *graph.Graph = graph.NewGraph()
	g.AddNode("X", true)
	g.AddNode("Y", true)
	g.AddNode("Z1", true)
	g.AddNode("Z2", true)
	g.AddEdge("Z1", "X")
	g.AddEdge("X", "Z2")
	g.AddEdge("Z2", "Y")
	g.AddBidirect("X", "Y")
	g.AddBidirect("X", "Z1")
	g.AddBidirect("Y", "Z1")
	g.AddBidirect("Z2", "Z1")
	fmt.Println(g)
	p := identify.Identify(y, x, g)
	fmt.Println(p)
}
*/