package jfs

import "os"

func (self *JFS) Remove(holder string, id string) error {
	filename := self.fullFileName(holder, id)
	return os.Remove(filename)
}
