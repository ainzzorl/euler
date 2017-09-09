package main

// https://projecteuler.net/problem=107
// Minimal network
// Prim's algorithm.

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func p107(intputFilePath string) int {
	edges := [][]int{}
	file, _ := os.Open(intputFilePath)
	defer file.Close()

	totalEdges := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		parts := strings.Split(line, ",")
		for _, part := range parts {
			var edge int
			if part == "-" {
				edge = 0
			} else {
				edge, _ = strconv.Atoi(part)
			}
			totalEdges += edge
			row = append(row, edge)
		}
		edges = append(edges, row)
	}

	totalEdges /= 2

	return totalEdges - p107PrimWeight(edges)
}

func p107PrimWeight(edges [][]int) int {
	n := len(edges)
	dist := make([]int, n)
	visited := make([]bool, n)
	for i := range dist {
		dist[i] = math.MaxUint32
		visited[i] = false
	}
	dist[0] = 0
	result := 0
	for {
		min := -1
		minDist := math.MaxUint32
		for vertex := range dist {
			if !visited[vertex] && dist[vertex] < minDist {
				min = vertex
				minDist = dist[vertex]
			}
		}
		if min == -1 {
			break
		}
		result += dist[min]
		visited[min] = true
		for vertex := range dist {
			if edges[min][vertex] > 0 && dist[vertex] > edges[min][vertex] {
				dist[vertex] = edges[min][vertex]
			}
		}
	}
	return result
}
