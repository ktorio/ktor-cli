package i18n

var en = map[Message]string{
	CannotDetermineHomeDir:             "Home directory cannot be determined.",
	CannotDetermineProjectDirOfProject: "Cannot determine the project directory for project %s.\n",
	CannotDetermineProjectDir:          "Cannot determine project directory %s\n",
	ErrorInitLogFile:                   "Failed to initialize log file: %s.\n",
	VersionInfo:                        "Ktor CLI version %s.\n",
	LogHint:                            "For more details, see the log: %s.\n",
	ProjectCreated:                     "Project \"%s\" created at %s.\n",
	JDKDetectedJavaHome:                "JDK detected in JAVA_HOME=%s.\n",
	JdkDetected:                        "Detected JDK %s.\n",
	JdkFoundLocally:                    "Local JDK found at %s.\n",
	JdkDownloaded:                      "JDK downloaded to %s.\n",
	JdkVerificationFailed:              "JDK verification failed. Retrying...",
	GenServerError:                     "Error connecting to the generation server. Please try again later.\n",
	GenServerTimeoutError:              "Connection to the generation server timed out. Please try again later.\n",
	NetworkError:                       "Network error. Check your Internet connection and try again.\n",
	InternalError:                      "Internal error. To get help, report the issue on YouTrack: https://youtrack.jetbrains.com/newIssue?project=ktor.\n",
	ProjectDirExistAndNotEmpty:         "The project directory %s already exists and is not empty.\n",
	NoPermsCreateProjectDir:            "Insufficient permissions to create project directory %s.\n",
	ProjectExtractError:                "Failed to extract downloaded archive to the directory %s.\n",
	JdkExtractError:                    "Failed to extract downloaded JDK to the directory %s.\n",
	DirAlreadyExist:                    "The directory already exists.",
	UnableExtractJdk:                   "Failed to extract downloaded JDK.\n",
	UnableDownloadJdk:                  "Failed to download JDK %s for %s %s\n",
	JdkServerError:                     "Error connecting to the JDK server. Please try again later.\n",
	JdkServerDownloadError:             "Error downloading from the JDK server. Please try again later.\n",
	ChecksumVerificationFailed:         "Checksum verification failed for the downloaded JDK.",
	UnableMakeFileExec:                 "Failed to make %s executable.\n",
	UnexpectedError:                    "An unexpected error occurred.\n",
	UnexpectedErrorWithArg:             "Unexpected error %s.\n",
	UnableCreateStoreJdkDir:            "Failed to create root directory %s for storing downloaded JDKs.\n",
	UnrecognizedFlagsError:             "Unrecognized flags: %s.\n",
	NoCommandError:                     "Expected a command",
	CommandNotFoundError:               "Command '%s' not recognized.\n",
	CommandArgumentsError:              "Command %s requires %d argument[s]: %s.\n",
	ToRunProject:                       "To run the project, use the following commands:\n\n",
	JavaHomeJdkIdeaInstruction:         "Either permanently set the JAVA_HOME environment variable, or add the following JDK in IntelliJ IDEA: \n",
	ToolSummary:                        "Ktor CLI is a tool for generating Ktor projects.\n\n",
	OptionsCaption:                     "The options are:",
	CommandsCaption:                    "The commands are:",
	VerifyingJdk:                       "Verifying %s...\n",
	CreatingDir:                        "Creating directory %s...\n",
	Extracting:                         "Extracting %s to %s...\n",
	RequestGenServer:                   "Requesting generation server...",
	ExtractingArchiveToDir:             "Extracting downloaded archive to directory %s\n...",
	ExtractProjectArchive:              "Extracting project archive... ",
	MakeFileExec:                       "Making %s file executable...\n",
	UsageLine:                          "Usage: ktor [options] <command> [arguments]\n\n",
	TermHeightSmall:                    "Terminal height %d is too small to display plugins.",
	SelectedPluginsCount:               "%d plugins added.",
	ProjectNameCaption:                 "Project name:",
	LocationCaption:                    "Location:",
	SearchPluginsCaption:               "Search for plugins:",
	CreateProjectButton:                "CREATE PROJECT (CTRL+G)",
	NoPluginsFound:                     "No plugins found for the search query.",
	DirNotEmptyError:                   "Directory %s is not empty",
	DirNotExist:                        "Directory %s does not exist",
	ProjectDirLong:                     "Project directory name is too long.",
	DownloadingJdk:                     "Downloading %s from %s...\n",
	DownloadingJdkProgress:             "Downloading JDK... ",
	ExtractingJdkFiles:                 "Extracting JDK files to %s\n",
	ExtractingJdkProgress:              "Extracting JDK... ",
	ByeMessage:                         "Goodbye!",
	UnableFetchPluginsError:            "Failed to fetch plugins from the generation server. Restart the app.",
	FetchingJdk:                        "Fetching %s...\n",
	DownloadingProjectArchiveProgress:  "Downloading project archive... ",
	ProjectNameRequired:                "Project name is empty",
	ProjectNameAllowedChars:            "Only Latin characters, digits, '_', '-' and '.' are allowed for the project name",
	DownloadOpenAPIJarError:            "Error downloading OpenAPI utility. Please try again later.\n",
	OpenApiExecuteJarError:             "Failed to execute OpenAPI utility.\n",
	ExternalCommandError:               "Error executing an external command.\n",
	OpenApiSpecNotExist:                "OpenAPI spec file %s does not exist.\n",
	CreateOpenApiJar:                   "Creating OpenAPI JAR file %s\n",
	ExecutingCommand:                   "Executing command: %s\n",
	FlagRequiresArgument:               "Flag %s requires one argument\n",
	DownloadingOpenApiJarProgress:      "Downloading OpenAPI utility... ",
	OpenApiCommandDescr:                "Generate a Ktor project with a given OpenAPI specification.",
	NewCommandDescr:                    "Generate a new Ktor project. If a project name is omitted, the interactive mode will be invoked.",
	VersionCommandDescr:                "Display the ktor CLI tool version.",
	HelpCommandDescr:                   "Show the ktor CLI tool usage text.",
	VerboseOptionDescr:                 "Enable verbose mode.",
	OutputDirOptionDescr:               "Project output directory.",
}
