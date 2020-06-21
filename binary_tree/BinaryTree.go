package binarytree

// 実装すべき機能
// -

type node struct {
	parent, right, left *node
}

type BinaryTree struct {
	r *node
}

// uを根とする二分技のサイズ(ノードの数)
func (u *node) size() int {
	return 1
}

//

//
func (u *node) depth() int {
	count := 0
	for {
		if u.parent != nil {
			count++
			u = u.parent
		} else {
			break
		}
	}
	return count
}
