package stream

import (
	"datalistener/src/models"
	"fmt"
)

type DebugStream struct {
}

func (s DebugStream) Notify(e models.EntryData) error {

	fmt.Println(e)
	return nil
}
