import io.ktor.client.*
import io.ktor.client.features.*
import io.ktor.client.features.json.*
import io.ktor.client.features.json.serializer.*
import io.ktor.generator.bundle.*
import io.ktor.generator.cli.installer.*
import kotlinx.cli.*

private const val DEFAULT_KTOR_URL = "https://ktor-plugin.europe-north1-gke.intellij.net"

@OptIn(ExperimentalCli::class)
abstract class KtorCommand(
    name: String,
    description: String,
    client: HttpClient
) : Subcommand(name, description) {
    private val host: String by option(
        ArgType.String,
        fullName = "host",
        description = PropertiesBundle.message("ktor.backend.url.description")
    ).default(DEFAULT_KTOR_URL)

    protected val projectName: String by argument(
        ArgType.String, description = PropertiesBundle.message("project.name.description")
    )

    protected val ktorInstaller: KtorInstaller by lazy { KtorInstaller(client, host) }
}

class GenerateProject(client: HttpClient) : KtorCommand(
    "generate",
    description = PropertiesBundle.message("generate.command.description"),
    client = client
) {
    override fun execute() {
        ktorInstaller.downloadKtorProject(projectName)
    }
}

class RunProject(client: HttpClient) : KtorCommand(
    "run",
    description = PropertiesBundle.message("run.command.description"),
    client = client
) {
    private val args: List<String> by argument(
        ArgType.String,
        fullName = "args",
        description = PropertiesBundle.message("run.arguments.description")
    ).vararg()

    override fun execute() {
        ktorInstaller.runKtorProject(projectName, args)
    }
}

class KtorParser(client: HttpClient) : ArgParser(PropertiesBundle.message("program.name")) {
    init {
        subcommands(GenerateProject(client), RunProject(client))
    }
}

@OptIn(ExperimentalCli::class)
fun main(args: Array<String>) {
    val client = HttpClient {
        install(JsonFeature) {
            serializer = KotlinxSerializer()
        }
        install(BodyProgress)
    }

    val parser = KtorParser(client)

    parser.parse(args)
}
