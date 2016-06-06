package homotopy

import (
    "simplex/struct/heap"
    "simplex/struct/item"
    "simplex/dp"
    . "simplex/geom"
    . "simplex/relations"
    . "simplex/struct/sset"
)


//Homotopy
type Homotopy struct {
    pln        []*Point
    queue      *heap.Heap
    relates    []Relations
    neighbours []Geometry
}

//homotopy
func NewHomotopy(coords []*Point, queue *heap.Heap, relates []Relations, neighbours []Geometry) *Homotopy {
    return &Homotopy{coords, queue, relates, neighbours}
}

//find spatial fit
func (self *Homotopy) FindSpatialFit() *LineString {
    var n = len(self.pln) - 1
    var vset = NewSSet()

    vset.Add(item.Int(0))
    vset.Add(item.Int(n))

    var pln = NewLineString(self.pln)

    var comparators = make([]Comparator, 0)
    for _, rlt := range self.relates {
        comparators = append(comparators, rlt.Relate(pln, self.neighbours))
    }

    var subpln = self.subgeom(vset)
    var isvalid = self.isvalid(subpln, comparators)

    for !isvalid && self.queue.Size() > 0 {
        nextint := self.queue.Pop().(*dp.Vertex).Index()
        subpln = self.subgeom(vset.Add(item.Int(nextint)))
        isvalid = self.isvalid(subpln, comparators)
    }
    return subpln
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
    s.Each(func(o item.Item) {
        coords = append(coords, self.pln[int(o.(item.Int))])
    })
    return NewLineString(coords)
}


