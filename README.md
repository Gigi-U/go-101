# go-101

1) create Github repository
2) clone repository to local folder
3) open folder in VSC
4) command: go mod init github.com/Gigi-U/go-101.git (a go.mod file is created)
5) Ctrl+Shift+p : search PREFERENCES: Open Workspace Settings(JSON) - .vscode folder and settings.json file are created.
6) in settings.json add: 
            "go.lintTool": "golangci-lint",
            "go.lintFlags": ["--fast"]
7) create main.go file
8) create folder(files inside the folder must have package named after the folder)
8) add to main.go access to  those packages. Example:

        package main

        import (
            "github.com/Gigi-U/go-101.git/datatypes"
        )

        func main() {
            datatypes.Hello()
        }      


Para correr una función sola: 