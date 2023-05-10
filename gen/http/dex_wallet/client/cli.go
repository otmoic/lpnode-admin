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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"accountId\": \"Et facilis deleniti ea sit praesentium.\",\n      \"address\": \"Consequatur dolorum.\",\n      \"chainId\": 7070136605965844913,\n      \"chainType\": \"Qui est sint maiores minima ex.\",\n      \"id\": \"Sint vel mollitia.\",\n      \"privateKey\": \"Deleniti cumque eum velit vero.\",\n      \"storeId\": \"Velit odit nobis in pariatur eaque.\",\n      \"vaultHostType\": \"Quidem officiis ipsam distinctio ducimus ipsum.\",\n      \"vaultName\": \"Nisi non dolore quidem.\",\n      \"vaultSecertType\": \"Voluptas doloremque quasi.\",\n      \"walletName\": \"Dolorem quod numquam.\",\n      \"walletType\": \"storeId\"\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Assumenda voluptas ullam culpa.\"\n   }'")
		}
	}
	v := &dexwallet.DeleteFilter{
		ID: body.ID,
	}

	return v, nil
}