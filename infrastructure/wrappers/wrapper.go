package wrappers

import (
	"fmt"
	"log"
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
		log.Println(w.Err)
		return req()
	}
	return w.Wrapper.Execute(req)
}
