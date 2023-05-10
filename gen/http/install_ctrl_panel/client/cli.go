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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"installType\": \"Odio totam.\"\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"awsAccessKeyId\": \"Et alias eum nostrum.\",\n         \"awsSecretAccessKey\": \"Repellat voluptas consectetur vitae ut aliquam sed.\",\n         \"connectionExplorerurl\": \"Neque provident quam necessitatibus.\",\n         \"connectionHelperurl\": \"Natus doloremque a quae.\",\n         \"connectionNodeurl\": \"Ex sit et voluptatem.\",\n         \"connectionWalleturl\": \"Laboriosam eaque aut aut aut quas et.\",\n         \"containerPort\": \"Quam nostrum pariatur error.\",\n         \"customEnv\": [\n            {\n               \"key\": \"Qui ipsa et.\",\n               \"value\": \"Ad vero omnis aspernatur aspernatur aperiam non.\"\n            },\n            {\n               \"key\": \"Qui ipsa et.\",\n               \"value\": \"Ad vero omnis aspernatur aspernatur aperiam non.\"\n            },\n            {\n               \"key\": \"Qui ipsa et.\",\n               \"value\": \"Ad vero omnis aspernatur aspernatur aperiam non.\"\n            }\n         ],\n         \"deploymentName\": \"Id esse aut labore.\",\n         \"imageRepository\": \"Iste similique vitae quibusdam.\",\n         \"install\": true,\n         \"rpcUrl\": \"Occaecati non aut.\",\n         \"serviceName\": \"Cum rerum tempore fuga.\",\n         \"startBlock\": \"Ad inventore quisquam ut tempora perspiciatis.\",\n         \"type\": \"Provident consequatur laudantium reiciendis asperiores sequi.\"\n      }\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"type\": \"Est ipsa pariatur.\",\n         \"uninstall\": false\n      }\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"containerPort\": \"Et est ipsum illum qui dolores vel.\",\n         \"customEnv\": [\n            {\n               \"key\": \"Qui ipsa et.\",\n               \"value\": \"Ad vero omnis aspernatur aspernatur aperiam non.\"\n            },\n            {\n               \"key\": \"Qui ipsa et.\",\n               \"value\": \"Ad vero omnis aspernatur aspernatur aperiam non.\"\n            },\n            {\n               \"key\": \"Qui ipsa et.\",\n               \"value\": \"Ad vero omnis aspernatur aspernatur aperiam non.\"\n            }\n         ],\n         \"imageRepository\": \"Earum saepe et cumque.\",\n         \"install\": false,\n         \"installType\": \"amm\",\n         \"name\": \"Cumque veritatis pariatur alias.\"\n      }\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"installType\": \"Soluta repellat et quam qui et et.\",\n         \"name\": \"Sed quis eligendi eaque.\",\n         \"uninstall\": true\n      }\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"setupConfig\": {\n         \"installContext\": \"Et est laborum numquam et ut laudantium.\",\n         \"installType\": \"Aut qui ex et porro saepe.\",\n         \"name\": \"Et sed.\",\n         \"update\": false\n      }\n   }'")
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