package action

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	common "solutionbuild/gameflexmatch/common"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	kps "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kps/v3"
	kpsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kps/v3/model"
	"github.com/urfave/cli/v2"
	"huaweicloud.com/esdk-obs-go/obs/v3"

	iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iamModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"

	coreregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

type InitHwCloudAction struct {
	c *cli.Context
}

func NewInitHwCloudAction(c *cli.Context) *InitHwCloudAction {
	return &InitHwCloudAction{
		c: c,
	}
}

// 内网不通
func (a *InitHwCloudAction) CreateKeyPair(authConf string) error {
	var err error
	keypairType := kpsModel.GetCreateKeypairActionTypeEnum().SSH
	keypairScope := kpsModel.GetCreateKeypairActionScopeEnum().DOMAIN
	req := &kpsModel.CreateKeypairRequest{}
	keypairName := a.c.String(common.FlagKeyPairName)
	if keypairName == "" {
		fmt.Printf("start create key pair using default: %s\n", keypairName)
		keypairName, err = common.GetConfKey(authConf, common.SectionCloud, common.KeypairName)
		if err != nil {
			return err
		}
	}
	fmt.Printf("start create key pair using %s\n", keypairName)
	req.Body = &kpsModel.CreateKeypairRequestBody{
		Keypair: &kpsModel.CreateKeypairAction{
			Name:  keypairName,
			Type:  &keypairType,
			Scope: &keypairScope,
			KeyProtection: &kpsModel.KeyProtection{
				Encryption: &kpsModel.Encryption{
					Type: kpsModel.GetEncryptionTypeEnum().DEFAULT,
				},
			},
		},
	}
	kpsClient, err := a.getKpsClient(authConf)
	if err != nil {
		return err
	}
	res, errR := kpsClient.CreateKeypair(req)
	if errR != nil {
		return errR
	}
	fmt.Printf("success to create keypair: %+v", res)
	return nil
}

func (a *InitHwCloudAction) CreateObsBucket(authConf string) error {
	var err error
	obsBucketName := a.c.String(common.FlagObsName)
	if obsBucketName == "" {
		obsBucketName, err = common.GetConfKey(authConf, common.SectionCloud, common.ObsBucketName)
		if err != nil {
			return err
		}
	}
	if obsBucketName == "" {
		obsBucketName = "gfm-obs-" + common.StringWithCharset(6)
		obsBucketName = strings.ToLower(obsBucketName)
	}
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.ObsBucketName, obsBucketName), authConf)); err != nil {
		return err
	}
	epid, err := common.GetConfKey(authConf, common.SectionCloud, common.EnterpriseProjectId)
	if err != nil {
		return err
	}
	req := &obs.CreateBucketInput{
		Bucket: obsBucketName,
		StorageClass: obs.StorageClassStandard,
		ACL: obs.AclPrivate,
		Epid: epid,
	}
	client, err := a.getObsClient(authConf)
	if err != nil {
		return err
	}
	res, err := client.CreateBucket(req)
	if err != nil {
		return err
	}
	fmt.Printf("success to create obs bucket: %+v", res)
	return nil
}

func (a *InitHwCloudAction) UploadAuxproxy(authConf string) error {
	auxproxyPath := a.c.String(common.FlagAuxproxyPath)
	if auxproxyPath == "" {
		auxproxyPath = common.RootPath + common.AuxproxyPath
	}
	return a.UploadToObs(authConf, auxproxyPath)
}

func (a *InitHwCloudAction) UploadFile(authConf string) error {
	filePath := a.c.String(common.FlagUploadFile)
	return a.UploadToObs(authConf, filePath)
}

func (a *InitHwCloudAction) UploadToObs(authConf string, fileKey string) error {
	var err error
	obsBucketName := a.c.String(common.FlagObsName)
	if obsBucketName == "" {
		obsBucketName, err = common.GetConfKey(authConf, common.SectionCloud, common.ObsBucketName)
		if err != nil {
			return err
		}
	}
	
	client, err := a.getObsClient(authConf)
	if err != nil {
		return err
	}
	file, err := os.Open(fileKey)
	if err != nil {
		return err
	}
	req := &obs.PutObjectInput{}
	req.Bucket = obsBucketName
	req.Body = file
	req.Key = common.ObsAuxproxyName
	res, err := client.PutObject(req)
	if err != nil {
		return err
	}
	fmt.Printf("success to upload auxproxy info to obs bucket: %s, response: %+v", obsBucketName, res)
	return nil
}

