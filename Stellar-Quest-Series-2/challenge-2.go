package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main_c2() {
	client := horizonclient.DefaultTestNetClient

	// Remember, these are just examples, so replace them with your own seeds.
	issuerSeed := "MASTER_ACCOUNT_SECRET_KEY"
	distributorSeed := "QUEST_SERIES_2_SECRET_KEY"

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

	// Create an object to represent the new asset
	astroCheems := txnbuild.CreditAsset{Code: "AstroCheems", Issuer: issuer.Address()}

	// First, the receiving (distribution) account must trust the asset from the
	// issuer.
	op1 := txnbuild.ChangeTrust{
		Line: astroCheems,
		// Limit:         "454545",
		// SourceAccount: &distributorAccount,
	}

	// Second, the issuing account actually sends a payment using the asset
	op2 := txnbuild.Payment{
		Destination:   distributor.Address(),
		Asset:         astroCheems,
		Amount:        "4545",
		SourceAccount: &issuerAccount,
	}

	// Assemble Transaction
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &distributorAccount,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(),
			Operations:           []txnbuild.Operation{&op1, &op2},
		},
	)

	// check(err)
	if err != nil {
		fmt.Println("Error 2")
		fmt.Println(err)
		return
	}

	tx, err = tx.Sign(network.TestNetworkPassphrase, issuer, distributor)
	// check(err)
	if err != nil {
		fmt.Println("Error 3")
		fmt.Println(err)
		return
	}

	txe, err := tx.Base64()
	// check(err)
	if err != nil {
		fmt.Println("Error 4")
		fmt.Println(err)
		return
	}
	fmt.Println(txe)

	// signedTx, err = tx.Sign(network.TestNetworkPassphrase, issuer)
	// resp, err = client.SubmitTransaction(signedTx)

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Printf("Pay: %s\n", resp.Hash)
	// }
}
