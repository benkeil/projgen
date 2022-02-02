# projgen

```bash
Projgen is a CLI library that generates new projects based on a template.

Usage:
  projgen [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generates a new repository
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.config/projgen/config.yaml)
  -h, --help            help for projgen
  -t, --toggle          Help message for toggle

Use "projgen [command] --help" for more information about a command. 
```

## Motivation

Most project generators are language specific or don't provide an executable. GitHub templates have fixed versions, that
are always outdated and need a lot of maintenance. With `projgen` you can easily execute commands and add files or
folders to initialize your project.

## Download and install

TBD

## Configuring default values

You can configure `projgen` with a config file in `$HOME/.config/projgen/config.yaml`.

### Config properties

Not yet available`

## Commands

### `generate`

```bash
Generates a new repository from a template.

Usage:
  projgen generate TEMPLATE [flags]

Examples:
  to use a local directory or repository
  $ projgen generate ./templates/typescript -n my-project
  to use a remote repository
  $ projgen generate git@github.com:user/template.git -n my-project


Flags:
  -h, --help                  help for generate
  -n, --project-name string   The project name

Global Flags:
      --config string   config file (default is $HOME/.config/projgen/config.yaml)
```

You can use any directory or git repository as `TEMPLATE`.

To create and organize new repositories, `projgen` uses [ghq](https://github.com/x-motemen/ghq).

## Templates

What does `generate` do:

1. Create a new local directory
1. Copy all files from the template into this directory
1. Execute `.projgen.yaml`

### Example

```yaml
steps:
  - command: npm init -y
    title: generate package.json
  - command: pnpm i --save-dev typescript
    title: install typescript
  - command: npx tsc --init
    title: initialize typescript
  - command: pnpm i --save-dev eslint prettier
    title: add additional dependencies
  - render: README.md

```

Contains steps that need to be done to initialize your project.

### Steps

`Command`

| property  | description                                                          |
|-----------|----------------------------------------------------------------------|
| `command` | The command that should be executed                                  |
| `title`   | An alternative title that should be displayed instead of the command |

`Render`

| property | description                                                    |
|----------|----------------------------------------------------------------|
| `render` | A file that should be rendered with the go templating language |
