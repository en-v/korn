package container

import (
	"github.com/en-v/log"
)

func (this *Container) catchError(err error) {
	log.Error(err)
	if err != nil {
		go func() {
			this.errs <- err
		}()
	}
}
