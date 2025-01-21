package i18n

type Message int

const (
	CannotDetermineHomeDir Message = iota
	CannotDetermineProjectDirOfProject
	CannotDetermineProjectDir
	ErrorInitLogFile
	VersionInfo
	LogHint
	ProjectCreated
	JDKDetectedJavaHome
	JdkDetected
	JdkFoundLocally
	JdkDownloaded
	JdkVerificationFailed
	GenServerError
	GenServerTimeoutError
	NetworkError
	InternalError
	ProjectDirExistAndNotEmpty
	NoPermsCreateProjectDir
	ProjectExtractError
	JdkExtractError
	DirAlreadyExist
	UnableExtractJdk
	UnableDownloadJdk
	JdkServerError
	JdkServerDownloadError
	ChecksumVerificationFailed
	UnableMakeFileExec
	UnexpectedError
	UnexpectedErrorWithArg
	UnableCreateStoreJdkDir
	UnrecognizedFlagsError
	NoCommandError
	CommandNotFoundError
	CommandArgumentsError
	ToRunProject
	JavaHomeJdkIdeaInstruction
	ToolSummary
	OptionsCaption
	CommandsCaption
	VerifyingJdk
	CreatingDir
	Extracting
	RequestGenServer
	ExtractingArchiveToDir
	ExtractProjectArchive
	MakeFileExec
	UsageLine
	TermHeightSmall
	SelectedPluginsCount
	ProjectNameCaption
	LocationCaption
	SearchPluginsCaption
	CreateProjectButton
	NoPluginsFound
	DirNotEmptyError
	DirNotExist
	ProjectDirLong
	DownloadingJdk
	DownloadingJdkProgress
	ExtractingJdkFiles
	ExtractingJdkProgress
	ByeMessage
	UnableFetchPluginsError
	FetchingJdk
	DownloadingProjectArchiveProgress
	ProjectNameRequired
	ProjectNameAllowedChars
	DownloadOpenAPIJarError
	OpenApiExecuteJarError
	ExternalCommandError
	OpenApiSpecNotExist
	CreateOpenApiJar
	ExecutingCommand
	FlagRequiresArgument
	DownloadingOpenApiJarProgress
	NewCommandDescr
	VersionCommandDescr
	HelpCommandDescr
	OpenApiCommandDescr
	AddCommandDescr
	CompletionCommandDescr
	VerboseOptionDescr
	OutputDirOptionDescr
	ProjectCreatedIn
)
