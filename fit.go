package homotopy

import (
    "simplex/dp"
    . "simplex/struct/item"
    . "simplex/geom"
    . "simplex/relations"
    . "simplex/struct/sset"
    "fmt"
)

//find spatial fit
func (self *Homotopy) FindSpatialFit(i, j int) *SSet {
    var vset = NewSSet()

    vset.Add(Int(i))
    vset.Add(Int(j))

    var pln = NewLineString(self.pln[i: j + 1])
    fmt.Println("pln:", pln.WKT())

    var comparators = make([]Comparator, 0)
    for _, r := range self.relations {
        comparators = append(comparators, r.Relate(pln, self.neighbours))
    }

    var subpln = self.subgeom(vset)
    fmt.Println("pln:\t", subpln.WKT())

    var isvalid = self.isvalid(subpln, comparators)

    for self.intpool.HasNext() && !isvalid {
        var queue = self.intpool.Next().Clone()
        var fixint Item
        for !isvalid && !queue.IsEmpty() {
            idx := queue.Pop().(*dp.Vertex).Index()
            if fixint == nil {
                fixint = idx
            }
            vset.Add(idx)
            fmt.Println("idx:\t", self.pln[int(idx)].WKT())

            self.DeflectionFilter(vset, idx, fixint)

            subpln = self.subgeom(vset)
            fmt.Println("pln:\t", subpln.WKT())

            isvalid = self.isvalid(subpln, comparators)
        }
    }

    return vset
}


