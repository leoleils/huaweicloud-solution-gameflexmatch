package cliapp

import (
	service "solutionbuild/gameflexmatch/action"

	"github.com/urfave/cli/v2"
)

const (
	AppName 	= "sac-gfm"
	AppUsage 	= "Usages: game-flex-match run script"
	AppVersion	= "v1.0"

	CommandStart = "start"
	CommandStop = "stop"
	CommandRestart = "restart"
	CommandInstall = "install"
	CommandCipher = "cipher"
	CommandInitHwcloud = "init-hwcloud"

)

type CliApp struct {
	action *service.Action
}

func NewGfmApp() *CliApp {
	
	return &CliApp{
		action: service.NewAction(),
	}
}

func (c *CliApp) GetCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.Version = AppVersion
	app.Commands = []*cli.Command {
		// 这里增加各个子命令
		c.getSubCommandCipher(),
		c.getSubCommandInstall(),
		c.getSubCommandStart(),
		c.getSubCommandStop(),
		c.getSubCommandRestart(),
	}
	return app
}

// get sub command start
func (c *CliApp) getSubCommandStart() *cli.Command {
	cmd := &cli.Command{
		Name: "start",
		Usage: "start a service, support: [fleetmanager, appgateway, aass]",
		Action: c.action.Start,
		Flags: []cli.Flag{
			c.getFlagService(),
		},
	}
	return cmd
}

// get sub command start
func (c *CliApp) getSubCommandStop() *cli.Command {
	cmd := &cli.Command{
		Name: "stop",
		Usage: "stop a service, support: [fleetmanager, appgateway, aass]",
		Action: c.action.Stop,
		Flags: []cli.Flag{
			c.getFlagService(),
		},
	}
	return cmd
}

// get sub command start
func (c *CliApp) getSubCommandRestart() *cli.Command {
	cmd := &cli.Command{
		Name: "restart",
		Usage: "restart a service, support: [fleetmanager, appgateway, aass]",
		Action: c.action.Restart,
		Flags: []cli.Flag{
			c.getFlagService(),
		},
	}
	return cmd
}

/*
get sub command instll
1. install fleetmanager env
2. install appgateway env
3. install aass env
4. install console env
*/
func (c *CliApp) getSubCommandInstall() *cli.Command {
	cmd := &cli.Command{
		Name: "install",
		Usage: "install service info, can help you to install fleetmanager/appgateway/aass",
		Action: c.action.Install,
		Flags: []cli.Flag{
			c.getFlagService(),
			c.getFlagInitConf(),
			c.getFlagStartLog(),
		},
	}
	return cmd
}

/* 
get sub command cipher:
1. generate gcm key and nonce
2. generate rsa public and private key
2. cipher key info
*/
func (c *CliApp) getSubCommandCipher() *cli.Command {
	cmd := &cli.Command{
		Name: "cipher",
		Usage: "cipher for key info, can help you encode or decode key info, and generate key file for gcm/rsa/tls/jwt",
		Action: c.action.Cipher,
		Flags: []cli.Flag{
			c.getFlagMode(),
			c.getFlagMethod(),
			c.getFlagText(),
			c.getFlagInitConf(),
			c.getFlagCreate(),
		},
	}
	return cmd
}

