package main

import (
	"fmt"
	"log"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main() {

	// Series 1
	kp, _ := keypair.Parse("QUEST_SERIES_SECRET_KEY_HERE")
	client := horizonclient.DefaultTestNetClient
	ar := horizonclient.AccountRequest{AccountID: kp.Address()}
	s1, err := client.AccountDetail(ar)
	// fmt.Println(sourceAccount)
	if err != nil {
		fmt.Println("Error 1: s1")
		log.Fatalln(err)
		return
	}

	// Tempest King
	kp_t1, _ := keypair.Parse("MASTER_ACCOUNT_SECRET_KEY_HERE")
	// client_t1 := horizonclient.DefaultTestNetClient
	ar_t1 := horizonclient.AccountRequest{AccountID: kp.Address()}
	t1, err := client.AccountDetail(ar_t1)
	// fmt.Println(sourceAccount)
	if err != nil {
		fmt.Println("Error 1: t1")
		log.Fatalln(err)
		return
	}

	op := txnbuild.Payment{
		Destination:   "MASTER_ACCOUNT_PUBLIC_KEY_HERE",
		Amount:        "100",
		Asset:         txnbuild.NativeAsset{},
		SourceAccount: &s1,
	}

	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &t1,
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
	tx, err = tx.Sign(network.TestNetworkPassphrase, kp_t1.(*keypair.Full))
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
