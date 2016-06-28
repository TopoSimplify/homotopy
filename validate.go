package homotopy

import (
    . "simplex/geom"
    . "simplex/relations"
)

//is valid
func (self *Homotopy) isvalid(g Geometry, comparators []Comparator) bool {
    //make true , proof otherwise
    var bln = true
    for i := 0; bln && i < len(comparators); i++ {
        bln = bln && comparators[i](g)
    }
    return bln
}
