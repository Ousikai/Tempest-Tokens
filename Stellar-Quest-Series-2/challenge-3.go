package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main_c3() {
	client := horizonclient.DefaultTestNetClient

	// Remember, these are just examples, so replace them with your own seeds.
	feeBumpSeed := "MASTER_ACCOUNT_SECRET_KEY"
	userSeed := "QUEST_SERIES_2_SECRET_KEY"

	/*
	 * We omit error checks here for brevity, but you should always check your
	 * return values.
	 */

	// Keys for accounts to issue and distribute the new asset.
	feeBumpKP, err := keypair.ParseFull(feeBumpSeed)
	userKP, err := keypair.ParseFull(userSeed)

	request := horizonclient.AccountRequest{AccountID: feeBumpKP.Address()}
	// feeBumpAccount, err := client.AccountDetail(request)

	request = horizonclient.AccountRequest{AccountID: userKP.Address()}
	userAccount, err := client.AccountDetail(request)

	// Create Operation
	op := txnbuild.BumpSequence{
		BumpTo: 0,
	}

	// First, the receiving (distribution) account must trust the asset from the
	// issuer.
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &userAccount,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(),
			Operations:           []txnbuild.Operation{&op},
		},
	)

	signedTx, err := tx.Sign(network.TestNetworkPassphrase, userKP)

	feeBumpTx, err := txnbuild.NewFeeBumpTransaction(
		txnbuild.FeeBumpTransactionParams{
			Inner:      signedTx,
			FeeAccount: feeBumpKP.Address(),
			BaseFee:    txnbuild.MinBaseFee,
		},
	)

	feeBumpTx, err = feeBumpTx.Sign(network.TestNetworkPassphrase, feeBumpKP)

	txe, err := feeBumpTx.Base64()
	fmt.Println(txe)
	if err != nil {
		fmt.Println("Error 4")
		fmt.Println(err)
		return
	}
	fmt.Println(txe)

	// submit transaction
	resp, err := client.SubmitTransactionXDR(txe)
	if err != nil {
		fmt.Println("Failed to sent Transactions")
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
