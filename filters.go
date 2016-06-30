package homotopy
import (
    . "simplex/struct/item"
    . "simplex/geom"
    . "simplex/struct/sset"
)

func (self *Homotopy) DeflectionFilter(vset *SSet, idx, fixint Item){
    b := vset.PrevItem(idx)
    f := vset.NextItem(idx)
    var nb , nf, ni []Item

    ni = []Item{b, idx, f};

    next_f := vset.PrevItem(f)
    prev_b := vset.PrevItem(b)

    if prev_b  != nil {
        nb = []Item{prev_b, b , vset.NextItem(b)}
    }

    if next_f  != nil {
        nf = []Item{vset.NextItem(f), f , next_f}
    }

    if len(nb) != 0 {

    }

    if len(ni) != 0 {

    }

    if len(nf) != 0 {

    }

    var rmlist = []Item
    //note this ordering b i f is important for _rmsubrange
    if (angb && angb > self._3vdefln) { //b
      rmlist.push(b);
    }
    if (angi && angi > self._3vdefln) { //i
      rmlist.push(i);
    }
    if (angf && angf > self._3vdefln) { //f
      rmlist.push(f);
    }
    if (rmlist.length) {
      this._rmsubrange(rmlist, subrange, fixint);
    }
}
