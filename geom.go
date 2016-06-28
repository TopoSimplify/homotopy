package homotopy

import (
    . "simplex/struct/item"
    . "simplex/geom"
    . "simplex/struct/sset"
)

//sub geometry
func (self *Homotopy)  subgeom(s *SSet) *LineString {
    coords := make([]*Point, 0)
    s.Each(func(o Item) {
        coords = append(coords, self.pln[int(o.(Int))])
    })
    return NewLineString(coords)
}

