package wrappers

import (
	"fmt"

	"github.com/alhamsya/boilerplate-go/lib/managers/custom_log"
)

func (w *Wrapper) GetWrapper(usecase string) *Wrapper {
	if wrapper, ok := w.CW[usecase]; ok {
		w.Wrapper = wrapper
		return w
	}

	w.Err = fmt.Errorf("[GetCallWrapper] unable to get wrapper(%s)", usecase)
	return w
}

func (w *Wrapper) Call(req func() (interface{}, error)) (interface{}, error) {
	if w.Err != nil {
		customLog.Error(w.Err.Error())
		return req()
	}
	return w.Wrapper.Execute(req)
}
