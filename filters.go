package homotopy

import (
    . "simplex/struct/item"
    . "simplex/struct/sset"
    . "simplex/geom"
)

func (self *Homotopy) DeflectionFilter(vset *SSet, idx, fixint Item) {
    b := vset.PrevItem(idx)
    f := vset.NextItem(idx)
    var nb, nf, ni []Item
    var angb, angi, angf float64

    ni = []Item{b, idx, f};

    next_f := vset.NextItem(f)
    prev_b := vset.PrevItem(b)

    if prev_b != nil {
        nb = []Item{prev_b, b, vset.NextItem(b)}
    }

    if next_f != nil {
        nf = []Item{vset.PrevItem(f), f, next_f}
    }

    if len(nb) != 0 {
        angb = self.Pt(nb[1]).AngleAtPoint(self.Pt(nb[0]), self.Pt(nb[2]))
    }

    if len(ni) != 0 {
        angi = self.Pt(ni[1]).AngleAtPoint(self.Pt(ni[0]), self.Pt(ni[2]))
    }

    if len(nf) != 0 {
        angf = self.Pt(nf[1]).AngleAtPoint(self.Pt(nf[0]), self.Pt(nf[2]))
    }

    _3vdefln := 3.1
    var rmlist = make([]Item, 0)
    //note this ordering b i f is important for _rmsubrange
    if (angb > _3vdefln) {
        //b
        rmlist = append(rmlist, b)
    }
    if (angi > _3vdefln) {
        //i
        rmlist = append(rmlist, idx)
    }
    if (angf > _3vdefln) {
        //f
        rmlist = append(rmlist, f)
    }
    if (len(rmlist) > 0) {
        _rmsubrange(rmlist, vset, fixint);
    }
}

func (self *Homotopy) Pt(i Item) *Point{
    return self.pln[int(i.(Int))]
}

func _rmsubrange(indexlist []Item, vset *SSet, fixint Item) {
    for i := 0; i < len(indexlist); i++ {
        item := indexlist[i];
        if (item.Compare(fixint) != 0) {
            vset.Remove(item);
        }
    }
};