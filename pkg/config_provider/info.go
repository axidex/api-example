package config_provider

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/pkg/logger"
	"github.com/axidex/api-example/pkg/version"
	"reflect"
)

type Infoer interface {
	Info() string
}

func PrintInfo(s interface{}, printFunc func(ctx context.Context, s string, attrs ...logger.Attribute)) {
	printVersion(printFunc)
	printInfoers(s, printFunc)
}

func printInfoers(s interface{}, printFunc func(ctx context.Context, s string, attrs ...logger.Attribute)) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		// Пропускаем неэкспортированные поля
		if !fieldType.IsExported() {
			continue
		}

		// Проверяем, реализует ли поле интерфейс Infoer
		if info, ok := field.Interface().(Infoer); ok {
			printFunc(context.Background(), info.Info())
		}

		// Рекурсивно проверяем вложенные структуры
		if field.Kind() == reflect.Struct {
			printInfoers(field.Addr().Interface(), printFunc)
		}
	}
}

func printVersion(printFunc func(ctx context.Context, s string, attrs ...logger.Attribute)) {
	v := version.NewVersion()
	printFunc(context.Background(), fmt.Sprintf("Build info: Version - %s, Date - %s, Commit - %s", v.Version(), v.Date(), v.Commit()))
}
