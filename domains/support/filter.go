package support

import "go.mongodb.org/mongo-driver/bson"

type FilterEntity struct{}

func (r *repo) filterToBson(e FilterEntity) bson.M {
	return bson.M{}
}
