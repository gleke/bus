package bus

import (
	"github.com/gleke/base"
	"github.com/gleke/pool/h"
)

func init() {
	h.BusPresence().Methods().AllowAllToGroup(base.GroupUser)
	h.BusPresence().Methods().AllowAllToGroup(base.GroupPortal)
}
