package util

import (
	"context"
	"fmt"
	"github.com/mintance/go-uniqid"
	"math/rand"
	"strings"
	"time"
)

func BusinessID(prefix string) string {
	date := time.Now().Format("20060102")
	id := uniqid.New(uniqid.Params{date + "-" + prefix, false}) + RandomInt(1000000, 1)

	return strings.ToUpper(id)
}

func RandomInt(max, min int64) string {
	return fmt.Sprint(rand.Int63n(max-min) + min)
}

func ContextWithTimeOut() (context_data context.Context) {
	context_data, _ = context.WithTimeout(context.Background(), time.Minute*10)
	return context_data
}
