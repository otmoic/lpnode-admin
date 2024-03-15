// Code generated by goa v3.11.0, DO NOT EDIT.
//
// installCtrlPanel HTTP client types
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	installctrlpanel "admin-panel/gen/install_ctrl_panel"

	goa "goa.design/goa/v3/pkg"
)

// ListInstallRequestBody is the type of the "installCtrlPanel" service
// "listInstall" endpoint HTTP request body.
type ListInstallRequestBody struct {
	// type of service installed
	InstallType string `form:"installType" json:"installType" xml:"installType"`
}

// InstallLpClientRequestBody is the type of the "installCtrlPanel" service
// "installLpClient" endpoint HTTP request body.
type InstallLpClientRequestBody struct {
	// ammClientInstallConfig
	SetupConfig *AmmClientSetupConfigRequestBody `form:"setupConfig" json:"setupConfig" xml:"setupConfig"`
}

// UninstallLpClientRequestBody is the type of the "installCtrlPanel" service
// "uninstallLpClient" endpoint HTTP request body.
type UninstallLpClientRequestBody struct {
	// ammClientUninstallConfig
	SetupConfig *AmmClientUnSetupConfigRequestBody `form:"setupConfig" json:"setupConfig" xml:"setupConfig"`
}

// InstallDeploymentRequestBody is the type of the "installCtrlPanel" service
// "installDeployment" endpoint HTTP request body.
type InstallDeploymentRequestBody struct {
	// deploymentSetupConfig
	SetupConfig *DeploymentSetupConfigRequestBody `form:"setupConfig" json:"setupConfig" xml:"setupConfig"`
}

// UninstallDeploymentRequestBody is the type of the "installCtrlPanel" service
// "uninstallDeployment" endpoint HTTP request body.
type UninstallDeploymentRequestBody struct {
	// UnDeploymentSetupConfig
	SetupConfig *UnDeploymentSetupConfigRequestBody `form:"setupConfig" json:"setupConfig" xml:"setupConfig"`
}

// UpdateDeploymentRequestBody is the type of the "installCtrlPanel" service
// "updateDeployment" endpoint HTTP request body.
type UpdateDeploymentRequestBody struct {
	// updateDeploymentConfig
	SetupConfig *UpdateDeploymentConfigRequestBody `form:"setupConfig" json:"setupConfig" xml:"setupConfig"`
}