func (a *InitHwCloudAction) CreateLtsAgency(authConf string) error {
	var err error
	ltsAgencyName := a.c.String(common.FlagLtsAgencyName)
	if ltsAgencyName == "" {
		ltsAgencyName, err = common.GetConfKey(authConf, common.SectionCloud, common.LtsAgencyName)
		if err != nil {
			return err
		}
	}

	domainId, err := common.GetConfKey(authConf, common.SectionCloud, common.DomianId)
	if err != nil {
		return err
	}
	client, err := a.getIamClient(authConf) 
	if err != nil {
		return err
	}
	ltsDomainId := common.LtsOpDomainName
	description := "game flex match lts agency"
	req := &iamModel.CreateAgencyRequest{
		Body: &iamModel.CreateAgencyRequestBody{
			Agency: &iamModel.CreateAgencyOption{
				Name: ltsAgencyName,
				DomainId: domainId,
				TrustDomainName: &ltsDomainId,
				Description: &description,
			},
		},
	}
	res, err := client.CreateAgency(req)
	if err != nil {
		return err
	}
	fmt.Printf("success to create lts agency: %+v", res)
	return nil
}

func (a *InitHwCloudAction) getAkSkRegion(authConf string) (string, string, string, error) {
	var err error
	akValue := a.c.String("ak")
	if akValue == "" {
		akValue, err = common.GetConfKey(authConf, common.SectionCloud, akKey)
		if err != nil {
			return "", "", "", err
		}
	}
	skValue := a.c.String("sk")
	if skValue == "" {
		skValue, err = common.GetConfKey(authConf, common.SectionCloud, skKey)
		if err != nil {
			return "", "", "", err
		}
	}
	regionValue := a.c.String("region")
	if regionValue == "" {
		regionValue, err = common.GetConfKey(authConf, common.SectionCloud, common.CloudRegion)
		if err != nil {
			return "", "", "", err
		}
		regionValue = strings.Split(regionValue, ",")[0]
	}
	return akValue, skValue, regionValue, nil
}

func (a *InitHwCloudAction) getKpsClient(authConf string) (*kps.KpsClient, error) {
	akValue, skValue, regionValue, err := a.getAkSkRegion(authConf)
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf(common.EndpointFormat, "kms", regionValue, common.MyhuaweicloudEndpoint)
	iamEndpoint := fmt.Sprintf(common.EndpointFormat, "iam", regionValue, common.MyhuaweicloudEndpoint)
	auth := basic.NewCredentialsBuilder().
		WithAk(akValue).
		WithSk(skValue).
		WithIamEndpointOverride(iamEndpoint).
		Build()
	
	client := kps.NewKpsClient(
		kps.KpsClientBuilder().
			WithRegion(coreregion.NewRegion(regionValue, endpoint)).
			WithCredential(auth).
			Build())
	return client, nil
}

func (a *InitHwCloudAction) getObsClient(authConf string) (*obs.ObsClient, error) {
	akValue, skValue, regionValue, err := a.getAkSkRegion(authConf)
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf(common.EndpointFormat, "obs", regionValue, common.UlanqabEndpoint)
	client, err := obs.New(akValue, skValue, endpoint)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (a *InitHwCloudAction) getIamClient(authConf string) (*iam.IamClient, error) {
	akValue, skValue, regionValue, err := a.getAkSkRegion(authConf)
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf(common.EndpointFormat, "iam", regionValue, common.MyhuaweicloudEndpoint)
	auth := global.NewCredentialsBuilder().
		WithAk(akValue).
		WithSk(skValue).
		WithIamEndpointOverride(endpoint).
		Build()
	
	client := iam.NewIamClient(
		iam.IamClientBuilder().
			WithRegion(coreregion.NewRegion(regionValue, endpoint)).
			WithCredential(auth).
			Build())
	return client, nil
}