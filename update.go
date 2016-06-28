package homotopy

import (
    . "simplex/geom"
    . "simplex/relations"
    . "simplex/interest"
)


func (self *Homotopy) UpdateHomotopy(coords []*Point, intpool *IntCandidates,
relations []Relations, neighbours []Geometry) *Homotopy {
    self.pln        = coords
    self.intpool    = intpool
    self.relations  = relations
    self.neighbours = neighbours
    return self
}
