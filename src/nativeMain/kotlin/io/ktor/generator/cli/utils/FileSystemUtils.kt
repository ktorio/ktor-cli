package io.ktor.generator.cli.utils

import io.ktor.utils.io.core.*
import kotlinx.cinterop.*
import okio.BufferedSink
import okio.BufferedSource
import okio.FileSystem
import okio.Path.Companion.toPath
import okio.buffer
import platform.posix.FILE
import platform.posix.fgets
import platform.posix.getcwd


private const val READ_MODE = "r"
private const val WRITE_MODE = "w"

expect val FS_DELIMETER: String
expect fun unzip(zipFile: File, outputDir: Directory)

expect fun addExecutablePermissions(file: File)

expect fun homePath(): String

interface FsUnit {
    val path: String
    val name: String get() = path.split(FS_DELIMETER).last()

    fun parentDir() = Directory(path.replaceAfterLast(FS_DELIMETER, "").drop(1))
    fun exists(): Boolean = FileSystem.SYSTEM.exists(path.toPath())
}

data class Directory(override val path: String) : FsUnit {
    private fun calcEntryPath(name: String): String = "$path$FS_DELIMETER$name"

    fun content(): List<FsUnit> {
        return FileSystem.SYSTEM.list(path.toPath()).map { entryPath ->
            val isFile: Boolean = FileSystem.SYSTEM.listOrNull(entryPath) == null //{

            if (isFile) File(entryPath.toString())
            else Directory(entryPath.toString())
        }
    }

    fun createFileIfNeeded(name: String): File {
        val filePath = calcEntryPath(name)
        if (!FileSystem.SYSTEM.exists(filePath.toPath())) {
            FileSystem.SYSTEM.sink(filePath.toPath())
        }

        return File(filePath)
    }

    fun subdir(name: String): Directory {
        val dirPath = calcEntryPath(name)
        return Directory(dirPath)
    }

    fun createDirIfNeeded(name: String): Directory {
        val dir = subdir(name)
        val dirPath = dir.path.toPath()
        if (!dir.exists()) {
            try {
                FileSystem.SYSTEM.createDirectory(dirPath, true)
            } catch (cause: Throwable) {
                throw Exception("Failed to find/create directory $dirPath")
            }
        }
        return dir
    }

    companion object {
        fun home(): Directory = Directory(homePath())

        fun current(): Directory = memScoped {
            val pathBufferSize = 1024
            val pathBuffer = allocArray<ByteVar>(pathBufferSize)
            getcwd(pathBuffer, pathBufferSize.toULong()) ?: throw Exception("Failed to locate working dir")

            Directory(pathBuffer.toKString())
        }
    }
}

class FileDataInputStream private constructor(private val source: BufferedSource) {
    companion object {
        fun from(file: File): FileDataInputStream {
            return FileDataInputStream(FileSystem.SYSTEM.source(file.path.toPath()).buffer())
        }
    }

    fun readInt(): Int = source.readInt()
    fun readByte(): Byte = source.readByte()
    fun readLine(): String? = source.readUtf8Line()
    fun readByteArray(): ByteArray = source.readByteArray()

    fun close() {
        source.close()
    }
}

class FileDataOutputStream private constructor(private val sink: BufferedSink) {
    companion object {
        fun from(file: File): FileDataOutputStream {
            return FileDataOutputStream(FileSystem.SYSTEM.sink(file.path.toPath()).buffer())
        }
    }

    fun writeInt(value: Int) {
        sink.writeInt(value)
    }

    fun writeByte(value: Byte) {
        sink.writeByte(value.toInt())
    }

    fun write(data: ByteArray) {
        sink.write(data)
    }

    fun close() {
        sink.close()
    }
}

data class File(override val path: String) : FsUnit {
    fun readContent(): ByteArray = FileSystem.SYSTEM.source(path.toPath()).buffer().readByteArray()

    fun readText(): String = FileSystem.SYSTEM.source(path.toPath()).buffer().readUtf8Line() ?: ""

    fun writeText(text: String) {
        FileSystem.SYSTEM.write(path.toPath()) { this.write(text.toByteArray()) }
    }

    fun writeContent(content: ByteArray) {
        FileSystem.SYSTEM.write(path.toPath()) {
            this.write(content)
        }
    }

    fun delete() {
        FileSystem.SYSTEM.delete(path.toPath(), false)
    }
}

fun File.readLines(): List<String> = readText().split("\n").filter { it.isNotEmpty() }

fun MemScope.handleOutput(file: CPointer<FILE>, handle: (String) -> Unit) {
    val readBufferLength = 64 * 1024
    val buffer = allocArray<ByteVar>(readBufferLength)
    var line = fgets(buffer, readBufferLength, file)?.toKString()
    while (line != null) {
        handle(line)
        line = fgets(buffer, readBufferLength, file)?.toKString()
    }
}