package caller_elling

import (
	"github.com/Elytrium/caller-elling/methods"
	"github.com/Elytrium/caller-elling/types"
	"github.com/Elytrium/elling/basic/common"
	"github.com/Elytrium/elling/elling"
	"github.com/Elytrium/elling/module"
	"github.com/Elytrium/elling/routing"
	"github.com/rs/zerolog/log"
	"reflect"
)

type Caller struct{}

type UserListener struct {}

type MasterListener struct {}

func (*Caller) OnModuleInit() {
	types.Instructions = common.ReadInstructions("caller", reflect.TypeOf(types.Method{}))
	types.LoadConfig()
	elling.RegisterListener(UserListener{})
}

func (*Caller) OnModuleRemove() {}

func (*Caller) GetMeta() *module.Meta {
	return &module.Meta{
		Name: "caller",
		Routes: map[string]routing.Method{
			"call": &methods.Call{},
			"verify": &methods.Verify{},
			"get": &methods.Get{},
		},
		DatabaseFields: []interface{}{},
	}
}

func (*UserListener) OnUserCreation(e elling.UserCreationEvent) {
	err := e.User.Deactivate()
	if err != nil {
		log.Error().Err(err).Msg("Deactivating before phone verification")
	}
}

var Module Caller
