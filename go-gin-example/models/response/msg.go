package response

var msgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "request parameters error",
	ErrorExistTag:              "tag already exists",
	ErrorNotExistTag:           "tag not exists",
	ErrorNotExistArticle:       "article not exists",
	ErrorAuthCheckTokenFail:    "check token failed",
	ErrorAuthCheckTokenTimeout: "token timeout",
	ErrorAuthToken:             "token generate failed",
	ErrorAuth:                  "token error",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[Error]
}
