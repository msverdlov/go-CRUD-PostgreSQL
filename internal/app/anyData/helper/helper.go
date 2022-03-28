package helper

import (
	"anyData/internal/app/anyData/localStorage"
	"errors"
	"fmt"
)

type Counter struct {
	counter uint64
}

func (cnt *Counter) Increment() uint64 {
	cnt.counter++
	return cnt.counter
}

func (cnt *Counter) Decrease() uint64 {
	cnt.counter--
	return cnt.counter
}

func (cnt *Counter) Current() uint64 {
	return cnt.counter
}

func IndexOf(id uint64) (uint64, error) {
	for index, data := range localStorage.DataStorage {
		if data.Id == id {
			return uint64(index), nil
		}
	}
	errorMessage := fmt.Sprintf("ID %v: data not found.", id)

	return 0, errors.New(errorMessage)
}

