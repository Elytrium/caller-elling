package methods

import (
	"fmt"
	"github.com/Elytrium/caller-elling/types"
	"github.com/Elytrium/elling/elling"
	"github.com/Elytrium/elling/routing"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/url"
	"strconv"
)

type Call struct{}

func (*Call) GetLimit() int {
	return 3
}

func (*Call) GetType() routing.MethodType {
	return routing.Http
}

func (*Call) IsPublic() bool {
	return false
}

func (*Call) Process(u *elling.User, v *url.Values) *routing.HTTPResponse {
	if !v.Has("number") {
		return routing.GenBadRequestResponse("no-number")
	}

	if !v.Has("method") {
		return routing.GenBadRequestResponse("no-method")
	}

	methodName := v.Get("method")
	method, ok := types.Instructions[methodName].(types.Method)

	if !ok {
		return routing.GenBadRequestResponse("no-such-method")
	}

	numStr := v.Get("number")

	if len(numStr) != 11 {
		return routing.GenBadRequestResponse("not-a-phone-number")
	}

	num, err := strconv.ParseInt(numStr, 10, 64)

	if err != nil {
		return routing.GenBadRequestResponse("not-a-phone-number")
	}

	var count int64
	elling.DB.Model(types.Mobile{}).Where("Number = ", num).Count(&count)

	if count != 0 {
		return routing.GenForbiddenResponse("number-already-set")
	}

	code := rand.Intn(10 * method.Length)
	codeStr := fmt.Sprintf("%0" + strconv.Itoa(method.Length) + "d", code)

	key := "caller:"+strconv.FormatUint(u.ID, 10)
	request := types.Request{
		Code:   codeStr,
		Number: "",
	}

	err = elling.Redis.Send(
		"HSET", redis.Args{}.Add(key).AddFlat(request))

	if err != nil {
		return routing.GenInternalServerError("save-set-failed")
	}

	err = elling.Redis.Send("EXPIRE", key, types.ModuleConfig.RequestExpiration)

	if err != nil {
		return routing.GenInternalServerError("save-exp-failed")
	}

	_, err = method.CallRequest.DoRequest(map[string]string{
		"code": codeStr,
	})

	if err != nil {
		log.Error().Err(err).Msg("Calling")
		return routing.GenInternalServerError("call-failed")
	}

	return routing.GenSuccessResponse("ok")
}
