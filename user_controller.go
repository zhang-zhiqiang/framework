package main

import "go_framework/framework"

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
