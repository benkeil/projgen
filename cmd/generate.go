// Package cmd
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package cmd

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/benkeil/projgen/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	giturls "github.com/whilp/git-urls"
	"os"
)

var projectName string
var devRoot string
var vcsProvider string
var vcsUser string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate TEMPLATE",
	Short: "Generates a new repository",
	Long: heredoc.Doc(`
		Generates a new repository from a template.
	`),
	Args: cobra.ExactArgs(1),
	Example: pkg.Doc(`
		to use a local directory or repository
		$ projgen generate ./templates/typescript -n my-project
		to use a remote repository
		$ projgen generate git@github.com:user/template.git -n my-project
	`),
	PreRun: func(cmd *cobra.Command, args []string) {
		if devRoot == "" {
			devRoot = viper.GetString("dev-root")
		}
		if vcsProvider == "" {
			vcsProvider = viper.GetString("vcs-provider")
		}
		if vcsUser == "" {
			switch user := viper.GetString("vcs-user"); {
			case user != "":
				vcsUser = user
			default:
				//cobra.CheckErr(errors.New("vcs-user not set, use the --vcs-user flag or configure  it in your ~/.config/projgen/config.yaml"))
			}
		}

		params := &pkg.Params{
			ProjectName: projectName,
		}

		projectTemplate, err := pkg.ReadRemoteTemplateFile(args[0])
		cobra.CheckErr(err)

		for key, value := range projectTemplate.Overrides {
			newValue := pkg.OverrideParam(value, *params)
			switch key {
			case "ProjectName":
				projectName = newValue
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		repository := args[0]
		_, err := giturls.Parse(repository)
		cobra.CheckErr(err)

		projectPath, err := pkg.CreateProject(projectName)
		cobra.CheckErr(err)
		fmt.Printf("create project from template at %v\n", projectPath)

		stat, err := os.Stat(repository)
		cobra.CheckErr(err)

		switch mode := stat.Mode(); {
		case mode.IsDir():
			err = pkg.CopyTemplate(repository, projectPath)
			cobra.CheckErr(err)
		default:
			err = pkg.CloneTemplate(repository, projectPath)
			cobra.CheckErr(err)
		}

		template, err := pkg.ReadTemplateFile(args[0])
		cobra.CheckErr(err)

		params := &pkg.Params{
			ProjectPath: projectPath,
			ProjectName: projectName,
			DevRoot:     devRoot,
			VcsProvider: vcsProvider,
		}

		for _, step := range template.Steps {
			executableStep, err := step.Transform(params)
			cobra.CheckErr(err)
			err = executableStep.Execute()
			cobra.CheckErr(err)
		}

		err = pkg.Cleanup(projectPath)
		cobra.CheckErr(err)

		err = pkg.GitInit(projectPath)
		cobra.CheckErr(err)

		err = pkg.GitCommit(projectPath)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&projectName, "project-name", "n", "", "The project name")
	//generateCmd.Flags().StringVarP(&devRoot, "dev-root", "r", "", "The development root directory")
	//generateCmd.Flags().StringVarP(&vcsProvider, "vcs-provider", "p", "", "The VCS provider")
	//generateCmd.Flags().StringVarP(&vcsUser, "vcs-user", "u", "", "The VCS user")
	generateCmd.MarkFlagRequired("project-name")
	viper.BindPFlag("dev-root", generateCmd.Flags().Lookup("dev-root"))
	viper.BindPFlag("vcs-provider", generateCmd.Flags().Lookup("vcs-provider"))
	viper.BindPFlag("vcs-user", generateCmd.Flags().Lookup("vcs-user"))
}
