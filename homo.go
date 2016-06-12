package homotopy

import (
    "simplex/dp"
    . "simplex/struct/item"
    . "simplex/geom"
    . "simplex/relations"
    . "simplex/struct/sset"
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

func (self *Homotopy) UpdateHomotopy(coords []*Point, intpool *IntCandidates,
relations []Relations, neighbours []Geometry) *Homotopy {
    self.pln        = coords
    self.intpool    = intpool
    self.relations  = relations
    self.neighbours = neighbours
    return self
}

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

    for self.intpool.HasNext() && !isvalid{
        var queue = self.intpool.Next()

        for !isvalid && !queue.IsEmpty() {
            vset.Add(queue.Pop().(*dp.Vertex).Index())
            subpln = self.subgeom(vset)
            isvalid = self.isvalid(subpln, comparators)
        }
    }

    return vset
}

//is valid
func (self *Homotopy) isvalid(g Geometry, comparators []Comparator) bool {
    //make true , proof otherwise
    var bln = true
    for i := 0; bln && i < len(comparators); i++ {
        bln = bln && comparators[i](g)
    }
    return bln
}

//sub geometry
func (self *Homotopy)  subgeom(s *SSet) *LineString {
    coords := make([]*Point, 0)
    s.Each(func(o Item) {
        coords = append(coords, self.pln[int(o.(Int))])
    })
    return NewLineString(coords)
}


