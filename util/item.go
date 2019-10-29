package util

import (
	"fmt"
	"sort"
)

type Item interface {
	fmt.Stringer
}

type ComparableItem interface {
	fmt.Stringer
	sort.Interface
}
