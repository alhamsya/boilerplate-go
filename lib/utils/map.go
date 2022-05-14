package utils

import (
	"github.com/alhamsya/boilerplate-go/lib/managers/custom_error"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func (l *thing) ToStruct(v map[string]interface{}, r interface{}) (err error) {
	ms, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &r,
	})
	if err != nil {
		return customError.Wrap(err, "init decoder")
	}
	err = ms.Decode(v)
	if err != nil {
		return customError.Wrap(err, "decode map to struct")
	}
	return
}

func (l *thing) ToMap(v interface{}) (r map[string]interface{}) {
	s := structs.New(v)
	s.TagName = "json"
	r = s.Map()
	return
}
