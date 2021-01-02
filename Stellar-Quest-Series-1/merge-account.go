package main

import (
	"fmt"
	"log"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

// func check(err string) {
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
// }

func main() {
	kp, _ := keypair.Parse("RANDOM_ACCOUNT_SECRET_KEY")
	client := horizonclient.DefaultTestNetClient
	ar := horizonclient.AccountRequest{AccountID: kp.Address()}
	sourceAccount, err := client.AccountDetail(ar)
	// fmt.Println(sourceAccount)
	if err != nil {
		fmt.Println("Error 1")
		log.Fatalln(err)
		return
	}

	op := txnbuild.AccountMerge{
		Destination:   "QUEST_SERIES_1_PUBLIC_KEY",
		SourceAccount: &sourceAccount,
	}

	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sourceAccount,
			IncrementSequenceNum: true,
			Operations:           []txnbuild.Operation{&op},
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(), // Use a real timeout in production!
		},
	)
	// check(err)
	if err != nil {
		fmt.Println("Error 2")
		fmt.Println(err)
		return
	}
	tx, err = tx.Sign(network.TestNetworkPassphrase, kp.(*keypair.Full))
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
}
