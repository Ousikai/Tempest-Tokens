package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main_c7() {
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

	// Because we are sponsoring the accounts minimum balance, revoking the sponsorship would cause the account to be in violation
	// of the minimum balance requirement if it has a nil balance. Therefore, we need to send XLM to the account to cover this
	// fee requirement before revoking the sponsorship.
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sponsorAcc,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewTimeout(300),
			Operations: []txnbuild.Operation{
				&txnbuild.Payment{
					Asset:       txnbuild.NativeAsset{},
					Amount:      "1",
					Destination: newAccAddress,
				},
				&txnbuild.RevokeSponsorship{
					Account:         &newAccAddress,
					SponsorshipType: txnbuild.RevokeSponsorshipTypeAccount,
				},
			},
		},
	)

	signedTx, err := tx.Sign(network.TestNetworkPassphrase, sponsorKP)
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
