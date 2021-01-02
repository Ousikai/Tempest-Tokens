package main

import (
	"log"

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

	// Create an object to represent the new asset
	astroCheems := txnbuild.CreditAsset{Code: "AstroCheems", Issuer: issuer.Address()}

	// First, the receiving (distribution) account must trust the asset from the
	// issuer.
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &distributorAccount,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(),
			Operations: []txnbuild.Operation{
				&txnbuild.ChangeTrust{
					Line:  astroCheems,
					Limit: "454545",
				},
			},
		},
	)

	signedTx, err := tx.Sign(network.TestNetworkPassphrase, distributor)
	resp, err := client.SubmitTransaction(signedTx)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Trust: %s\n", resp.Hash)
	}

	// Second, the issuing account actually sends a payment using the asset
	tx, err = txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &issuerAccount,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(),
			Operations: []txnbuild.Operation{
				&txnbuild.Payment{
					Destination: distributor.Address(),
					Asset:       astroCheems,
					Amount:      "4545",
				},
			},
		},
	)

	signedTx, err = tx.Sign(network.TestNetworkPassphrase, issuer)
	resp, err = client.SubmitTransaction(signedTx)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Pay: %s\n", resp.Hash)
	}
}
