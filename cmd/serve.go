// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/hori-ryota/talkiego/assets"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var servePort uint

var serveRootOption = RootOption{}
var customCSSs = []string{}
var customScripts = []string{}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve [index.md]",
	Short: "Serve markdown file",
	Long:  `Serve markdown file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		targetPath := "index.md"
		if len(args) >= 1 {
			targetPath = args[0]
		}

		log.Printf("targetPath = %s", targetPath)

		if !IsExistFile(targetPath) {
			return errors.Errorf("target file not found: targetPath='%s'", targetPath)
		}

		http.Handle("/assets/Talkie/dist/", http.StripPrefix("/assets/Talkie/dist", http.FileServer(assets.Talkie)))
		http.HandleFunc("/", handler(targetPath))

		log.Printf("ListenAndServe http://localhost:%d/\n", servePort)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", servePort), nil); err != nil {
			return errors.Wrap(err, "ListenAndServe failed")
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().UintVarP(&servePort, "port", "p", 3000, "listen port")
	serveCmd.Flags().BoolVarP(&serveRootOption.Wide, "wide", "w", false, "wide option")
	serveCmd.Flags().BoolVarP(&serveRootOption.Control, "control", "c", false, "control option")
	serveCmd.Flags().BoolVarP(&serveRootOption.Pointer, "pointer", "b", false, "pointer option")
	serveCmd.Flags().BoolVarP(&serveRootOption.Progress, "progress", "r", false, "progress option")
	serveCmd.Flags().BoolVarP(&serveRootOption.Backface, "backface", "f", false, "backface option")
	serveCmd.Flags().BoolVarP(&serveRootOption.NoTransition, "noTransition", "t", false, "noTransition option")
	serveCmd.Flags().BoolVarP(&serveRootOption.LinkShouldBlank, "linkShouldBlank", "l", false, "linkShouldBlank option")

	serveCmd.Flags().StringSliceVarP(&customCSSs, "customCSS", "s", []string{}, "custom css")
	serveCmd.Flags().StringSliceVarP(&customScripts, "customScripts", "j", []string{}, "custom scripts")
}

func handler(targetPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" || r.URL.Path == "/" || r.URL.Path == "index.html" {
			indexHandler(w, r, targetPath)
			return
		}

		fileRoot := filepath.Dir(targetPath)
		target := filepath.Join(fileRoot, r.URL.Path)
		http.ServeFile(w, r, target)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, targetPath string) {
	tmpl, err := template.New("index").Parse(string(assets.Template.Files["/index.html"].Data))
	if err != nil {
		err = errors.Wrap(err, "parse template for index.html failed)")
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	// parse
	bs, err := ioutil.ReadFile(targetPath)
	if err != nil {
		err = errors.Wrap(err, "read file failed")
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	param := TalkieParam{
		RootOption:    serveRootOption,
		CustomCSSs:    customCSSs,
		CustomScripts: customScripts,
	}
	if err := param.Parse(string(bs), divider); err != nil {
		err = errors.Wrap(err, "parse file failed")
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	if err := tmpl.Execute(w, param); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
