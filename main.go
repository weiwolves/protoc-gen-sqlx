package main

import (
	"github.com/gogo/protobuf/vanity/command"
	"github.com/weiwolves/protoc-gen-sqlx/plugin"
)

func main() {
	response := command.GeneratePlugin(command.Read(), &plugin.SqlxPlugin{}, ".pb.sqlx.go")
	for _, file := range response.GetFile() {
		file.Content = plugin.CleanImports(file.Content)
	}
	command.Write(response)
}
