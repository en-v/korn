package holder

import (
	"github.com/en-v/log"
)

func (self *_Holder) catchError(err error) {
	log.Error(err)
	if err != nil {
		go func() {
			self.errch <- err
		}()
	}
}
