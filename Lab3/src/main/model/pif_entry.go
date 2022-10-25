package model

import "fmt"

type PIFEntry struct {
	Token string
	Type  uint8
	Hash  int
	Index int
}

func (e *PIFEntry) String() string {

	return fmt.Sprintf(`PIFEntry{ Token="%s"; Type="%d"; Hash=%d; Index=%d }`, e.Token, e.Type, e.Hash, e.Index)
}