// ListInstallResponseBody is the type of the "installCtrlPanel" service
// "listInstall" endpoint HTTP response body.
type ListInstallResponseBody struct {
	Code *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// list of installed services
	Result  []*CtrlDeploayItemResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                        `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// InstallLpClientResponseBody is the type of the "installCtrlPanel" service
// "installLpClient" endpoint HTTP response body.
type InstallLpClientResponseBody struct {
	Code   *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result *struct {
		// rendered template content
		Template  *string `form:"Template" json:"Template" xml:"Template"`
		CmdStdout *string `form:"CmdStdout" json:"CmdStdout" xml:"CmdStdout"`
		CmdStderr *string `form:"CmdStderr" json:"CmdStderr" xml:"CmdStderr"`
	} `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UninstallLpClientResponseBody is the type of the "installCtrlPanel" service
// "uninstallLpClient" endpoint HTTP response body.
type UninstallLpClientResponseBody struct {
	Code   *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result *struct {
		// rendered template content
		Template  *string `form:"Template" json:"Template" xml:"Template"`
		CmdStdout *string `form:"CmdStdout" json:"CmdStdout" xml:"CmdStdout"`
		CmdStderr *string `form:"CmdStderr" json:"CmdStderr" xml:"CmdStderr"`
	} `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// InstallDeploymentResponseBody is the type of the "installCtrlPanel" service
// "installDeployment" endpoint HTTP response body.
type InstallDeploymentResponseBody struct {
	Code *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// install result
	Result  *InstallDeploymentDataResultResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                                  `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UninstallDeploymentResponseBody is the type of the "installCtrlPanel"
// service "uninstallDeployment" endpoint HTTP response body.
type UninstallDeploymentResponseBody struct {
	Code    *int64                                     `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result  *UnInstallDeploymentDataResultResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                                    `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateDeploymentResponseBody is the type of the "installCtrlPanel" service
// "updateDeployment" endpoint HTTP response body.
type UpdateDeploymentResponseBody struct {
	Code   *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result *struct {
		CmdStdout *string `form:"CmdStdout" json:"CmdStdout" xml:"CmdStdout"`
		CmdStderr *string `form:"CmdStderr" json:"CmdStderr" xml:"CmdStderr"`
		// rendered template content
		Template *string `form:"Template" json:"Template" xml:"Template"`
	} `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CtrlDeploayItemResponseBody is used to define fields on response body types.
type CtrlDeploayItemResponseBody struct {
	// install type
	InstallType *string `form:"installType,omitempty" json:"installType,omitempty" xml:"installType,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// install status
	Status *int64 `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// install context
	InstallContext *string `form:"installContext,omitempty" json:"installContext,omitempty" xml:"installContext,omitempty"`
	// yaml
	Yaml *string `form:"yaml,omitempty" json:"yaml,omitempty" xml:"yaml,omitempty"`
}

// AmmClientSetupConfigRequestBody is used to define fields on request body
// types.
type AmmClientSetupConfigRequestBody struct {
	CustomEnv []*DeploymentSetupConfigEnvItemRequestBody `form:"customEnv,omitempty" json:"customEnv,omitempty" xml:"customEnv,omitempty"`
	// imageRepository
	ImageRepository string `form:"imageRepository" json:"imageRepository" xml:"imageRepository"`
	// serviceName
	ServiceName *string `form:"serviceName,omitempty" json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// deploymentName
	DeploymentName *string `form:"deploymentName,omitempty" json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// type
	Type                  string  `form:"type" json:"type" xml:"type"`
	StartBlock            *string `form:"startBlock,omitempty" json:"startBlock,omitempty" xml:"startBlock,omitempty"`
	RPCURL                *string `form:"rpcUrl,omitempty" json:"rpcUrl,omitempty" xml:"rpcUrl,omitempty"`
	ConnectionNodeurl     *string `form:"connectionNodeurl,omitempty" json:"connectionNodeurl,omitempty" xml:"connectionNodeurl,omitempty"`
	ConnectionWalleturl   *string `form:"connectionWalleturl,omitempty" json:"connectionWalleturl,omitempty" xml:"connectionWalleturl,omitempty"`
	ConnectionHelperurl   *string `form:"connectionHelperurl,omitempty" json:"connectionHelperurl,omitempty" xml:"connectionHelperurl,omitempty"`
	ConnectionExplorerurl *string `form:"connectionExplorerurl,omitempty" json:"connectionExplorerurl,omitempty" xml:"connectionExplorerurl,omitempty"`
	AwsAccessKeyID        *string `form:"awsAccessKeyId,omitempty" json:"awsAccessKeyId,omitempty" xml:"awsAccessKeyId,omitempty"`
	ContainerPort         *string `form:"containerPort,omitempty" json:"containerPort,omitempty" xml:"containerPort,omitempty"`
	AwsSecretAccessKey    *string `form:"awsSecretAccessKey,omitempty" json:"awsSecretAccessKey,omitempty" xml:"awsSecretAccessKey,omitempty"`
	Install               bool    `form:"install" json:"install" xml:"install"`
}

// DeploymentSetupConfigEnvItemRequestBody is used to define fields on request
// body types.
type DeploymentSetupConfigEnvItemRequestBody struct {
	Key   *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// AmmClientUnSetupConfigRequestBody is used to define fields on request body
// types.
type AmmClientUnSetupConfigRequestBody struct {
	Type      *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	Uninstall *bool   `form:"uninstall,omitempty" json:"uninstall,omitempty" xml:"uninstall,omitempty"`
}

// DeploymentSetupConfigRequestBody is used to define fields on request body
// types.
type DeploymentSetupConfigRequestBody struct {
	ImageRepository string  `form:"imageRepository" json:"imageRepository" xml:"imageRepository"`
	ContainerPort   *string `form:"containerPort,omitempty" json:"containerPort,omitempty" xml:"containerPort,omitempty"`
	Install         bool    `form:"install" json:"install" xml:"install"`
	InstallType     string  `form:"installType" json:"installType" xml:"installType"`
	Name            string  `form:"name" json:"name" xml:"name"`
	// env list
	CustomEnv []*DeploymentSetupConfigEnvItemRequestBody `form:"customEnv,omitempty" json:"customEnv,omitempty" xml:"customEnv,omitempty"`
}

// InstallDeploymentDataResultResponseBody is used to define fields on response
// body types.
type InstallDeploymentDataResultResponseBody struct {
	CmdStdout *string `form:"CmdStdout,omitempty" json:"CmdStdout,omitempty" xml:"CmdStdout,omitempty"`
	CmdStderr *string `form:"CmdStderr,omitempty" json:"CmdStderr,omitempty" xml:"CmdStderr,omitempty"`
	Template  *string `form:"Template,omitempty" json:"Template,omitempty" xml:"Template,omitempty"`
}

// UnDeploymentSetupConfigRequestBody is used to define fields on request body
// types.
type UnDeploymentSetupConfigRequestBody struct {
	Uninstall   bool   `form:"uninstall" json:"uninstall" xml:"uninstall"`
	InstallType string `form:"installType" json:"installType" xml:"installType"`
	Name        string `form:"name" json:"name" xml:"name"`
}

// UnInstallDeploymentDataResultResponseBody is used to define fields on
// response body types.
type UnInstallDeploymentDataResultResponseBody struct {
	CmdStdout *string `form:"CmdStdout,omitempty" json:"CmdStdout,omitempty" xml:"CmdStdout,omitempty"`
	CmdStderr *string `form:"CmdStderr,omitempty" json:"CmdStderr,omitempty" xml:"CmdStderr,omitempty"`
	Template  *string `form:"Template,omitempty" json:"Template,omitempty" xml:"Template,omitempty"`
}

// UpdateDeploymentConfigRequestBody is used to define fields on request body
// types.
type UpdateDeploymentConfigRequestBody struct {
	InstallType    string  `form:"installType" json:"installType" xml:"installType"`
	Name           string  `form:"name" json:"name" xml:"name"`
	InstallContext *string `form:"installContext,omitempty" json:"installContext,omitempty" xml:"installContext,omitempty"`
	Update         bool    `form:"update" json:"update" xml:"update"`
}

// NewListInstallRequestBody builds the HTTP request body from the payload of
// the "listInstall" endpoint of the "installCtrlPanel" service.
func NewListInstallRequestBody(p *installctrlpanel.ListInstallPayload) *ListInstallRequestBody {
	body := &ListInstallRequestBody{
		InstallType: p.InstallType,
	}
	return body
}

// NewInstallLpClientRequestBody builds the HTTP request body from the payload
// of the "installLpClient" endpoint of the "installCtrlPanel" service.
func NewInstallLpClientRequestBody(p *installctrlpanel.InstallLpClientPayload) *InstallLpClientRequestBody {
	body := &InstallLpClientRequestBody{}
	if p.SetupConfig != nil {
		body.SetupConfig = marshalInstallctrlpanelAmmClientSetupConfigToAmmClientSetupConfigRequestBody(p.SetupConfig)
	}
	return body
}

// NewUninstallLpClientRequestBody builds the HTTP request body from the
// payload of the "uninstallLpClient" endpoint of the "installCtrlPanel"
// service.
func NewUninstallLpClientRequestBody(p *installctrlpanel.UninstallLpClientPayload) *UninstallLpClientRequestBody {
	body := &UninstallLpClientRequestBody{}
	if p.SetupConfig != nil {
		body.SetupConfig = marshalInstallctrlpanelAmmClientUnSetupConfigToAmmClientUnSetupConfigRequestBody(p.SetupConfig)
	}
	return body
}

// NewInstallDeploymentRequestBody builds the HTTP request body from the
// payload of the "installDeployment" endpoint of the "installCtrlPanel"
// service.
func NewInstallDeploymentRequestBody(p *installctrlpanel.InstallDeploymentPayload) *InstallDeploymentRequestBody {
	body := &InstallDeploymentRequestBody{}
	if p.SetupConfig != nil {
		body.SetupConfig = marshalInstallctrlpanelDeploymentSetupConfigToDeploymentSetupConfigRequestBody(p.SetupConfig)
	}
	return body
}

// NewUninstallDeploymentRequestBody builds the HTTP request body from the
// payload of the "uninstallDeployment" endpoint of the "installCtrlPanel"
// service.
func NewUninstallDeploymentRequestBody(p *installctrlpanel.UninstallDeploymentPayload) *UninstallDeploymentRequestBody {
	body := &UninstallDeploymentRequestBody{}
	if p.SetupConfig != nil {
		body.SetupConfig = marshalInstallctrlpanelUnDeploymentSetupConfigToUnDeploymentSetupConfigRequestBody(p.SetupConfig)
	}
	return body
}

// NewUpdateDeploymentRequestBody builds the HTTP request body from the payload
// of the "updateDeployment" endpoint of the "installCtrlPanel" service.
func NewUpdateDeploymentRequestBody(p *installctrlpanel.UpdateDeploymentPayload) *UpdateDeploymentRequestBody {
	body := &UpdateDeploymentRequestBody{}
	if p.SetupConfig != nil {
		body.SetupConfig = marshalInstallctrlpanelUpdateDeploymentConfigToUpdateDeploymentConfigRequestBody(p.SetupConfig)
	}
	return body
}

// NewListInstallResultOK builds a "installCtrlPanel" service "listInstall"
// endpoint result from a HTTP "OK" response.
func NewListInstallResultOK(body *ListInstallResponseBody) *installctrlpanel.ListInstallResult {
	v := &installctrlpanel.ListInstallResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = make([]*installctrlpanel.CtrlDeploayItem, len(body.Result))
		for i, val := range body.Result {
			v.Result[i] = unmarshalCtrlDeploayItemResponseBodyToInstallctrlpanelCtrlDeploayItem(val)
		}
	}

	return v
}

// NewInstallLpClientResultOK builds a "installCtrlPanel" service
// "installLpClient" endpoint result from a HTTP "OK" response.
func NewInstallLpClientResultOK(body *InstallLpClientResponseBody) *installctrlpanel.InstallLpClientResult {
	v := &installctrlpanel.InstallLpClientResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = &struct {
			// rendered template content
			Template  *string
			CmdStdout *string
			CmdStderr *string
		}{
			Template:  body.Result.Template,
			CmdStdout: body.Result.CmdStdout,
			CmdStderr: body.Result.CmdStderr,
		}
	}

	return v
}

// NewUninstallLpClientResultOK builds a "installCtrlPanel" service
// "uninstallLpClient" endpoint result from a HTTP "OK" response.
func NewUninstallLpClientResultOK(body *UninstallLpClientResponseBody) *installctrlpanel.UninstallLpClientResult {
	v := &installctrlpanel.UninstallLpClientResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = &struct {
			// rendered template content
			Template  *string
			CmdStdout *string
			CmdStderr *string
		}{
			Template:  body.Result.Template,
			CmdStdout: body.Result.CmdStdout,
			CmdStderr: body.Result.CmdStderr,
		}
	}

	return v
}

// NewInstallDeploymentResultOK builds a "installCtrlPanel" service
// "installDeployment" endpoint result from a HTTP "OK" response.
func NewInstallDeploymentResultOK(body *InstallDeploymentResponseBody) *installctrlpanel.InstallDeploymentResult {
	v := &installctrlpanel.InstallDeploymentResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = unmarshalInstallDeploymentDataResultResponseBodyToInstallctrlpanelInstallDeploymentDataResult(body.Result)
	}

	return v
}

// NewUninstallDeploymentResultOK builds a "installCtrlPanel" service
// "uninstallDeployment" endpoint result from a HTTP "OK" response.
func NewUninstallDeploymentResultOK(body *UninstallDeploymentResponseBody) *installctrlpanel.UninstallDeploymentResult {
	v := &installctrlpanel.UninstallDeploymentResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = unmarshalUnInstallDeploymentDataResultResponseBodyToInstallctrlpanelUnInstallDeploymentDataResult(body.Result)
	}

	return v
}

// NewUpdateDeploymentResultOK builds a "installCtrlPanel" service
// "updateDeployment" endpoint result from a HTTP "OK" response.
func NewUpdateDeploymentResultOK(body *UpdateDeploymentResponseBody) *installctrlpanel.UpdateDeploymentResult {
	v := &installctrlpanel.UpdateDeploymentResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = &struct {
			CmdStdout *string
			CmdStderr *string
			// rendered template content
			Template *string
		}{
			CmdStdout: body.Result.CmdStdout,
			CmdStderr: body.Result.CmdStderr,
			Template:  body.Result.Template,
		}
	}

	return v
}

// ValidateDeploymentSetupConfigRequestBody runs the validations defined on
// DeploymentSetupConfigRequestBody
func ValidateDeploymentSetupConfigRequestBody(body *DeploymentSetupConfigRequestBody) (err error) {
	if !(body.InstallType == "ammClient" || body.InstallType == "market" || body.InstallType == "amm" || body.InstallType == "userApp") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.installType", body.InstallType, []interface{}{"ammClient", "market", "amm", "userApp"}))
	}
	return
}
