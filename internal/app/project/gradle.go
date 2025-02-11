package project

func GuessGradleTasks(projectDir string) (runTask, buildTask string, err error) {
	//	If the build gradle has Ktor plugin and it's applied
	//		Use run and classes command
	//	Else
	//		Search one level in the subdirectories for build.gradle.kts
	//		For each found
	//			If the build gradle has Ktor plugin and it's applied
	//				Use :subdir:run and :subdir:classes command

	return
}
