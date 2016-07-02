package homotopy
import (
    . "simplex/struct/item"
    "simplex/geom"
    . "simplex/struct/sset"
    "simplex/vect"
)

func (self *Homotopy) DeflectionFilter(vset *SSet, idx, fixint Item){
    b := vset.PrevItem(idx)
    f := vset.NextItem(idx)
    var nb , nf, ni []Item
    var angb , angi , angf float64

    ni = []Item{b, idx, f};

    next_f := vset.NextItem(f)
    prev_b := vset.PrevItem(b)

    if prev_b  != nil {
        nb = []Item{prev_b, b , vset.NextItem(b)}
    }

    if next_f  != nil {
        nf = []Item{vset.PrevItem(f), f , next_f}
    }

    if len(nb) != 0 {
        angb = self.pln[nb[1]].AngleAtPoint(self.pln[nb[0]], self.pln[nb[2]])
    }

    if len(ni) != 0 {
        angi = self.pln[ni[1]].AngleAtPoint(self.pln[ni[0]], self.pln[ni[2]])
    }

    if len(nf) != 0 {
        angf = self.pln[nf[1]].AngleAtPoint(self.pln[nf[0]], self.pln[nf[2]])
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
