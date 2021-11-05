package main

import (
    "go_framework/framework"
)


func UserLoginController(c *framework.Context) error {
    // 打印控制器名字
    c.Json(200, "ok, UserLoginController")
    return nil
}


func SubjectDelController(c *framework.Context) error {
    // 打印控制器名字
    c.Json(200, "del, SubjectDelController")
    return nil
}

func SubjectUpdateController(c *framework.Context) error {
    // 打印控制器名字
    c.Json(200, "update, SubjectUpdateController")
    return nil
}

func SubjectGetController(c *framework.Context) error {
    // 打印控制器名字
    c.Json(200, "get, SubjectGetController")
    return nil
}

func SubjectListController(c *framework.Context) error {
    // 打印控制器名字
    c.Json(200, "list, SubjectListController")
    return nil
}