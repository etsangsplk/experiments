package util

import (
	"fmt"
	"sort"
)

type Item interface {
	fmt.Stringer
	sort.Interface
}
