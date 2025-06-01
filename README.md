# Chop MCP (Message Control Protocol)

Este proyecto implementa un servidor MCP (Message Control Protocol) para interactuar con configuraciones de setup en una aplicación. El servidor está desarrollado en Go y proporciona una herramienta para obtener información de setup a través de una API local.

## Descripción

El servidor MCP implementa una herramienta llamada "obtener-setup" que permite recuperar información detallada de configuraciones de setup mediante un ID específico. La información del setup incluye detalles internos y externos sobre la configuración de la aplicación.

## Requisitos

- Go 1.x o superior
- Acceso a un servidor local en el puerto 8080 que proporcione los endpoints de setup

## Estructura del Proyecto

```
.
├── main.go         # Archivo principal con la implementación del servidor MCP
├── setup.go        # Definición de estructuras de datos para el setup
├── go.mod          # Archivo de dependencias de Go
└── go.sum          # Checksums de las dependencias
```

## Herramientas Disponibles

### obtener-setup

Esta herramienta permite obtener la información de un setup específico.

**Parámetros:**
- `setup_id` (string, requerido): El ID del setup que se desea obtener

## Uso

El servidor se ejecuta y escucha en la entrada/salida estándar para comandos MCP. Para obtener un setup, se debe proporcionar el ID del setup como parámetro.

## Dependencias Principales

- github.com/mark3labs/mcp-go/mcp
- github.com/mark3labs/mcp-go/server

## Endpoints

El servidor se comunica con un servidor local que debe estar ejecutándose en:
- `http://localhost:8080/backoffice/setup/{setup_id}`

## Manejo de Errores

El servidor maneja varios tipos de errores:
- Errores de conexión con el servidor local
- Errores en la respuesta del servidor (códigos diferentes a 200)
- Errores en la decodificación de la respuesta JSON

## Contribución

Para contribuir al proyecto:
1. Fork del repositorio
2. Crear una rama para tu feature
3. Commit de tus cambios
4. Push a la rama
5. Crear un Pull Request 

## Development
- Agregarlo al client modo Dev
```JSON  
    "ChopSetup": {
      "command": "/Users/lucdelgado/Works/chop-mcp/chop-mcp",
      "args": []
    }
```
- Correr el debug
- `npx @modelcontextprotocol/inspector go run .`
