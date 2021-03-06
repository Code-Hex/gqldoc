# Document Generator for GraphQL

> gqldoc is now alpha

gqldoc is command line tool to generate documents from GraphQL schema or your GraphQL endpoint. the command is written by Go 1.16. So You need Go 1.16 and above If you want to build.

## Demo

- [Star Wars API](example/starwars)
- [Github V4 API](example/github)
  - Public schema is [here](https://docs.github.com/en/graphql/overview/public-schema)

## How to use

### Generate docs from graphql endpoint

You need to enable GraphQL [Introspection](https://graphql.org/learn/introspection/). `--header` and `--query` options are supported.

```sh
$ gqldoc -e http://127.0.0.1:8081/query --header 'Authorization: Bearer token' -o ./doc_dir
```

### Generate docs from graphql schema files

Supported `.json` and `.graphql`, `.gql` extensions. `.json` must be contained the result of introspection.

```sh
$ gqldoc -s schema.graphql -o ./doc_dir
```

If you want to specify multiple schema, you can use `--schema` flag repeatedly.

```sh
$ gqldoc -s a.graphql -s b.graphql -o ./doc_dir
```

You can also use glob.

```sh
$ gqldoc -s "schema/**/*.graphql" -o ./doc_dir
```

## Installation

### Mac and Linux users via Homebrew

```sh
$ brew install Code-Hex/tap/gqldoc
```

### Manually via go command

Again, this tool is supported Go 1.16 and above. If you are one of those users, you can install like below.

```sh
$ go install github.com/Code-Hex/gqldoc/cmd/gqldoc@latest
```

### Manually download

You can download binary from [here](https://github.com/Code-Hex/gqldoc/releases)

## Todo

This todo list is written the order of priority.

- [x] ~~GitHub Actions for this~~ **Available**
  - Update document and push it to current branch If target schema has diff
  - https://github.com/Code-Hex/gqldoc-actions
- [ ] Subscription
- [ ] Directive and it's location.
- [x] ~~Put README.md in output directory~~
  - TOC purpose
- [ ] Custom Template
- [ ] Render HTML
  - I'm wondering if I should really support this feature.

