// Package lengthconv e2.2 长度转换程序
package lengthconv

import "fmt"

type KM float64
type M float64
type DM float64

func (km KM) String() string { return fmt.Sprintf("%gKM", km) }
func (m M) String() string   { return fmt.Sprintf("%gM", m) }
func (dm DM) String() string { return fmt.Sprintf("%gDM", dm) }
