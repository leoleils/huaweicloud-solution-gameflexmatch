package cliapp

import (
	"github.com/urfave/cli/v2"
	"solutionbuild/gameflexmatch/common"
)
// 这里增加命令标识标签


func (c *CliApp) getFlagService() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagService,
		Usage:    "service name, Options: [fleetmanager, appgateway, aass]",
		Value:    "",
		Required: true,
	}
}

func (c *CliApp) getFlagConfig() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     common.FlagConfig,
		Usage:    "config path for service: [fleetmanager, appgateway, aass]",
		Value:    "",
		Required: false,
	}
}

func (c *CliApp) getFlagMode() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagMode,
		Usage:    "mode for cipher, options: [decode | encode], default: encode",
		Value:    "encode",
		Required: false,
	}
}

func (c *CliApp) getFlagMethod() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagMethod,
		Usage:    "method for cipher, options: [gcm | rsa], default: gcm",
		Value:    "gcm",
		Required: false,
	}
}

func (c *CliApp) getFlagText() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagText,
		Usage:    "text to decode or encode for cipher, will use the text to cipher if it had",
		Value:    "",
		Required: false,
	}
}

func (c *CliApp) getFlagCreate() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagCreate,
		Usage:    "create gcm key and nonce, create rsa key file, create tls key file, options: [gcm, rsa, tls, jwt]",
		Value:    "",
		Required: false,
	}
}

func (c *CliApp) getFlagInitConf() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     common.FlagAuthConf,
		Usage:    "init conf using for cipher",
		Value:    "",
		Required: false, // 正式使用时改为true
	}
}

func (c *CliApp) getFlagStartBin() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     common.FlagStartBin,
		Usage:    "service start path for configuring deamon process",
		Value:    "",
		Required: false, // 正式使用时改为true
	}
}

func (c *CliApp) getFlagStartConf() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     common.FlagStartConfig,
		Usage:    "service config file for configuring deamon process",
		Value:    "",
		Required: false, // 正式使用时改为true
	}
}

func (c *CliApp) getFlagStartLog() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     common.FlagStartLog,
		Usage:    "service log path for configuring deamon process",
		Value:    "",
		Required: false, // 正式使用时改为true
	}
}


func (c *CliApp) getFlagAk() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagAk,
		Usage:    "hw cloud ak for tenant",
		Value:    "",
		Required: false, 
	}
}

func (c *CliApp) getFlagSk() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagSk,
		Usage:    "hw cloud sk for tenant",
		Value:    "",
		Required: false, 
	}
}

func (c *CliApp) getFlagObsName() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagObsName,
		Usage:    "hw cloud obs bucket name, default: ObsBucketName in init.conf'",
		Value:    "",
		Required: false, 
	}
}

func (c *CliApp) getFlagCreateOBSBucket() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     common.FlagCreateOBSBucket,
		Usage:    "create obs bucket or not, default: false",
		Required: false,
	}
}

func (c *CliApp) getFlagKeyPairName() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagKeyPairName,
		Usage:    "hw cloud keypair name, default: KeypairName in init.conf'",
		Value:    "",
		Required: false,
	}
}

func (c *CliApp) getFlagCreateKeyPair() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     common.FlagCreateKeyPair,
		Usage:    "create keypair or not, default: false",
		Required: false,
	}
}

func (c *CliApp) getFlagCreateLtsAgency() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     common.FlagCreateLtsAgency,
		Usage:    "create lts agency or not, default false",
		Required: false,
	}
}

func (c *CliApp) getFlagUploadAuxproxy() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     common.FlagUploadAuxproxy,
		Usage:    "upload auxproxy or not, default: ./bin/auxproxy.zip",
		Required: false,
	}
}

func (c *CliApp) getFlagAuxproxyPath() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagAuxproxyPath,
		Usage:    "auxproxy path",
		Value:    "",
		Required: false,
	}
}

func (c *CliApp) getFlagRegion() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagRegion,
		Usage:    "init cloud env of region, default: Region in 'init.conf', if have more, will use first",
		Value:    "",
		Required: false,
	}
}

func (c *CliApp) getFlagUploadFile() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     common.FlagUploadFile,
		Usage:    "upload file to obs",
		Value:    "",
		Required: false,
	}
}

func (c *CliApp) getFlagLtsAgencyName() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     common.FlagLtsAgencyName,
		Usage:    "lts agency name, default: LtsAgencyName in 'init.conf'",
		Value:    "",
		Required: false,
	}
}