## Directories

The `grammars` directory contains ANTLR4 grammars:

* Toml from https://github.com/antlr/grammars-v4/tree/master/toml
* Kotlin from https://github.com/Kotlin/kotlin-spec/tree/release/grammar/src/main/antlr

The `parsers` directory contains generated parsers for Go language.

## Generate a parser

1. Download ANTLR tool https://www.antlr.org/download/antlr-4.13.2-complete.jar
2. Execute the following commands to generate a lexer and a parser for the specific language:
```shell
cd internal/app/lang
java -jar antlr-4.13.2-complete.jar -Dlanguage=Go -o parsers/{lang} grammars/{lang}/{Lang}Lexer.g4
java -jar antlr-4.13.2-complete.jar -Dlanguage=Go -o parsers/{lang} grammars/{lang}/{Lang}Parser.g4
```
