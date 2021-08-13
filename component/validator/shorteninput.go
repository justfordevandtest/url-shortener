package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"shorturl/component/shortener"
)

func (v *GoPlayGroundValidator) ShortenInputValidation(sl validator.StructLevel) {
	blacklist, err := v.blacklistRepo.List()
	if err != nil {
		reportURLErr(sl.Current().Interface().(shortener.ShortenInput).URL, "fetch-blacklist", sl)
	}

	for _, regex := range blacklist {
		match, err := regexp.MatchString(regex, sl.Current().Interface().(shortener.ShortenInput).URL)
		if err != nil {
			reportURLErr(sl.Current().Interface().(shortener.ShortenInput).URL, "invalid-blacklist", sl)
			continue
		}
		if match {
			reportURLErr(sl.Current().Interface().(shortener.ShortenInput).URL, "blacklist", sl)
		}
	}
}

func reportURLErr(url string, tag string, sl validator.StructLevel) {
	sl.ReportError(url, "URL", "URL", tag, "")
}
