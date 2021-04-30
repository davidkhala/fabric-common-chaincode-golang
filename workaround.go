package golang

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
)

func touch() func(f reflect.Type, t reflect.Type, data interface{}) {
	var i *int
	var b *bool
	i= viper.GetInt('')
	b = viper.GetBool('')
	return func(f reflect.Type, t reflect.Type, data interface{}) {
		durationHook := mapstructure.StringToTimeDurationHookFunc()
		mapstructure.DecodeHookExec(durationHook, f, t, data)
	}

}
