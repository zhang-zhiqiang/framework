package command

import "github.com/gohade/hade/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	//root.AddCommand(DemoCommand)

	// provider
	root.AddCommand(initProviderCommand())
	// env
	root.AddCommand(envCommand)

	// cron
	root.AddCommand(initCronCommand())

	// cmd
	root.AddCommand(initCmdCommand())

	//// build
	//buildCommand.AddCommand(buildSelfCommand)
	//buildCommand.AddCommand(buildBackendCommand)
	//buildCommand.AddCommand(buildFrontendCommand)
	//buildCommand.AddCommand(buildAllCommand)
	//root.AddCommand(buildCommand)
	//
	//// app
	root.AddCommand(initAppCommand())
	//
	//// dev
	root.AddCommand(initDevCommand())
	//
	// middleware
	root.AddCommand(initMiddlewareCommand())
	//
	//// swagger
	//swagger.IndexCommand.AddCommand(swagger.InitServeCommand())
	//swagger.IndexCommand.AddCommand(swagger.GenCommand)
	//root.AddCommand(swagger.IndexCommand)
	//
	//// provider
	//providerCommand.AddCommand(providerListCommand)
	//providerCommand.AddCommand(providerCreateCommand)
	//root.AddCommand(providerCommand)
	//
	//// new
	root.AddCommand(initNewCommand())
	// build
	root.AddCommand(initBuildCommand())
}
