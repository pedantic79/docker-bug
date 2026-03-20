package main

import (
	"fmt"

	"github.com/docker/cli/cli/command"
	"github.com/docker/compose/v5/pkg/compose"
)

func main() {
	dockerCli, _ := command.NewDockerCli()
	composeService, _ := compose.NewComposeService(dockerCli)
	fmt.Println(composeService) // avoid compiler complaining about unused variable
}
