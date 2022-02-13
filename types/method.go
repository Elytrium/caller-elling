package types

import (
	"github.com/Elytrium/elling/basic/common"
	"github.com/Elytrium/elling/elling"
)

var Instructions common.Instructions

type Method struct {
	Name        string            `yaml:"name"`
	Length      int               `yaml:"length"`
	CallRequest elling.NetRequest `yaml:"call-request"`
}
