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

    var comparators = make([]Comparator, 0)
    for _, r := range self.relations {
        comparators = append(comparators, r.Relate(pln, self.neighbours))
    }

    var subpln = self.subgeom(vset)
    var isvalid = self.isvalid(subpln, comparators)

    fmt.Println("vset       --|>", vset.String())
    fmt.Println("vset geom  --|>", subpln.WKT())
    fmt.Println("bool state --|>", isvalid)

    for self.intpool.HasNext() && !isvalid{
        var queue = self.intpool.Next()

        for !isvalid && !queue.IsEmpty() {
            idx:= queue.Pop().(*dp.Vertex).Index()
            vset.Add(idx)
            subpln = self.subgeom(vset)
            isvalid = self.isvalid(subpln, comparators)

            fmt.Println("vset       --|>", vset.String())
            fmt.Println("vset geom  --|>", subpln.WKT())
            fmt.Println("bool state --|>", isvalid)
            fmt.Println(vset.PrintBaseStruct())
        }
    }

    return vset
}

