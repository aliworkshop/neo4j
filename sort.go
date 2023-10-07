package neo4j

import (
	"fmt"
	"github.com/aliworkshop/error"
)

func (n *neo) sort(query Query, prefix ...string) (string, error.ErrorModel) {
	var pre string
	if len(prefix) > 0 {
		pre = fmt.Sprintf("%s.", prefix[0])
	}
	sortItems := query.GetSort()
	out := ""
	if sortItems != nil {
		for _, sortItem := range sortItems {
			sort := pre + sortItem.Field
			if sortItem.Order != "" {
				if sortItem.Order.IsDescending() {
					sort += " DESC, "
				} else if sortItem.Order.IsAscending() {
					sort += " ASC, "
				} else {
					return "", error.New().
						WithType(error.TypeValidation).
						WithId("InvalidSortQuery").
						WithDetail("invalid sort order. order takes one of values of `DESC` or `ASC`")
				}
				out += sort
			}
		}
		out = out[:len(out)-2]
	}
	return out, nil
}
