package handler

import "context"

type ContextKey string

var (
	keyCategoryId    ContextKey = "category_id"
	keySubCategoryId ContextKey = "sub_category_id"
)

func SetCategoryId(ctx context.Context, value int) context.Context {
	return context.WithValue(ctx, keyCategoryId, value)
}

func GetCategoryId(ctx context.Context) (value int, ok bool) {
	value, ok = ctx.Value(keyCategoryId).(int)
	return
}

func SetSubCategoryId(ctx context.Context, value int) context.Context {
	return context.WithValue(ctx, keySubCategoryId, value)
}

func GetSubCategoryId(ctx context.Context) (value int, ok bool) {
	value, ok = ctx.Value(keySubCategoryId).(int)
	return
}
