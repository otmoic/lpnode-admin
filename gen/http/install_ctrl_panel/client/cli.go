// Code generated by goa v3.11.0, DO NOT EDIT.
//
// installCtrlPanel HTTP client CLI support package
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	installctrlpanel "admin-panel/gen/install_ctrl_panel"
	"encoding/json"
	"fmt"

	goa "goa.design/goa/v3/pkg"
)

// BuildListInstallPayload builds the payload for the installCtrlPanel
// listInstall endpoint from CLI flags.
func BuildListInstallPayload(installCtrlPanelListInstallBody string) (*installctrlpanel.ListInstallPayload, error) {
	var err error
	var body ListInstallRequestBody
	{
		err = json.Unmarshal([]byte(installCtrlPanelListInstallBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"installType\": \"In numquam et.\"\n   }'")
		}
	}
	v := &installctrlpanel.ListInstallPayload{
		InstallType: body.InstallType,
	}

	return v, nil
}

// BuildInstallLpClientPayload builds the payload for the installCtrlPanel
// installLpClient endpoint from CLI flags.
func BuildInstallLpClientPayload(installCtrlPanelInstallLpClientBody string) (*installctrlpanel.InstallLpClientPayload, error) {
	var err error
	var body InstallLpClientRequestBody
	{
		err = json.Unmarshal([]byte(installCtrlPanelInstallLpClientBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"awsAccessKeyId\": \"Vel dolores ullam incidunt labore rem quibusdam.\",\n         \"awsSecretAccessKey\": \"In ratione labore molestiae.\",\n         \"connectionExplorerurl\": \"Aut rerum repellendus.\",\n         \"connectionHelperurl\": \"Aut ut rerum praesentium omnis.\",\n         \"connectionNodeurl\": \"Eaque et ex.\",\n         \"connectionWalleturl\": \"Quo et.\",\n         \"containerPort\": \"Rerum illum recusandae.\",\n         \"customEnv\": [\n            {\n               \"key\": \"Fugit quia qui quisquam illum non.\",\n               \"value\": \"Quaerat aut.\"\n            },\n            {\n               \"key\": \"Fugit quia qui quisquam illum non.\",\n               \"value\": \"Quaerat aut.\"\n            }\n         ],\n         \"deploymentName\": \"Iure impedit nesciunt ut rerum sed.\",\n         \"imageRepository\": \"Et est facilis a nulla ipsa ratione.\",\n         \"install\": false,\n         \"rpcUrl\": \"Voluptas officiis sed voluptates.\",\n         \"serviceName\": \"Aut voluptatem repellat.\",\n         \"startBlock\": \"Voluptatem dolorem.\",\n         \"type\": \"Et iusto voluptatem debitis.\"\n      }\n   }'")
		}
		if body.SetupConfig == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("setupConfig", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &installctrlpanel.InstallLpClientPayload{}
	if body.SetupConfig != nil {
		v.SetupConfig = marshalAmmClientSetupConfigRequestBodyToInstallctrlpanelAmmClientSetupConfig(body.SetupConfig)
	}

	return v, nil
}

// BuildUninstallLpClientPayload builds the payload for the installCtrlPanel
// uninstallLpClient endpoint from CLI flags.
func BuildUninstallLpClientPayload(installCtrlPanelUninstallLpClientBody string) (*installctrlpanel.UninstallLpClientPayload, error) {
	var err error
	var body UninstallLpClientRequestBody
	{
		err = json.Unmarshal([]byte(installCtrlPanelUninstallLpClientBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"type\": \"Magnam sed.\",\n         \"uninstall\": false\n      }\n   }'")
		}
		if body.SetupConfig == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("setupConfig", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &installctrlpanel.UninstallLpClientPayload{}
	if body.SetupConfig != nil {
		v.SetupConfig = marshalAmmClientUnSetupConfigRequestBodyToInstallctrlpanelAmmClientUnSetupConfig(body.SetupConfig)
	}

	return v, nil
}

// BuildInstallDeploymentPayload builds the payload for the installCtrlPanel
// installDeployment endpoint from CLI flags.
func BuildInstallDeploymentPayload(installCtrlPanelInstallDeploymentBody string) (*installctrlpanel.InstallDeploymentPayload, error) {
	var err error
	var body InstallDeploymentRequestBody
	{
		err = json.Unmarshal([]byte(installCtrlPanelInstallDeploymentBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"containerPort\": \"Quia esse fugit vel laboriosam tenetur.\",\n         \"customEnv\": [\n            {\n               \"key\": \"Fugit quia qui quisquam illum non.\",\n               \"value\": \"Quaerat aut.\"\n            },\n            {\n               \"key\": \"Fugit quia qui quisquam illum non.\",\n               \"value\": \"Quaerat aut.\"\n            }\n         ],\n         \"imageRepository\": \"Ut veniam sint.\",\n         \"install\": false,\n         \"installType\": \"ammClient\",\n         \"name\": \"Sunt quam sunt aliquam.\"\n      }\n   }'")
		}
		if body.SetupConfig == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("setupConfig", "body"))
		}
		if body.SetupConfig != nil {
			if err2 := ValidateDeploymentSetupConfigRequestBody(body.SetupConfig); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	v := &installctrlpanel.InstallDeploymentPayload{}
	if body.SetupConfig != nil {
		v.SetupConfig = marshalDeploymentSetupConfigRequestBodyToInstallctrlpanelDeploymentSetupConfig(body.SetupConfig)
	}

	return v, nil
}

// BuildUninstallDeploymentPayload builds the payload for the installCtrlPanel
// uninstallDeployment endpoint from CLI flags.
func BuildUninstallDeploymentPayload(installCtrlPanelUninstallDeploymentBody string) (*installctrlpanel.UninstallDeploymentPayload, error) {
	var err error
	var body UninstallDeploymentRequestBody
	{
		err = json.Unmarshal([]byte(installCtrlPanelUninstallDeploymentBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"installType\": \"Vel itaque sint.\",\n         \"name\": \"Et asperiores.\",\n         \"uninstall\": false\n      }\n   }'")
		}
		if body.SetupConfig == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("setupConfig", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &installctrlpanel.UninstallDeploymentPayload{}
	if body.SetupConfig != nil {
		v.SetupConfig = marshalUnDeploymentSetupConfigRequestBodyToInstallctrlpanelUnDeploymentSetupConfig(body.SetupConfig)
	}

	return v, nil
}

// BuildUpdateDeploymentPayload builds the payload for the installCtrlPanel
// updateDeployment endpoint from CLI flags.
func BuildUpdateDeploymentPayload(installCtrlPanelUpdateDeploymentBody string) (*installctrlpanel.UpdateDeploymentPayload, error) {
	var err error
	var body UpdateDeploymentRequestBody
	{
		err = json.Unmarshal([]byte(installCtrlPanelUpdateDeploymentBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"installContext\": \"Aliquid quibusdam deserunt aut.\",\n         \"installType\": \"Commodi quos itaque et quasi est.\",\n         \"name\": \"Est quam vel modi.\",\n         \"update\": true\n      }\n   }'")
		}
		if body.SetupConfig == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("setupConfig", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &installctrlpanel.UpdateDeploymentPayload{}
	if body.SetupConfig != nil {
		v.SetupConfig = marshalUpdateDeploymentConfigRequestBodyToInstallctrlpanelUpdateDeploymentConfig(body.SetupConfig)
	}

	return v, nil
}
