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
	distributorSeed := "QUEST_SERIES_1_SECRET_KEY"

	// Keys for accounts to issue and distribute the new asset.
	distributor, err := keypair.ParseFull(distributorSeed)

	request := horizonclient.AccountRequest{AccountID: distributor.Address()}
	distributorAccount, err := client.AccountDetail(request)

	// Create an object to represent the new asset
	srt := txnbuild.CreditAsset{Code: "SRT", Issuer: "GCDNJUBQSX7AJWLJACMJ7I4BC3Z47BQUTMHEICZLE6MU4KQBRYG5JY6B"}

	// First, the receiving (distribution) account must trust the asset from the source
	op1 := txnbuild.ChangeTrust{
		Line:  srt,
		Limit: "1000",
	}

	// Second, the issuing account actually sends a payment using the asset
	op2 := txnbuild.PathPaymentStrictSend{
		SendAsset:     txnbuild.NativeAsset{},
		SendAmount:    "500",
		Destination:   distributor.Address(),
		DestAsset:     srt,
		DestMin:       "1",
		Path:          []txnbuild.Asset{srt},
		SourceAccount: &distributorAccount,
	}

	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &distributorAccount,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(),
			Operations:           []txnbuild.Operation{&op1, &op2},
		},
	)

	tx, err = tx.Sign(network.TestNetworkPassphrase, distributor)
	txe, err := tx.Base64()
	if err != nil {
		fmt.Println("Error 4")
		fmt.Println(err)
		return
	}
	fmt.Println(txe)
}
