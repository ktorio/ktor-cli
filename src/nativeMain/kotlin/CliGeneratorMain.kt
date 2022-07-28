import io.ktor.client.*
import io.ktor.client.features.*
import io.ktor.client.features.json.*
import io.ktor.client.features.json.serializer.*
import io.ktor.generator.api.*
import io.ktor.generator.bundle.*
import io.ktor.generator.cli.installer.*
import kotlinx.cli.*

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
        ktorInstaller.downloadKtorProject(projectName)
    }
}

class RunProject(client: HttpClient) : KtorCommand(
    "start", description = PropertiesBundle.message("run.command.description"), client = client
) {
    override fun execute() = executeCatching {
        ktorInstaller.runKtorProject(projectName)
    }
}

@OptIn(ExperimentalCli::class)
class KtorParser(client: HttpClient) : ArgParser(PropertiesBundle.message("program.name")) {
    init {
        subcommands(GenerateProject(client), RunProject(client))
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
    val client = createHttpClient()

    val parser = KtorParser(client)
    if (args.isEmpty()) {
        parser.parse(arrayOf("--help"))
        return
    }
    parser.parse(args)
}
