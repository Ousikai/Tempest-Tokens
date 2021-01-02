package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main() {
	client := horizonclient.DefaultTestNetClient

	// Remember, these are just examples, so replace them with your own seeds.
	issuerSeed := "MASTER_ACCOUNT_SECRET_KEY"
	distributorSeed := "QUEST_SERIES_1_SECRET_KEY"

	/*
	 * We omit error checks here for brevity, but you should always check your
	 * return values.
	 */

	// Keys for accounts to issue and distribute the new asset.
	issuer, err := keypair.ParseFull(issuerSeed)
	distributor, err := keypair.ParseFull(distributorSeed)

	request := horizonclient.AccountRequest{AccountID: issuer.Address()}
	issuerAccount, err := client.AccountDetail(request)

	request = horizonclient.AccountRequest{AccountID: distributor.Address()}
	distributorAccount, err := client.AccountDetail(request)

	// Create Operation
	op1 := txnbuild.BumpSequence{
		BumpTo: 0,
	}

	op2 := txnbuild.Payment{
		Destination:   issuer.Address(),
		Amount:        "10",
		Asset:         txnbuild.NativeAsset{},
		SourceAccount: &distributorAccount,
	}

	// First, the receiving (distribution) account must trust the asset from the
	// issuer.
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &issuerAccount,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(),
			Operations:           []txnbuild.Operation{&op1, &op2},
		},
	)

	signedTx, err := tx.Sign(network.TestNetworkPassphrase, issuer)

	txe, err := signedTx.Base64()
	// check(err)
	if err != nil {
		fmt.Println("Error 4")
		fmt.Println(err)
		return
	}
	fmt.Println(txe)
}
