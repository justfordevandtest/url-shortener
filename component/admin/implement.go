package admin

import (
	"shorturl/entity"
)

type impl struct {
	urlRepo   URLRepo
	validator Validator
}

func New(urlRepo URLRepo, v Validator) (comp Comp) {
	return &impl{
		urlRepo:   urlRepo,
		validator: v,
	}
}

func (i *impl) List(input *ListInput) (output *ListOutput, err error) {
	err = i.validator.Validate(input)
	if err != nil {
		return nil, entity.ValidatorListErr(err)
	}

	filters := make(map[string]interface{})
	if len(input.ID) > 0 {
		filters["ID"] = input.ID
	}
	if len(input.Keyword) > 0 {
		filters["URL"] = input.Keyword
	}

	total, items, err := i.urlRepo.List(input.Page, input.PerPage, filters)
	if err != nil {
		return nil, entity.ListRecordsErr(err)
	}

	return &ListOutput{
		Total: total,
		Items: items,
	}, nil
}

func (i *impl) Delete(input *DelInput) (err error) {
	err = i.validator.Validate(input)
	if err != nil {
		return entity.ValidatorListErr(err)
	}

	url, err := i.urlRepo.Read(input.ID)
	if err != nil {
		return entity.ReadRecordErr(err)
	}

	err = i.urlRepo.Delete(url.ID)
	if err != nil {
		return entity.DeleteRecordErr(err)
	}

	return nil
}
