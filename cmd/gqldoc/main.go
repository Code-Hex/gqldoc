package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/Code-Hex/gqldoc/internal/markdown"
	"github.com/Code-Hex/gqldoc/loader"
	"github.com/machinebox/graphql"
	gqlcli "github.com/machinebox/graphql"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// Version holds the complete version number. Filled in at linking time.
var Version = "v0.0.0+unknown"

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "err: %+v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	app := cli.NewApp()
	app.Name = "gqldoc"
	app.Description = "This is a command for quickly generating documents from graphql schema in golang."
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "endpoint",
			Aliases: []string{"e"},
			Usage:   `GraphQL http endpoint. e.g. "https://example.com/graphql"`,
		},
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"x"},
			Usage:   `HTTP header string for request. (use with --endpoint) e.g. "Authorization: Token cb8795e7"`,
		},
		&cli.StringSliceFlag{
			Name:    "query",
			Aliases: []string{"q"},
			Usage:   `HTTP query string for request. (use with --endpoint) e.g. "token=cb8795e7"`,
		},
		&cli.StringSliceFlag{
			Name:    "schema",
			Aliases: []string{"s"},
			Usage:   `GraphQL schema file. "1 json graphql schema" or "1 or more graphql files (.graphqls, .graphql, .gql)".`,
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   `Output directory.`,
		},
	}
	app.Version = Version
	app.Action = func(ctx *cli.Context) error {
		return generateCommand(ctx)
	}

	return app.Run(os.Args)
}

func generateCommand(ctx *cli.Context) (err error) {
	config := &markdown.Config{
		BaseDir: ctx.String("output"),
	}

	endpoint := ctx.String("endpoint")
	if endpoint != "" {
		schema, err := requestSchema(endpoint,
			ctx.StringSlice("query"),
			ctx.StringSlice("header"),
		)
		if err != nil {
			return errors.WithStack(err)
		}
		if err := config.Render(schema.Schema); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}

	schemaFiles := ctx.StringSlice("schema")

	schema, err := convertSchemaTree(schemaFiles)
	if err != nil {
		if err == errHelp {
			return cli.ShowAppHelp(ctx)
		}
		return errors.WithStack(err)
	}

	if err := config.Render(schema.Schema); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

var errHelp = errors.New("help")

func convertSchemaTree(schemaFiles []string) (*introspection.Root, error) {
	if len(schemaFiles) == 1 && filepath.Ext(schemaFiles[0]) == ".json" {
		f, err := os.Open(schemaFiles[0])
		if err != nil {
			return nil, errors.WithStack(err)
		}
		defer f.Close()
		var schema introspection.Root
		if err := json.NewEncoder(f).Encode(&schema); err != nil {
			return nil, errors.WithStack(err)
		}
		return &schema, nil
	}

	if len(schemaFiles) == 0 {
		return nil, errHelp
	}

	for _, schema := range schemaFiles {
		switch filepath.Ext(schema) {
		case ".graphqls", ".graphql", ".gql":
		default:
			return nil, errHelp
		}
	}

	schema, err := loader.LoadSchema(schemaFiles...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return schema, nil
}

func requestSchema(endpoint string, queries, headers []string) (*introspection.Root, error) {
	client := graphql.NewClient(endpoint)

	req := gqlcli.NewRequest(introspection.Query)
	if len(queries) > 0 {
		for _, q := range queries {
			qSlice := strings.Split(q, "=")
			if len(qSlice) != 2 {
				continue
			}
			key, val := qSlice[0], qSlice[1]
			req.Var(key, val)
		}
	}

	if len(headers) > 0 {
		for _, h := range headers {
			hSlice := strings.Split(h, ": ")
			if len(hSlice) != 2 {
				continue
			}
			key, val := hSlice[0], hSlice[1]
			req.Header.Set(key, val)
		}
	}

	var schema introspection.Root
	if err := client.Run(context.Background(), req, &schema); err != nil {
		return nil, errors.WithStack(err)
	}
	return &schema, nil
}
