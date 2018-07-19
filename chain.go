package homotopy

import (
	"github.com/intdxdt/geom"
)

type Vertex struct {
	Point *geom.Point
	i     int
	prev  *Vertex
	next  *Vertex
}

type Chain struct {
	link *Vertex
	size int
	simple *geom.LineString
}

func NewChain(coordinates []geom.Point) *Chain {
	var n = len(coordinates)
	var chain = &Chain{
		size: n,
		link: &Vertex{Point: &coordinates[0]},
		simple: geom.NewLineString([]geom.Point{
			coordinates[0], coordinates[n-1],
		}),
	}
	var prev, next *Vertex
	prev = chain.link
	for i := 1; i < chain.size; i++ {
		next = &Vertex{Point: &coordinates[i], i: i, prev: prev}
		ptrs(prev, next)
		prev = next
	}
	return chain
}

func remove(link *Vertex) {
	if link != nil {
		ptrs(link.prev, link.next)
	}
}

//Updates next and prev pointers
func ptrs(prev, next *Vertex) {
	if next != nil {
		next.prev = prev
	}
	if prev != nil {
		prev.next = next
	}
}
