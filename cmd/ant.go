// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
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
	"bytes"
	"fmt"
	"log"

	"github.com/EDDYCJY/fake-useragent"
	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

// antCmd represents the ant command
var antCmd = &cobra.Command{
	Use:   "ant",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flags().Lookup("url").Value.String()
		fmt.Println("ant called:", url)
		if url == "" {
			fmt.Println("input a valid url")
			return
		}
		Start(url)
	},
}

func init() {
	rootCmd.AddCommand(antCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// antCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	antCmd.Flags().StringP("url", "u", "https://detail.1688.com/offer/578840012945.html?spm=a261y.7663282.descBanner.15.303731e481lUtC", "fetch url")
}

//Start 开始
func Start(url string) {
	c := colly.NewCollector(
		colly.UserAgent(browser.Random()),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println("onhtml:", e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Printf("Proxy Address: %s\n", r.Request.URL)
		log.Printf("Body: %s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	c.Visit(url)
}
