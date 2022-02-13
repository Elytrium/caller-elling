package methods

import (
	"github.com/Elytrium/caller-elling/types"
	"github.com/Elytrium/elling/elling"
	"github.com/Elytrium/elling/routing"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
	"net/url"
	"strconv"
)

type Verify struct{}

func (*Verify) GetLimit() int {
	return 3
}

func (*Verify) GetType() routing.MethodType {
	return routing.Http
}

func (*Verify) IsPublic() bool {
	return false
}

func (*Verify) Process(u *elling.User, v *url.Values) *routing.HTTPResponse {
	if !v.Has("code") {
		return routing.GenBadRequestResponse("no-code")
	}

	key := "caller:" + strconv.FormatUint(u.ID, 10)
	requestBytes, err := redis.Values(elling.Redis.Do("HGETALL", key))

	if err != nil {
		log.Error().Err(err).Msg("Getting call request from Redis")
		return routing.GenInternalServerError("get-failed")
	}

	if len(requestBytes) == 0 {
		return routing.GenBadRequestResponse("no-code-for-user")
	}

	var request types.Request
	err = redis.ScanStruct(requestBytes, &request)

	if err != nil {
		log.Error().Err(err).Msg("Scanning call request")
		return routing.GenInternalServerError("scan-failed")
	}

	codeStr := v.Get("code")

	if codeStr == request.Code {
		num, err := strconv.ParseInt(request.Number, 10, 64)

		if err != nil {
			log.Error().Err(err).Msg("Parsing phone number")
			return routing.GenInternalServerError("phone-parse-failed")
		}

		mobile := types.Mobile{
			Number: num,
			UserID: u.ID,
		}

		mobile.Save()
	}

	return routing.GenSuccessResponse("ok")
}
