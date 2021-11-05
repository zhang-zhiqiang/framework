package main

import "go_framework/framework"

func registerRouter(core *framework.Core) {
    core.Get("/user/login", UserLoginController)

    // 需求3:批量通用前缀
    subjectApi := core.Group("/subject")
    {
        // 需求4:动态路由
        subjectApi.Delete("/:id", SubjectDelController)
        subjectApi.Put("/:id", SubjectUpdateController)
        subjectApi.Get("/:id", SubjectGetController)
        subjectApi.Get("/list/all", SubjectListController)
    }
}