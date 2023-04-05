package redblack

import "github.com/playgroundgo/genlib/generic"

type nodeColor uint8

const (
	black nodeColor = iota
	red
)

type node[K, V any] struct {
	key      K
	value    V
	children [2]*node[K, V]
	parent   *node[K, V]
}

// Tree implements a red-black tree.
type Tree[K, V any] struct {
	root  *node[K, V]
	less  generic.LessFn[K]
	size  uint32
	color nodeColor
}

// New returns an empty red-black tree.
func New[K, V any](less generic.LessFn[K]) *Tree[K, V] {
	return &Tree[K, V]{
		less: less,
	}
}

func (t *Tree[K, V]) IsEmpty() bool {
	return t.root == nil
}

func (t *Tree[K, V]) Size() uint32 {
	return t.size
}

func (t *Tree[K, V]) Find(key K) (*V, bool) {
	if t.root == nil {
		return nil, false
	}

	node := t.root

	for node != nil {
		compare := generic.CompareBy(key, node.key, t.less)
		switch {
		case compare < 0:
			node = node.children[0]
		case compare > 0:
			node = node.children[1]
		default:
			return &node.value, true
		}
	}
	return nil, false
}

func (t *Tree[K, V]) getSibling()
