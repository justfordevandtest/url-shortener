package shortener

import (
	"fmt"
	"math/rand"
	"shorturl/entity"
	"time"
)

type impl struct {
	baseURL        string
	cacheThreshold int
	urlRepo        URLRepo
	urlCache       URLCache
	validator      Validator
}

func New(baseURL string, cacheThreshold int, urlRepo URLRepo, urlCache URLCache, v Validator) (comp Comp) {
	return &impl{
		baseURL:        baseURL,
		cacheThreshold: cacheThreshold,
		urlRepo:        urlRepo,
		urlCache:       urlCache,
		validator:      v,
	}
}

func (i *impl) ShortenURL(input *ShortenInput) (output *ShortenOutput, err error) {
	err = i.validator.Validate(input)
	if err != nil {
		return nil, entity.ValidatorShortenErr(err)
	}

	var ent *entity.URL
	used := true
	for used {
		ent = entity.MakeURL(rand.Uint64(), input.URL, input.Expired)
		cnt, err := i.urlRepo.CountID(ent.ID)
		if err != nil {
			return nil, entity.CountRecordErr(err)
		}
		used = cnt > 0
	}

	err = i.urlRepo.Create(ent)
	if err != nil {
		return nil, entity.CreateRecordErr(err)
	}

	output = &ShortenOutput{
		URL: fmt.Sprintf("%s/%s", i.baseURL, ent.ID),
	}
	return output, nil
}

func (i *impl) AccessURL(input *AccessInput) (output *ShortenOutput, err error) {
	url := i.urlCache.Read(input.ID)
	if url == nil {
		url, err = i.urlRepo.Read(input.ID)
		if err != nil {
			return nil, entity.NotFoundRecordErr(err)
		}
	}

	if url.Expired != nil && time.Now().Unix() > *url.Expired {
		return nil, entity.ExpiredURLErr(nil)
	}

	err = i.urlRepo.IncrHit(input.ID)
	if err != nil {
		return nil, entity.UpdateRecordErr(err)
	}

	url.HitCount++
	if url.HitCount >= i.cacheThreshold {
		err = i.urlCache.Write(url)
		if err != nil {
			return nil, entity.CacheWriteErr(err)
		}
	}

	return &ShortenOutput{URL: url.URL}, nil
}
