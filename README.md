# Funlen linter

Funlen is a linter that checks for long functions. It can check both on the number of lines and the number of statements.

The default limits are 60 lines and 40 statements. You can configure these.

## Description

The intent for the funlen linter is to fit a function within one screen. If you need to scroll through a long function, tracing variables back to their definition or even just finding matching brackets can become difficult.

Besides checking lines there's also a separate check for the number of statements, which gives a clearer idea of how much is actually being done in a function.

The default values are used internally, but might to be adjusted for your specific environment.

## Installation

Funlen is included in [golangci-lint](https://github.com/golangci/golangci-lint/). Install it and enable funlen.
