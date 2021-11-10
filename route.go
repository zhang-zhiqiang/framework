package main

import (
	"go_framework/framework"
	"time"
)

func registerRouter(core *framework.Core) {
	core.Get("/user/login", framework.TimeoutHandler(UserLoginController, time.Second))

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4:动态路由
		subjectApi1 := subjectApi.Group("/subject")
		{
			subjectApi1.Delete("/:id", SubjectDelController)
			subjectApi1.Put("/:id", SubjectUpdateController)
			subjectApi1.Get("/:id", SubjectGetController)
			subjectApi1.Get("/list/all", SubjectListController)
		}

	}
}
