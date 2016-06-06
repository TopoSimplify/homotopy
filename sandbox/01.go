package main

import (
    "fmt"
    //"simplex/vect"
    . "simplex/geom"
    . "simplex/relations"
    "simplex/dp"
    "simplex/struct/heap"
    . "simplex/homotopy"
)

var A = NewPointXY(2, -3)
var B = NewPointXY(7, 2)
var C = NewPointXY(9, -2)
var D = NewPointXY(13, 2)
var E = NewPointXY(16, -3)
var F = NewPointXY(13, -4)
var G = NewPointXY(17, -5)
var H = NewPointXY(13, -8)

var a = dp.NewVObj(0, 0.0)
var b = dp.NewVObj(1, 6.62)
var c = dp.NewVObj(2, 3.81)
var d = dp.NewVObj(3, 9.10)
var e = dp.NewVObj(4, 5.83)
var f = dp.NewVObj(5, 3.64)
var g = dp.NewVObj(6, 5.00)
var h = dp.NewVObj(7, 0.0)

func main() {
    gA := NewPolygonFromWKT("POLYGON (( 5.98753212697853 -4.812514603172059, 5.98753212697853 -3.9952946175296855, 6.334231514826809 -3.9952946175296855, 6.334231514826809 -4.812514603172059, 5.98753212697853 -4.812514603172059 ))")
    //gB := NewPolygonFromWKT("POLYGON (( 6.951548493807859 -2.7181585575501734, 6.951548493807859 -2.3219306857235678, 7.57065454353693 -2.3219306857235678, 7.57065454353693 -2.7181585575501734, 6.951548493807859 -2.7181585575501734 ))")
    var neighbours = []Geometry{gA}
    fmt.Println(gA.BBox())

    var ln = NewLineString([]*Point{A, B, C, D, E, F, G, H})
    fmt.Println("Polyline: ", ln)



    hp := heap.NewHeap(heap.NewHeapType().AsMax())
    hp.Push(b); hp.Push(c); hp.Push(d); hp.Push(e); hp.Push(f); hp.Push(g);

    //grlt := NewGeometryRelate()
    //qrlt := NewQuadRelate()
    drlt := NewMinDistanceRelate(20.5)

    relates := []Relations{drlt}

    var topy = NewHomotopy(ln.Coordinates(), hp, relates, neighbours)
    var poly = topy.FindSpatialFit()


    fmt.Println(poly)

}

