package methods

import (
	"github.com/Elytrium/caller-elling/types"
	"github.com/Elytrium/elling/elling"
	"github.com/Elytrium/elling/routing"
	"net/url"
)

type Get struct{}

func (*Get) GetLimit() int {
	return 10
}

func (*Get) GetType() routing.MethodType {
	return routing.Http
}

func (*Get) IsPublic() bool {
	return false
}

func (*Get) Process(u *elling.User, _ *url.Values) *routing.HTTPResponse {
	var mobile types.Mobile
	elling.DB.Find(&mobile, u.ID)

	return routing.GenSuccessResponse(mobile.Number)
}
