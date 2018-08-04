package homotopy

import (
	"github.com/TopoSimplify/ctx"
	"github.com/intdxdt/geom"
)

//Homotopy Relate
func Homotopy(coordinates geom.Coords, contexts *ctx.ContextGeometries) bool {
	var bln bool
	var disjoint, _, xor = filterContext(coordinates, contexts)
	if xor.Len() > 0 {
		return false
	}

	if disjoint.Len() == 0 {
		return true
	}

	var ch = chainDeformation(coordinates, disjoint)
	var n = coordinates.Len() - 1
	if ch.size == 2 {
		var a, b = ch.link.Point, ch.link.next.Point
		bln = coordinates.Pt(0).Equals2D(a) && coordinates.Pt(n).Equals2D(b)
	}
	return bln
}

func filterContext(coordinates geom.Coords, contexts *ctx.ContextGeometries) (
	*ctx.ContextGeometries, *ctx.ContextGeometries, *ctx.ContextGeometries,
) {
	var n = coordinates.Len()
	var simpleCoords = coordinates
	simpleCoords.Idxs = []int{0, n-1}
	var simple = geom.NewLineString(simpleCoords)
	var ln = geom.NewLineString(coordinates)
	var exclude = ctx.NewContexts()
	var disjoint = ctx.NewContexts()
	var xor = ctx.NewContexts()

	for _, c := range contexts.DataView() {
		ln_bln := ln.Intersects(c.Geom)
		simple_bln := simple.Intersects(c.Geom)
		if ln_bln && simple_bln {
			exclude.Push(c)
		} else if !ln_bln && !simple_bln {
			disjoint.Push(c)
		} else {
			xor.Push(c)
		}
	}
	return disjoint, exclude, xor
}
