package constant

const (
	ResponseMsgOK            = "OK"
	ResponseMsgErrorNotFound = "no book found"
	ResponseMsgErrorInternal = "Internal Server Error"
	ResponseMsgBadRequest    = "Bad Request"

	ResponseMsgBookAlreadyExists = "book already exists"
	ResponseMsgUserAlreadyExists = "user already exists"

	ResponseMsgAuthorDoesNotExist  = "author does not exist"
	ResponseMsgUserDoesNotExist    = "user does not exist"
	ResponseMsgBookDoesNotExist    = "book does not exist"
	ResponseMsgRecordDoesNotExist  = "record does not exist"
	ResponseMsgBookAlreadyReturned = "book already returned"

	ResponseMsgErrorCredentials = "email or password incorrect"

	ResponseMsgSQLError = "bad request database error"
)
