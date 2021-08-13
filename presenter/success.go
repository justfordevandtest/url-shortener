package presenter

const (
	okStatus       = "OK"
)

type SuccessResp struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
} // @Name SuccessResponse

func MakeSuccessResp(status int, data interface{}) (resp *SuccessResp) {
	return &SuccessResp{
		Status: okStatus,
		Code:   status,
		Data:   data,
	}
}
