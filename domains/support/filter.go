package support

import "go.mongodb.org/mongo-driver/bson"

type FilterEntity struct {
	Query string `query:"q,omitempty" validate:"omitempty,max=100"`
	State string `query:"state,omitempty" validate:"omitempty,oneof=open answered closed deleted"`
}

func (r *repo) filterToBson(e FilterEntity, defaultFilters ...bson.M) bson.M {
	list := make([]bson.M, 0)
	if len(defaultFilters) > 0 {
		list = append(list, defaultFilters...)
	}
	list = r.filterByState(list, e)
	list = r.filterByQuery(list, e)
	listLen := len(list)
	if listLen == 0 {
		return bson.M{}
	}
	if listLen == 1 {
		return list[0]
	}
	return bson.M{
		"$and": list,
	}
}

func (r *repo) filterByState(list []bson.M, filter FilterEntity) []bson.M {
	if filter.State != "" {
		list = append(list, bson.M{
			fields.State: filter.State,
		})
	}
	return list
}

func (r *repo) filterByQuery(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Query != "" {
		list = append(list, bson.M{
			"$or": []bson.M{
				{
					fields.Subject: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					messageField(messageFields.Text): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					userField(userFields.Name): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}
