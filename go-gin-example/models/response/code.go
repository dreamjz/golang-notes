package response

const (
	Success       = 10001
	Error         = 10002
	InvalidParams = 10003

	ErrorExistTag        = 20001
	ErrorNotExistTag     = 20002
	ErrorNotExistArticle = 20003

	ErrorAuthCheckTokenFail    = 30001
	ErrorAuthCheckTokenTimeout = 30002
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004

	DataBaseError = 40001
)
