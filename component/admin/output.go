package admin

import "shorturl/entity"

type ListOutput struct {
	Total int
	Items []entity.URL
}
