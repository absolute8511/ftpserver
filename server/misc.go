package server

import (
	"fmt"
)

func (self *Paradise) handleSyst() {
	self.writeMessage(215, "UNIX Type: L8")
}

func (self *Paradise) handlePwd() {
	self.writeMessage(257, "\""+self.path+"\" is the current directory")
}
