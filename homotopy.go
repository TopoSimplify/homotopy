package homotopy

import (
	"github.com/TopoSimplify/ctx"
	"github.com/intdxdt/geom"
)

//Homotopy Relate
func Homotopy(coordinates []*geom.Point, contexts *ctx.ContextGeometries) bool {
	var bln bool
	var disjoint, _, xor = filterContext(coordinates, contexts)
	if xor.Len() > 0 {
		return false
	}

	if disjoint.Len() == 0 {
		return true
	}

	var ch = chainDeformation(coordinates, disjoint)
	var n = len(coordinates) - 1
	if ch.size == 2 {
		var a, b = ch.link.Point, ch.link.next.Point
		bln = coordinates[0].Equals2D(a) && coordinates[n].Equals2D(b)
	}
	return bln
}

func filterContext(coordinates []*geom.Point, contexts *ctx.ContextGeometries) (
	*ctx.ContextGeometries, *ctx.ContextGeometries, *ctx.ContextGeometries,
) {
	var n = len(coordinates)
	var simple = geom.NewLineString([]*geom.Point{
		coordinates[0], coordinates[n-1],
	})
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
