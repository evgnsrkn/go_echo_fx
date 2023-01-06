### Project
This project is being developed to learn golang, fx, echo, gorm.
___
### Makefile
__To build, run, compile or clean project use Makefile:__ </br>
- `make run` run application
- `make build` compile main.go and put binary in `bin` directory
- `make clean` delete `bin` directory
- `make compile` compile binaries for linux, macos and windows in `bin` directory
- `make live-reload` run application with live reloading

__You also can run application with live reloading using [air](https://github.com/cosmtrek/air):__
1. Install air using `go install github.com/cosmtrek/air@latest`
2. Use `make live-reload` command for run application with live reloading
