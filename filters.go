package homotopy
import (
    . "simplex/struct/item"
    . "simplex/geom"
    . "simplex/struct/sset"
)

func (self *Homotopy) DeflectionFilter(vset *SSet, idx Item){
    b := vset.PrevItem(idx)
    f := vset.NextItem(idx)

}
