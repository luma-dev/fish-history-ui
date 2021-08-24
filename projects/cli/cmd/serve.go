package cmd

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/adrg/xdg"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-ptr/go-ptr/ptr"
	"github.com/rs/cors"
	"github.com/spf13/cobra"

	"github.com/luma-dev/fish-history-ui/projects/cli/filter"
	"github.com/luma-dev/fish-history-ui/projects/cli/gen/models"
	"github.com/luma-dev/fish-history-ui/projects/cli/gen/restapi"
	"github.com/luma-dev/fish-history-ui/projects/cli/gen/restapi/operations"
	"github.com/luma-dev/fish-history-ui/projects/cli/parser"
	"github.com/luma-dev/fish-history-ui/projects/cli/util"
)

func createAPIHandler(basePath string, histories []models.History, tzName string, offset int) http.Handler {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	swaggerSpec.Spec().BasePath = basePath

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

			blocks := filter.ChunkHistory(histories, unit, limit, params.Body.Filter)
			return operations.NewChunkHistoryOK().WithPayload(&blocks)
		})
	api.GetTimezoneHandler = operations.GetTimezoneHandlerFunc(
		func(params operations.GetTimezoneParams) middleware.Responder {
			timezone := models.TimezoneResponse{
				Name:   ptr.NewString(tzName),
				Offset: ptr.NewInt64(int64(offset)),
			}
			return operations.NewGetTimezoneOK().WithPayload(&timezone)
		})
	return api.Serve(nil)
}

func getOffset(prefered string) (string, int, error) {
	if prefered == "" || prefered == "system" {
		t := time.Now()
		name, offset := t.Zone()
		return name, offset, nil
	}
	t := time.Now()
	loc, err := time.LoadLocation(prefered)
	if err != nil {
		return "", 0, err
	}
	name, offset := t.In(loc).Zone()
	return name, offset, nil
}

func error500(w http.ResponseWriter) {
	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
}

var (
	port     int
	host     string
	timezone string
	open     bool
	root     embed.FS
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 3210, "port number to serve (default 4310)")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "b", "", "interface to bind (default 0.0.0.0)")
	rootCmd.PersistentFlags().StringVarP(&timezone, "timezone", "t", "", "Timezone used in calendar Web UI (e.g. UTC, Japan, Etc/GMT+3. default \"system\")")
	rootCmd.PersistentFlags().BoolVarP(&open, "open", "", false, "Open browser.")
}

var rootCmd = &cobra.Command{
	Use:  "serve",
	Long: "Serve Fish shell history visualization Web UI.",
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()

		if len(args) >= 2 {
			log.Fatal(errors.New("not implemented to support multiple input files"))
		}
		var historyPath string
		if len(args) == 1 {
			historyPath = args[0]
		} else {
			// Reference: https://fishshell.com/docs/current/cmds/history.html
			// Fish history default path is $XDG_DATA_HOME/fish/fish_history
			historyPath = path.Join(xdg.DataHome, "fish/fish_history")
		}

		bytes, err := ioutil.ReadFile(historyPath)
		if err != nil {
			log.Fatal(err)
		}

		tzName, offset, err := getOffset(timezone)
		if err != nil {
			log.Fatal(err)
		}

		histories := util.ShiftHistories(parser.ParseHistories(bytes), int64(offset))

		f := http.FS(root)

		index, err := f.Open(path.Join("web-ui-dist", "index.html"))
		if err != nil {
			log.Fatal(err)
		}
		indexStat, err := index.Stat()
		if err != nil {
			log.Fatal(err)
		}

		mux.Handle("/api/", createAPIHandler("/api", histories, tzName, offset))
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fallback := func() {
				http.ServeContent(w, r, indexStat.Name(), indexStat.ModTime(), index)
			}
			file, err := f.Open(path.Join("web-ui-dist", r.URL.Path))
			if errors.Is(err, fs.ErrNotExist) {
				fallback()
				return
			} else if err != nil {
				error500(w)
				return
			}
			stat, err := file.Stat()
			if err != nil {
				error500(w)
				return
			}
			if stat.IsDir() {
				fallback()
				return
			}
			http.ServeContent(w, r, stat.Name(), stat.ModTime(), file)
		})

		log.Printf("Listening on %s:%d...\n", host, port)
		handler := cors.Default().Handler(mux)
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), handler)
		if err != nil {
			log.Fatal(err)
		}
	},
}
