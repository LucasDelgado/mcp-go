package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"ChopSetup",
		"1.0.0",
	)

	setupTool := mcp.NewTool(
		"obtener-setup",
		mcp.WithDescription("Obtener el setup de la aplicación"),
		mcp.WithString("setup_id", mcp.Required(), mcp.Description("El id del setup a obtener")),
	)

	s.AddTool(
		setupTool,
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			arguments := req.GetArguments()
			setup_id := arguments["setup_id"].(string)

			url := fmt.Sprintf("http://localhost:8080/backoffice/setup/%s", setup_id)

			request, _ := http.NewRequest("GET", url, nil)
			request.Header.Set("Content-Type", "application/json")

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				return mcp.NewToolResultErrorFromErr("Error al obtener el setup", err), nil
			}
			defer response.Body.Close()

			if response.StatusCode != 200 {
				return mcp.NewToolResultErrorFromErr("Error al obtener el setup", errors.New("error al obtener el setup")), nil
			}

			var setup Setup
			err = json.NewDecoder(response.Body).Decode(&setup)
			if err != nil {
				return mcp.NewToolResultErrorFromErr("Error al obtener el setup", err), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("%+v", setup)), nil
		},
	)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v", err)
	}
}
