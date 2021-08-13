package entity

import "fmt"

var (
	UnknownErr = defineErr("UNKNOWN", "UNKNOWN")

	UnauthorizedErr   = defineErr("UNAUTHORIZED", "CREDENTIAL")
	NotFoundRecordErr = defineErr("NOTFOUND", "RECORD")
	ExpiredURLErr     = defineErr("EXPIRED", "URL")

	ValidatorShortenErr = defineErr("VALIDATOR", "SHORTEN")
	ValidatorListErr    = defineErr("VALIDATOR", "LIST")

	ListRecordsErr  = defineErr("LIST", "RECORD")
	CreateRecordErr = defineErr("CREATE", "RECORD")
	ReadRecordErr   = defineErr("READ", "RECORD")
	UpdateRecordErr = defineErr("UPDATE", "RECORD")
	DeleteRecordErr = defineErr("DELETE", "RECORD")
	CountRecordErr  = defineErr("COUNT", "RECORD")

	CacheWriteErr  = defineErr("CACHE", "WRITE")
)

type Error struct {
	Cause   error  `json:"cause"`
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
}

type ErrorFunc func(cause error) (err *Error)

func defineErr(code string, subCode string) (errFn ErrorFunc) {
	return func(cause error) (err *Error) {
		return &Error{
			Cause:   cause,
			Code:    code,
			SubCode: subCode,
		}
	}
}

func TypeOfErr(err error) (typeOf *Error) {
	if e0, ok := err.(*Error); ok {
		return e0
	}
	return &Error{Cause: err}
}

func (e *Error) Error() (msg string) {
	causeStr := ""
	if e.Cause != nil {
		causeStr = e.Cause.Error()
	}
	return fmt.Sprintf("%s (%s:%s)", causeStr, e.Code, e.SubCode)
}

func (e *Error) IsType(errFn ErrorFunc) (is bool) {
	test := errFn(nil)
	return e.Code == test.Code && e.SubCode == test.SubCode
}
