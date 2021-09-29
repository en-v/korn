package jsf

import "os"

func (self *JSF) Remove(holder string, id string) error {
	filename := self.fullFileName(holder, id)
	return os.Remove(filename)
}
