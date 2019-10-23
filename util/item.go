package util

import (
	"fmt"
	"sort"
)

type Item interface {
	fmt.Stringer
}

type SortableItem interface {
	fmt.Stringer
	sort.Interface
}
