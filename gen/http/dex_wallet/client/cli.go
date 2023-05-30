// Code generated by goa v3.11.0, DO NOT EDIT.
//
// dexWallet HTTP client CLI support package
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	dexwallet "admin-panel/gen/dex_wallet"
	"encoding/json"
	"fmt"

	goa "goa.design/goa/v3/pkg"
)

// BuildCreateDexWalletPayload builds the payload for the dexWallet
// createDexWallet endpoint from CLI flags.
func BuildCreateDexWalletPayload(dexWalletCreateDexWalletBody string) (*dexwallet.WalletRow, error) {
	var err error
	var body CreateDexWalletRequestBody
	{
		err = json.Unmarshal([]byte(dexWalletCreateDexWalletBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"accountId\": \"Dolores officiis est sed.\",\n      \"address\": \"Autem harum ullam ut dicta rerum eveniet.\",\n      \"chainId\": 7935885781834240208,\n      \"chainType\": \"Assumenda voluptas ullam culpa.\",\n      \"id\": \"Nisi non dolore quidem.\",\n      \"privateKey\": \"Accusamus maxime blanditiis cumque.\",\n      \"storeId\": \"Vitae repellendus rerum enim consectetur corporis.\",\n      \"vaultHostType\": \"At id quis neque ad dolorem.\",\n      \"vaultName\": \"Sed quidem eius numquam natus.\",\n      \"vaultSecertType\": \"Iste nihil explicabo quia.\",\n      \"walletName\": \"Voluptas doloremque quasi.\",\n      \"walletType\": \"storeId\"\n   }'")
		}
		if !(body.WalletType == "privateKey" || body.WalletType == "storeId") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.walletType", body.WalletType, []interface{}{"privateKey", "storeId"}))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &dexwallet.WalletRow{
		ID:              body.ID,
		WalletName:      body.WalletName,
		PrivateKey:      body.PrivateKey,
		Address:         body.Address,
		ChainType:       body.ChainType,
		AccountID:       body.AccountID,
		ChainID:         body.ChainID,
		StoreID:         body.StoreID,
		VaultHostType:   body.VaultHostType,
		VaultName:       body.VaultName,
		VaultSecertType: body.VaultSecertType,
		WalletType:      body.WalletType,
	}

	return v, nil
}

// BuildDeleteDexWalletPayload builds the payload for the dexWallet
// deleteDexWallet endpoint from CLI flags.
func BuildDeleteDexWalletPayload(dexWalletDeleteDexWalletBody string) (*dexwallet.DeleteFilter, error) {
	var err error
	var body DeleteDexWalletRequestBody
	{
		err = json.Unmarshal([]byte(dexWalletDeleteDexWalletBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Provident non maxime.\"\n   }'")
		}
	}
	v := &dexwallet.DeleteFilter{
		ID: body.ID,
	}

	return v, nil
}
