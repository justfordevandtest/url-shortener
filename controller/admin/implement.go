package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/component/admin"
	"shorturl/presenter"
)

type Ctrl struct {
	Cmp admin.Comp
}

func New(cmp admin.Comp) (ctrl *Ctrl) {
	return &Ctrl{
		Cmp: cmp,
	}
}

// List godoc
// @Tags Admin
// @Summary List a page of URLs
// @Description Return a list of URLs according to the given paginator options
// @param Page query string true "A page number"
// @param PerPage query string true "A total number of items per page"
// @param Filters query string false "Condition for URLs retrieval, ex. 'Filters[id]=Lb'"
// @Produce json
// @Success 200 {object} presenter.SuccessResp{data=admin.ListOutput}
// @Failure 400 {object} presenter.ErrResp
// @Failure 401 {object} presenter.ErrResp
// @Failure 500 {object} presenter.ErrResp
// @Router /admin [get]
func (ctrl *Ctrl) List(c *gin.Context) {
	input := &admin.ListInput{}
	_ = c.ShouldBindQuery(input)

	list, err := ctrl.Cmp.List(input)
	if err != nil {
		errResp := presenter.MakeErrResp(err)
		c.JSON(errResp.Code, errResp)
		return
	}

	resp := presenter.MakeSuccessResp(http.StatusOK, list)
	c.JSON(resp.Code, resp)
}

// Delete godoc
// @Tags Admin
// @Summary Delete a URL with a given ID
// @Description Accessing a deleted URL will get a 410 response
// @param id path string true "URL ID"
// @Produce json
// @Success 200 {object} presenter.SuccessResp{data=admin.ListOutput}
// @Failure 400 {object} presenter.ErrResp
// @Failure 401 {object} presenter.ErrResp
// @Failure 404 {object} presenter.ErrResp
// @Failure 500 {object} presenter.ErrResp
// @Router /admin/{id} [delete]
func (ctrl *Ctrl) Delete(c *gin.Context) {
	input := &admin.DelInput{
		ID: c.Param("id"),
	}

	err := ctrl.Cmp.Delete(input)
	if err != nil {
		errResp := presenter.MakeErrResp(err)
		c.JSON(errResp.Code, errResp)
		return
	}

	resp := presenter.MakeSuccessResp(http.StatusOK, nil)
	c.JSON(resp.Code, resp)
}