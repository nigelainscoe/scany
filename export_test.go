package dbscan

import (
	"reflect"
)

var ParseDestination = parseDestination

func (rs *RowScanner) SetStartFn(f startRowsFunc)         { rs.startFn = f }
func (rs *RowScanner) Start(dstValue reflect.Value) error { return rs.start(dstValue) }
