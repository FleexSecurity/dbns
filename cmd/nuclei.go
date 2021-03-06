/*
Copyright © 2021 FleexSecurity

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"
	"log"

	"github.com/FleexSecurity/dbns/config"
	"github.com/FleexSecurity/dbns/nuclei"
	"github.com/FleexSecurity/dbns/nuclei/repositories"
	"github.com/FleexSecurity/dbns/nuclei/services"
	"github.com/spf13/cobra"
)

var (
	psqlDB *sql.DB
)

// nucleiCmd represents the nuclei command
var nucleiCmd = &cobra.Command{
	Use:   "nuclei",
	Short: "Nuclei Scanner command",
	Long:  "Nuclei Scanner command",
	Run: func(cmd *cobra.Command, args []string) {
		listPath, _ := cmd.Flags().GetString("list")
		url, _ := cmd.Flags().GetString("url")
		info, _ := cmd.Flags().GetBool("info")
		if url == "" && listPath == "" {
			log.Fatal("ERR:", nuclei.ErrInvalidUrlOrList)
		}
		psqlDB = config.Connect()
		repository := repositories.PsqlNucleiRepository{
			DB:    psqlDB,
			Table: "nuclei",
		}
		service := services.NucleiService{
			Repository: repository,
		}
		err := service.Scan(url, listPath, info)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(nucleiCmd)

	nucleiCmd.Flags().StringP("list", "l", "", "path to file containing a list of target URLs/hosts to scan (one per line)")
	nucleiCmd.Flags().StringP("url", "u", "", "target URLs/hosts to scan")
	nucleiCmd.Flags().BoolP("info", "i", false, "scan also for info severity")
}
