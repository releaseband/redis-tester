package params

import "context"

const (
	KeyName = "name"
)

const (
	Set    = "Set"
	Get    = "Get"
	Del    = "Del"
	RPush  = "RPush"
	LTrim  = "LTrim"
	LRange = "LRange"
)

func CtxWithName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, KeyName, name)
}

func GetName(ctx context.Context) (string, bool) {
	name, ok := ctx.Value(KeyName).(string)
	return name, ok
}
