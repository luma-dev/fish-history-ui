package main

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	// "github.com/go-openapi/swag"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/rs/cors"

	"github.com/luma-dev/fishis/projects/cli/filter"
	"github.com/luma-dev/fishis/projects/cli/gen/models"
	"github.com/luma-dev/fishis/projects/cli/gen/restapi"
	"github.com/luma-dev/fishis/projects/cli/gen/restapi/operations"
	"github.com/luma-dev/fishis/projects/cli/parser"
)

//go:embed web-ui-dist/*
var root embed.FS

var (
	port = flag.Int("port", 3210, "Port to run this service on")
)

type contentsFS struct {
	content embed.FS
}

func (c contentsFS) Open(name string) (fs.File, error) {
	return c.content.Open(path.Join("web-ui-dist", name))
}

func createAPIHandler(h []models.History) http.Handler {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	swaggerSpec.Spec().BasePath = "/api"

	api := operations.NewFishHistoryAPI(swaggerSpec)
	api.ChunkHistoryHandler = operations.ChunkHistoryHandlerFunc(
		func(params operations.ChunkHistoryParams) middleware.Responder {
			if err := params.Body.Validate(strfmt.NewFormats()); err != nil {
				return operations.NewChunkHistoryBadRequest()
			}

			unit := models.TimeUnitAll
			if params.Body.Unit != nil {
				unit = *params.Body.Unit
			}

			var limit int64 = -1
			if params.Body.Limit != nil {
				limit = *params.Body.Limit
			}

			blocks := filter.ChunkHistory(h, unit, limit, params.Body.Filter)
			return operations.NewChunkHistoryOK().WithPayload(&blocks)
		})
	return api.Serve(nil)
}

func main() {
	mux := http.NewServeMux()

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadFile(path.Join(home, ".local/share/fish/fish_history"))
	if err != nil {
		log.Fatal(err)
	}
	h := parser.ParseHistories(bytes)

	fs := http.FileServer(http.FS(contentsFS{root}))
	mux.Handle("/", fs)
	mux.Handle("/api/", createAPIHandler(h))

	log.Printf("Listening on :%d...\n", *port)
	handler := cors.Default().Handler(mux)
	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), handler)
	if err != nil {
		log.Fatal(err)
	}
}
