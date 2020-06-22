package bus

import (
	_ "github.com/gleke/base"
	"github.com/gleke/bus/controllers"
	"github.com/gleke/hexya/src/server"
	"github.com/gleke/hexya/src/tools/logging"
	_ "github.com/gleke/web"
)

const MODULE_NAME string = "bus"

var log logging.Logger

func init() {
	log = logging.GetLogger("bus")
	server.RegisterModule(&server.Module{
		Name:    MODULE_NAME,
		PreInit: func() {},
		PostInit: func() {
			controllers.Dispatcher.Start()
		},
	})
}
