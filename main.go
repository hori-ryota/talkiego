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

package main

import "github.com/hori-ryota/talkiego/cmd"

//go:generate go-assets-builder -p assets -o assets/talkie.go -s /assets/Talkie/dist -v Talkie assets/Talkie/dist
//go:generate go-assets-builder -p assets -o assets/template.go -s /assets/template -v Template assets/template

func main() {
	cmd.Execute()
}
