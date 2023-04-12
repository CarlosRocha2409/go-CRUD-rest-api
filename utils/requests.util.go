package utils

import "go.mongodb.org/mongo-driver/mongo/options"

func GetPaginationOptions(limit *int64, page *int64) *options.FindOptions {
	skip := (*page)*(*limit) - (*limit)
	options := options.FindOptions{Limit: limit, Skip: &skip}
	return &options
}
