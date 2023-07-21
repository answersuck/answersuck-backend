package appctx

import "context"

var FootPrintKey = struct{}{}

type FootPrint struct {
	RemoteAddr string
	UserAgent  string
}

func GetFootPrint(ctx context.Context) FootPrint {
	fp, ok := ctx.Value(FootPrintKey).(FootPrint)
	if !ok {
		return FootPrint{}
	}

	return fp
}
