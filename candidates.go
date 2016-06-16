package homotopy

//func (self *Homotopy) _childints(n *bst.Node) *heap.Heap {
//    var node = self.AsDPNode(n)
//    var stack = stack.NewStack()
//    var nextint = node.Ints.Peek().(*dp.Vertex).Index()
//    var intlist = heap.NewHeap(heap.NewHeapType().AsMax())
//    var intobj    *dp.Vertex
//
//    //node stack
//    if n.Right != nil {
//        stack.Add(n.Right)
//
//    }
//    if n.Left != nil {
//        stack.Add(n.Left)
//    }
//
//    for !stack.IsEmpty() {
//        node = self.AsDPNode_BSTNode_Any(stack.Pop())
//        intobj = node.Ints.Peek().(*dp.Vertex)
//
//        if nextint != intobj.Index() && intobj.Value() > 0.0 {
//            intlist.Push(intobj)
//        }
//        if n.Right != nil {
//            stack.Add(n.Right)
//        }
//        if n.Left != nil {
//            stack.Add(n.Left)
//        }
//    }
//
//    return intlist
//}
