package config_provider

import "reflect"

type Infoer interface {
	Info() string
}

func PrintInfo(s interface{}, printFunc func(s string, args ...interface{})) {
	printVersion(printFunc)
	printInfoers(s, printFunc)
}

func printInfoers(s interface{}, printFunc func(s string, args ...interface{})) {
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
			printFunc(info.Info())
		}

		// Рекурсивно проверяем вложенные структуры
		if field.Kind() == reflect.Struct {
			printInfoers(field.Addr().Interface(), printFunc)
		}
	}
}

func printVersion(printFunc func(s string, args ...interface{})) {
	version := NewVersion()
	printFunc("Build info: Version - %s, Date - %s, Commit - %s", version.Version(), version.Date(), version.Commit())
}
