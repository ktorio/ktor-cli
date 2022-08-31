import io.ktor.client.*
import io.ktor.client.features.*
import io.ktor.client.features.json.*
import io.ktor.client.features.json.serializer.*
import io.ktor.generator.api.*
import io.ktor.generator.bundle.*
import io.ktor.generator.cli.installer.*
import io.ktor.generator.cli.utils.*
import kotlinx.cli.*
import platform.posix.exit

private val ALLOWED_PROJECT_NAME = "[a-zA-Z]\\w+".toRegex()

private fun executeCatching(action: () -> Unit) {
    try {
        action()
    } catch (e: Exception) {
        PropertiesBundle.writeMessage("error.happened", e.message ?: e.stackTraceToString())
    }
}

@OptIn(ExperimentalCli::class)
abstract class KtorCommand(
    name: String, description: String, client: HttpClient
) : Subcommand(name, description) {
    protected val projectName: String by argument(
        ArgType.String, description = PropertiesBundle.message("project.name.description")
    )

    protected val ktorInstaller: KtorInstaller by lazy { KtorInstaller(KtorGeneratorWebImpl(client)) }
}

class GenerateProject(client: HttpClient) : KtorCommand(
    "generate", description = PropertiesBundle.message("generate.command.description"), client = client
) {
    override fun execute() = executeCatching {
        if (!projectName.matches(ALLOWED_PROJECT_NAME)) {
            PropertiesBundle.writeErrorMessage("project.name.invalid")
        } else {
            ktorInstaller.downloadKtorProject(projectName)
        }
    }
}

class RunProject(client: HttpClient) : KtorCommand(
    "start", description = PropertiesBundle.message("run.command.description"), client = client
) {
    override fun execute() = executeCatching {
        ktorInstaller.runKtorProject(projectName)
    }
}

class Version(client: HttpClient) : KtorCommand(
    "version",
    description = PropertiesBundle.message("version.command.description"),
    client = client
) {
    override fun execute() = executeCatching {
        PropertiesBundle.writeMessage("version.command", APP_VERSION)
    }
}

@OptIn(ExperimentalCli::class)
class KtorParser(client: HttpClient) : ArgParser(PropertiesBundle.message("program.name")) {
    init {
        subcommands(GenerateProject(client), RunProject(client), Version(client))
    }
}

fun createHttpClient(): HttpClient = HttpClient {
    install(JsonFeature) {
        serializer = KotlinxSerializer()
    }
    install(BodyProgress)
}

@OptIn(ExperimentalCli::class)
fun main(args: Array<String>) {
    assertLibcurlExists()
    val client = createHttpClient()

    val parser = KtorParser(client)
    if (args.isEmpty()) {
        parser.parse(arrayOf("--help"))
        exit(1)
        return
    }
    parser.parse(args)
}
