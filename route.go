package main

import "go_framework/framework"

func registerRouter(core *framework.Core) {
    // core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
    core.Get("foo", FooControllerHandler)
}
