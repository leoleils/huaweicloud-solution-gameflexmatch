package action

import (
	"fmt"
	"os/exec"

	common "solutionbuild/gameflexmatch/common"

	"github.com/urfave/cli/v2"
)

type Action struct{

}

func NewAction() *Action {
	return &Action{}
}

// start a service
func (a *Action) Start(c *cli.Context) error {
	service := c.String(common.FlagService)
	if err := common.ExecCmd(exec.Command("systemctl", "start", service)); err != nil {
		return err
	}
	fmt.Printf("Succeed to start %s\n", service)
	fmt.Printf("You alse can exec this command to start %s:\n", service)
	fmt.Printf("systemctl start %s\n", service)
	return nil
}

// stop a service
func (a *Action) Stop(c *cli.Context) error {
	service := c.String(common.FlagService)
	if err := common.ExecCmd(exec.Command("systemctl", "stop", service)); err != nil {
		return err
	}
	fmt.Printf("Succeed to stop %s\n", service)
	fmt.Printf("You alse can exec this command to stop %s:\n", service)
	fmt.Printf("systemctl stop %s\n", service)
	return nil
}

// restart a service
func (a *Action) Restart(c *cli.Context) error {
	service := c.String(common.FlagService)
	if err := common.ExecCmd(exec.Command("systemctl", "restart", service)); err != nil {
		return err
	}
	fmt.Printf("Succeed to restart %s\n", service)
	fmt.Printf("You alse can exec this command to restart %s:\n", service)
	fmt.Printf("systemctl restart %s\n", service)
	return nil
}

// cipher key info
func (a *Action) Cipher(c *cli.Context) error {
	cipherAction := NewCipherAction(c)
	initConf := a.getInitConf(c)

	create := c.String(Create)
	if create != "" {
		return cipherAction.Create(initConf)
	}

	mode := c.String(Mode)
	method := c.String(Method)
	fmt.Printf("mode: %s method: %s \n", mode, method)
	if mode == "encode" && method == "gcm" {
		return cipherAction.GcmEncode(initConf)
	} else if mode == "encode" && method == "rsa" {
		return cipherAction.RsaEncode(initConf)
	} else if mode == "decode" && method == "gcm" {
		return cipherAction.GcmDecode(initConf)
	} else if mode == "decode" && method == "rsa" {
		return cipherAction.RsaDecode(initConf)
	}
	return fmt.Errorf("the 'mode' or 'method' must be provided correct if you want to use cipher," +
	"now privade mode: %s method: %s", mode, method)
}

// install env and other info for starting service normally 
func (a *Action) Install(c *cli.Context) error {
	installAction := NewInstallAction(c)
	initConf := a.getInitConf(c)
	service := c.String(common.FlagService)
	fmt.Printf("Would install service %s with initConf %s \n", service, initConf)
	log := c.String(common.FlagStartLog)
	if log == "" {
		log = fmt.Sprintf("/local/log/%s", service)
		fmt.Printf("Would use %s as the service log path", log)
	}
	param := map[string]string {
		"log": log,
	}
	if err := installAction.install(initConf, service, param); err != nil {
		return err
	}
	return nil
}

// init env for starting gfm on hwcloud
func (a *Action) InitHwcloud(c *cli.Context) error {
	initHwCloudAction := NewInitHwCloudAction(c)
	initConf := a.getInitConf(c)
	if c.Bool(common.FlagCreateKeyPair) {
		if err := initHwCloudAction.CreateKeyPair(initConf); err != nil {
			return err
		}
	}
	if c.Bool(common.FlagCreateOBSBucket) {
		if err := initHwCloudAction.CreateObsBucket(initConf); err != nil {
			return err
		}
	}


	if c.Bool(common.FlagUploadAuxproxy) {
		if err := initHwCloudAction.UploadAuxproxy(initConf); err != nil {
			return err
		}
	}

	uploadFile := c.String(common.FlagUploadFile)
	if uploadFile != "" {
		if err := initHwCloudAction.UploadFile(initConf); err != nil {
			return err
		}
	}

	if c.Bool(common.FlagCreateLtsAgency) {
		if err := initHwCloudAction.CreateLtsAgency(initConf); err != nil {
			return err
		}
	}

	return nil
}

func (a *Action) getInitConf(c *cli.Context) string {
	initConf := c.String(InitConf)
	if initConf == "" {
		initConf = common.DefaultConfPath + "/init.yaml"
	}

	fmt.Printf("init conf: %s\n", initConf)
	return initConf
}

