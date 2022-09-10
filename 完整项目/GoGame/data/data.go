package data

import (
	"GoGame/def"
	"fmt"
)

var PlayerMap map[string]def.Player

func init() {
	PlayerMap = make(map[string]def.Player)
	if PlayerMap == nil {
		fmt.Errorf("PlaterMap Init Error")
	}
}
