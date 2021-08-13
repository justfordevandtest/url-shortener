package shortener

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/component/shortener"
	"shorturl/presenter"
)

type Ctrl struct {
	Cmp shortener.Comp
}

func New(cmp shortener.Comp) (ctrl *Ctrl) {
	return &Ctrl{
		Cmp: cmp,
	}
}

// Shorten godoc
// @Tags Public
// @Summary Shorten a given URL
// @Description Return shorten version of a given URL
// @Param input body shortener.ShortenInput true "URL will not expire if 'expired' is set to null or excluded"
// @Accept json
// @Produce json
// @Success 200 {object} presenter.SuccessResp{data=shortener.ShortenOutput}
// @Failure 400 {object} presenter.ErrResp
// @Failure 410 {object} presenter.ErrResp
// @Failure 500 {object} presenter.ErrResp
// @Router /shorten [post]
func (ctrl *Ctrl) Shorten(c *gin.Context) {
	input := &shortener.ShortenInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		errResp := presenter.MakeErrResp(err)
		errResp.Code = http.StatusBadRequest
		c.JSON(errResp.Code, errResp)
		return
	}

	shortenURL, err := ctrl.Cmp.ShortenURL(input)
	if err != nil {
		errResp := presenter.MakeErrResp(err)
		c.JSON(errResp.Code, errResp)
		return
	}

	resp := presenter.MakeSuccessResp(http.StatusOK, shortenURL)
	c.JSON(resp.Code, resp)
}

// Access godoc
// @Tags Public
// @Summary Access a given shortened URL
// @Description Return a decoded URL of a given shortened URL
// @param id path string true "ID"
// @Produce json
// @Header 302 {string} url {https://example.com}
// @Failure 410 {object} presenter.ErrResp
// @Failure 500 {object} presenter.ErrResp
// @Router /:id [get]
func (ctrl *Ctrl) Access(c *gin.Context) {
	input := &shortener.AccessInput{
		ID: c.Param("id"),
	}

	url, err := ctrl.Cmp.AccessURL(input)
	if err != nil {
		errResp := presenter.MakeErrResp(err)
		c.JSON(errResp.Code, errResp)
		return
	}

	c.Redirect(http.StatusFound, url.URL)
}
