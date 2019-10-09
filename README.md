# Funlen linter

Funlen is a linter that checks for long functions. It can checks both on the number of lines and the number of statements.

## Run
Funlen by default, requires you specify your lines or statement limits in
order to run the scanning. Otherwise, it will skip the runs.

## Default Limits
Limits are 60 lines and 40 statements respectively when only either one is
provided.

If both limits are not provided (0,0), Funlen will skip the runs.

## Installation guide

Funlen is included in [https://github.com/golangci/golangci-lint/](golangci-lint). Install it and enable funlen.
