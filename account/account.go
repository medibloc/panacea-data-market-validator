package account

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-data-market-validator/crypto"
	tmcrypto "github.com/tendermint/tendermint/crypto"
)

type ValidatorAccount struct {
	privKey tmcrypto.PrivKey
	pubKey  tmcrypto.PubKey
}

func NewValidatorAccount(mnemonic string) (ValidatorAccount, error) {
	privKey, err := crypto.GeneratePrivateKeyFromMnemonic(mnemonic)

	if err != nil {
		return ValidatorAccount{}, err
	}

	return ValidatorAccount{
		privKey: privKey,
		pubKey:  privKey.PubKey(),
	}, nil
}

func (v ValidatorAccount) GetAddress() string {
	return sdk.AccAddress(v.pubKey.Address().Bytes()).String()
}

func (v ValidatorAccount) GetPrivKey() tmcrypto.PrivKey {
	return v.privKey
}

func (v ValidatorAccount) GetPubKey() tmcrypto.PubKey {
	return v.pubKey
}