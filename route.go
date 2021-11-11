package main

import (
	"go_framework/framework"
	"go_framework/framework/middleware"
)

func registerRouter(core *framework.Core) {
	core.Use(middleware.Test1(), middleware.Test2())
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4:动态路由
		subjectApi1 := subjectApi.Group("/subject")
		subjectApi1.Use(middleware.Test3(), middleware.Test2(), middleware.Test1())
		{
			subjectApi1.Delete("/:id", SubjectDelController)
			subjectApi1.Put("/:id", SubjectUpdateController)
			subjectApi1.Get("/:id", SubjectGetController)
			subjectApi1.Get("/list/all", SubjectListController)
		}

	}
}
