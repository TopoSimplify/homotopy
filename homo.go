package homotopy

import (
    . "simplex/geom"
    . "simplex/relations"
    . "simplex/interest"
)


//Homotopy
type Homotopy struct {
    pln        []*Point
    intpool    *IntCandidates
    relations  []Relations
    neighbours []Geometry
}

//homotopy
func NewHomotopy(coords []*Point, intpool *IntCandidates, relations []Relations,
neighbours []Geometry) *Homotopy {
    return &Homotopy{coords, intpool, relations, neighbours}
}




