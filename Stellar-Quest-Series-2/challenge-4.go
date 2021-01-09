package main

import (
	"fmt"
	"time"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main_c4() {
	client := horizonclient.DefaultTestNetClient

	// Remember, these are just examples, so replace them with your own seeds.
	userSeed := "QUEST_SERIES_2_SECRET_KEY"

	/*
	 * We omit error checks here for brevity, but you should always check your
	 * return values.
	 */

	// Keys for accounts to issue and distribute the new asset.
	// feeBumpKP, err := keypair.ParseFull(feeBumpSeed)
	userKP, err := keypair.ParseFull(userSeed)
	request := horizonclient.AccountRequest{AccountID: userKP.Address()}
	userAccount, err := client.AccountDetail(request)

	// Create a claimable balance with our two above-described conditions.
	soon := time.Now().Add(time.Second * 60)
	// bCanClaim := txnbuild.BeforeRelativeTimePredicate(60)
	aCanReclaim := txnbuild.NotPredicate(
		txnbuild.BeforeAbsoluteTimePredicate(soon.Unix()),
	)

	claimants := []txnbuild.Claimant{
		// txnbuild.NewClaimant(B, &bCanClaim),
		txnbuild.NewClaimant(userKP.Address(), &aCanReclaim),
	}

	// Create the operation and submit it in a transaction.
	claimableBalanceEntry := txnbuild.CreateClaimableBalance{
		Destinations: claimants,
		Asset:        txnbuild.NativeAsset{},
		Amount:       "100",
	}

	// First, the receiving (distribution) account must trust the asset from the
	// issuer.
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &userAccount,
			IncrementSequenceNum: true,
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(),
			Operations:           []txnbuild.Operation{&claimableBalanceEntry},
		},
	)

	signedTx, err := tx.Sign(network.TestNetworkPassphrase, userKP)

	feeBumpTx, err := txnbuild.NewFeeBumpTransaction(
		txnbuild.FeeBumpTransactionParams{
			Inner:      signedTx,
			FeeAccount: userKP.Address(),
			BaseFee:    txnbuild.MinBaseFee,
		},
	)

	feeBumpTx, err = feeBumpTx.Sign(network.TestNetworkPassphrase, userKP)

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
