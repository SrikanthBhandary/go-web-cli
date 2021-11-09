package cmd

import (
	"html/template"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"

	"github.com/SrikanthBhandary/go-web-cli/tpl/cmd"
	"github.com/SrikanthBhandary/go-web-cli/tpl/controller"
	"github.com/SrikanthBhandary/go-web-cli/tpl/entity"
	"github.com/SrikanthBhandary/go-web-cli/tpl/repository"
	"github.com/SrikanthBhandary/go-web-cli/tpl/service"
)

var pkgName = ""

const AppName = "app"

var resources = []string{
	"controller",
	"service",
	"repository",
	"cmd",
	"entity",
	"config",
	"db",
	"logger",
}

func CreateApp() {
	os.Mkdir(AppName, 0777)
}

func CreateFolders() {
	for _, resource := range resources {
		os.MkdirAll(path.Join(AppName, resource), 0777)
	}
}

func ExecuteTemplates(content, resource string) {
	tmpl := template.Must(template.New("test").Parse(content))
	f, _ := os.Create(path.Join(AppName, resource, "todo.go"))
	tmpl.Execute(f, pkgName)
}

func CreateTemplates() {
	for _, resource := range resources {
		switch resource {
		case "controller":
			ExecuteTemplates(controller.Tmpl, resource)
		case "repository":
			ExecuteTemplates(repository.Tmpl, resource)
		case "service":
			ExecuteTemplates(service.Tmpl, resource)
		case "entity":
			ExecuteTemplates(entity.Tmpl, resource)
		case "cmd":
			ExecuteTemplates(cmd.Tmpl, resource)
		default:
			continue
		}
	}
}

func GetDependencies() {
	command := exec.Command("/bin/sh", "-c", "go mod init "+pkgName)
	command.Dir = "app"
	command.Run()
	command = exec.Command("/bin/sh", "-c", "go mod tidy")
	command.Dir = "app"
	command.Run()

}

var rootCmd = &cobra.Command{
	Use:   "web-cli",
	Short: "Code generation tool for webserer.",
	Long:  `CLI tool to generate the boiler plate code for go microservices.`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateApp()
		CreateFolders()
		CreateTemplates()
		GetDependencies()
	},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&pkgName, "pkgname", "", "Package name")
	rootCmd.MarkFlagRequired("pkgname")
}
