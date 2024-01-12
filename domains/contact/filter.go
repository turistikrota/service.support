package contact

import "go.mongodb.org/mongo-driver/bson"

type FilterEntity struct {
	Query string `query:"q,omitempty" validate:"omitempty,max=100"`
}

func (r *repo) filterToBson(e FilterEntity) bson.M {
	list := make([]bson.M, 0)
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
					fields.Message: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					fields.Email: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}
