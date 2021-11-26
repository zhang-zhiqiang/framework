package env

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/provider/app"
	tests "github.com/gohade/hade/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHadeEnvProvider(t *testing.T) {
	Convey("test hade env norm case", t, func() {
		basePath := tests.BasePath
		c := framework.NewHadeContainer()
		sp := &app.HadeAppProvider{BaseFolder: basePath}

		err := c.Bind(sp)
		So(err, ShouldBeNil)

		sp2 := &HadeEnvProvider{}
		err = c.Bind(sp2)
		So(err, ShouldBeNil)

		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		So(envServ.AppEnv(), ShouldEqual, "development")
	})
}
