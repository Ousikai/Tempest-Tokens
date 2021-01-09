package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main_c6() {
	client := horizonclient.DefaultTestNetClient

	// Remember, these are just examples, so replace them with your own seeds.
	newAccountSeed := "NEW_ACCOUNT_SECRET_KEY"
	sponsorSeed := "QUEST_SERIES_2_SECRET_KEY"

	/*
	 * We omit error checks here for brevity, but you should always check your
	 * return values.
	 */

	// Keys for accounts to issue and distribute the new asset.
	newAccKP, err := keypair.ParseFull(newAccountSeed)
	newAccAddress := newAccKP.Address()
	sponsorKP, err := keypair.ParseFull(sponsorSeed)

	request := horizonclient.AccountRequest{AccountID: newAccKP.Address()}
	// newAcc, err := client.AccountDetail(request)

	request = horizonclient.AccountRequest{AccountID: sponsorKP.Address()}
	sponsorAcc, err := client.AccountDetail(request)

	// Sponser will create an new account
	sponsorNewAccCreation := []txnbuild.Operation{
		&txnbuild.BeginSponsoringFutureReserves{
			SourceAccount: &sponsorAcc,
			SponsoredID:   newAccAddress,
		},
		&txnbuild.CreateAccount{
			Destination:   newAccAddress,
			Amount:        "0",
			SourceAccount: &sponsorAcc,
		},
		&txnbuild.EndSponsoringFutureReserves{
			SourceAccount: &txnbuild.SimpleAccount{
				AccountID: newAccKP.Address(),
			},
		},
	}
	// First, the receiving (distribution) account must trust the asset from the
	// issuer.
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sponsorAcc,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(),
			Operations:           sponsorNewAccCreation,
		},
	)

	signedTx, err := tx.Sign(network.TestNetworkPassphrase, sponsorKP, newAccKP)
	txe, err := signedTx.Base64()
	fmt.Println(txe)
	if err != nil {
		fmt.Println("Error 4")
		fmt.Println(err)
		return
	}
	// submit transaction
	// resp, err := client.SubmitTransactionXDR(txe)
	// if err != nil {
	// 	fmt.Println("Failed to sent Transactions")
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(resp)
}
