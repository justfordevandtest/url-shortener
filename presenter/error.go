package presenter

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"shorturl/entity"
)

const errorStatus = "ERROR"

type ErrResp struct {
	Status string     `json:"status"`
	Code   int        `json:"code"`
	Errors []*ErrItem `json:"errors"`
} // @Name ErrorResponse

type ErrItem struct {
	Cause   string `json:"-"`
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
} // @Name ErrorItemResponse

func MakeErrResp(err error) (errResp *ErrResp) {
	return &ErrResp{
		Status: errorStatus,
		Code:   getHTTPStatusCode(err),
		Errors: getRespErrors(err),
	}
}

func getHTTPStatusCode(err error) (code int) {
	switch err := entity.TypeOfErr(err); {
	case err.IsType(entity.ExpiredURLErr):
		return http.StatusGone
	case err.IsType(entity.ValidatorShortenErr):
		return http.StatusUnprocessableEntity
	case err.IsType(entity.ReadRecordErr):
		return http.StatusNotFound
	case err.IsType(entity.NotFoundRecordErr):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func getRespErrors(err error) (errs []*ErrItem) {
	switch err.(type) {
	case *entity.Error:
		return utilToResp(err.(*entity.Error))
	default:
		ukErr := entity.UnknownErr(err)
		return []*ErrItem{
			{
				Cause:   ukErr.Error(),
				Code:    ukErr.Code,
				SubCode: ukErr.SubCode,
			},
		}
	}
}

func utilToResp(err *entity.Error) (errs []*ErrItem) {
	switch err := entity.TypeOfErr(err); {
	case err.IsType(entity.ValidatorShortenErr):
		return validateToResp(err)
	default:
		return []*ErrItem{
			{
				Cause:   err.Error(),
				Code:    err.Code,
				SubCode: err.SubCode,
			},
		}
	}
}

func validateToResp(err *entity.Error) (errs []*ErrItem) {
	vErrs := err.Cause.(validator.ValidationErrors)
	errs = make([]*ErrItem, len(vErrs))
	for i, vErr := range vErrs {
		errs[i] = &ErrItem{
			Cause:   vErr.Translate(nil),
			Code:    vErr.Tag(),
			SubCode: vErr.Field(),
		}
	}

	return errs
}