// Package cmd
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/benkeil/projgen/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	giturls "github.com/whilp/git-urls"
	"os"
	"path/filepath"
)

var projectName string
var devRoot string
var vcsProvider string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate the new repository.",
	Long:    `Generate the new repository.`,
	Args:    cobra.ExactArgs(1),
	Example: heredoc.Doc(`$ projgen generate templates/typescript -n my-project`),
	Run: func(cmd *cobra.Command, args []string) {
		repository := args[0]
		_, err := giturls.Parse(repository)
		cobra.CheckErr(err)

		projectPath, err := pkg.CreateProject(projectName)
		cobra.CheckErr(err)

		stat, err := os.Stat(repository)
		switch mode := stat.Mode(); {
		case mode.IsDir():
			err = pkg.CopyA(filepath.Join(repository, ".*"), projectPath)
			cobra.CheckErr(err)
		default:
			err = pkg.CloneTemplate(repository, projectPath)
			cobra.CheckErr(err)
		}

		template, err := pkg.ReadTemplateFile(args[0])
		cobra.CheckErr(err)

		if devRoot == "" {
			devRoot = viper.GetString("dev-root")
		}
		if vcsProvider == "" {
			vcsProvider = viper.GetString("vcs-provider")
		}

		params, err := pkg.ReadParams(pkg.InputParams{
			ProjectPath: projectPath,
			ProjectName: projectName,
			DevRoot:     devRoot,
			VcsProvider: vcsProvider,
		})
		cobra.CheckErr(err)

		for _, step := range template.Steps {
			executableStep, err := step.Transform(params)
			cobra.CheckErr(err)
			err = executableStep.Execute()
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&projectName, "project-name", "n", "", "The project name")
	generateCmd.Flags().StringVarP(&devRoot, "dev-root", "r", "", "The development root directory")
	generateCmd.Flags().StringVarP(&vcsProvider, "vcs-provider", "p", "", "The VCS provider")
	generateCmd.MarkFlagRequired("project-name")
	viper.BindPFlag("dev-root", generateCmd.Flags().Lookup("dev-root"))
	viper.BindPFlag("vcs-provider", generateCmd.Flags().Lookup("vcs-provider"))
}
