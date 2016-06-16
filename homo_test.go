package homotopy

import (
    "fmt"
    . "github.com/franela/goblin"
    . "simplex/geom"
    . "simplex/relations"
    "simplex/dp"
    "simplex/struct/heap"
    . "simplex/interest"
    "testing"
    "simplex/struct/item"
)

func TestHomotopy(t *testing.T) {

    g := Goblin(t)

    var A = NewPointXY(2, -3)
    var B = NewPointXY(7, 2)
    var C = NewPointXY(9, -2)
    var D = NewPointXY(13, 2)
    var E = NewPointXY(16, -3)
    var F = NewPointXY(13, -4)
    var G = NewPointXY(17, -5)
    var H = NewPointXY(13, -8)

    //var a = dp.NewVObj(0, 0.0)
    var b = dp.NewVObj(1, 6.62)
    var c = dp.NewVObj(2, 3.81)
    var d = dp.NewVObj(3, 9.10)
    var e = dp.NewVObj(4, 5.83)
    var f = dp.NewVObj(5, 3.64)
    var _g = dp.NewVObj(6, 5.00)
    //var h = dp.NewVObj(7, 0.0)

    g.Describe("Homotopy", func() {
        gA := NewPolygonFromWKT("POLYGON (( 5.98753212697853 -4.812514603172059, 5.98753212697853 -3.9952946175296855, 6.334231514826809 -3.9952946175296855, 6.334231514826809 -4.812514603172059, 5.98753212697853 -4.812514603172059 ))")
        //gB := NewPolygonFromWKT("POLYGON (( 6.951548493807859 -2.7181585575501734, 6.951548493807859 -2.3219306857235678, 7.57065454353693 -2.3219306857235678, 7.57065454353693 -2.7181585575501734, 6.951548493807859 -2.7181585575501734 ))")
        var neighbours = []Geometry{gA}
        fmt.Println(gA.BBox())

        var ln = NewLineString([]*Point{A, B, C, D, E, F, G, H})
        fmt.Println("Polyline: ", ln)

        hp := heap.NewHeap(heap.NewHeapType().AsMax())
        hp.Push(b); hp.Push(c); hp.Push(d); hp.Push(e); hp.Push(f); hp.Push(_g);
        var intfn = func() *heap.Heap {
            return hp
        }
        var intpool = NewIntCandidates([]IntFunctor{intfn})

        grlt := NewGeometryRelate()
        qrlt := NewQuadRelate()
        drlt := NewMinDistanceRelate(20.5)

        relates := []Relations{grlt, qrlt, drlt}

        var pln = ln.Coordinates()
        var topy = NewHomotopy(pln, intpool, relates, neighbours)
        var ints = topy.FindSpatialFit(0, len(pln) - 1)
        for _, i := range []item.Int{0, 1, 3, 7} {
            g.Assert(ints.Contains(i)).IsTrue()
        }

        fmt.Println(pln)
        topy.UpdateHomotopy(pln, intpool, relates, neighbours)
        ints = topy.FindSpatialFit(0, len(pln) - 1)
        fmt.Println(ints)
        //for _, i := range []item.Int{0, 1, 3, 7} {
        //    g.Assert(ints.Contains(i)).IsTrue()
        //}
    })
}
