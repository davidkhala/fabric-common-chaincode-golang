package golang

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
)

func touch() func(f reflect.Type, t reflect.Type, data interface{}) {
	viper.GetViper()
	return func(f reflect.Type, t reflect.Type, data interface{}) {
		durationHook := mapstructure.StringToTimeDurationHookFunc()
		mapstructure.DecodeHookExec(durationHook, f, t, data)
	}

}
