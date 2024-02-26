package gameflexmatch

import (
	"fmt"
	"log"
	"os"

	"solutionbuild/gameflexmatch/cliapp"
	"solutionbuild/gameflexmatch/common"

)

var (
	command string
	input string

	optionConfig string
)

func Run() {
	var err error
	common.RootPath, err = os.Getwd()
	if err != nil {
		fmt.Printf("get root path err: %+v", err)
		os.Exit(1)
	}
	fmt.Printf("the work dir is: %s\n", common.RootPath)
	cliApp := cliapp.NewGfmApp()

	err = cliApp.GetCliApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}