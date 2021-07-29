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

	// Create Quest Account
	// Use the default pubnet client
	kp, _ := keypair.Parse("SECRET_KEY")
	client := horizonclient.DefaultTestNetClient
	ar := horizonclient.AccountRequest{AccountID: kp.Address()}
	sourceAccount, err := client.AccountDetail(ar)
	// check(err)
	if err != nil {
		fmt.Println("Error 1")
		log.Fatalln(err)
		return
	}

	op := txnbuild.CreateAccount{
		Destination: "GCL4WZZZ6ROA7NOZJIOJXYULVQIJDSAUDR2WPP6GM33BXWKUQAN4XBTV",
		Amount:      "100",
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

	// submit transaction
	resp, err := client.SubmitTransactionXDR(txe)
	if err != nil {
		fmt.Println("Failed to sent Transactions")
		fmt.Println(err)
		return
	}

	fmt.Println(resp)

	// Add Signer

	// // Use the default pubnet client
	// kp_quest, _ := keypair.Parse("SD6QHAZNUQHN44GJZN5674DCIN5CEJEPBFRQGFKYMBKNUT7PYCSUENE2")
	// client := horizonclient.DefaultTestNetClient
	// ar := horizonclient.AccountRequest{AccountID: kp_quest.Address()}
	// sourceAccount_quest, err := client.AccountDetail(ar)
	// // check(err)
	// if err != nil {
	// 	fmt.Println("Error 1")
	// 	log.Fatalln(err)
	// 	return
	// }

	// op := txnbuild.SetOptions{
	// 	// InflationDestination: NewInflationDestination("GCCOBXW2XQNUSL467IEILE6MMCNRR66SSVL4YQADUNYYNUVREF3FIV2Z"),
	// 	// ClearFlags:           []AccountFlag{AuthRevocable},
	// 	// SetFlags:             []AccountFlag{AuthRequired, AuthImmutable},
	// 	// MasterWeight:         NewThreshold(10),
	// 	// LowThreshold:         NewThreshold(1),
	// 	// MediumThreshold:      NewThreshold(2),
	// 	// HighThreshold:        NewThreshold(2),
	// 	// HomeDomain:           NewHomeDomain("LovelyLumensLookLuminous.com"),
	// 	Signer: &txnbuild.Signer{Address: "b42f743feaa0ecff7c9fb32c548ed56de849858b4eb5d17a64d6a90c89f60e5f", Weight: txnbuild.Threshold(1)},
	// }

	// tx, err := txnbuild.NewTransaction(
	// 	txnbuild.TransactionParams{
	// 		SourceAccount:        &sourceAccount_quest,
	// 		IncrementSequenceNum: true,
	// 		Operations:           []txnbuild.Operation{&op},
	// 		BaseFee:              txnbuild.MinBaseFee,
	// 		Timebounds:           txnbuild.NewInfiniteTimeout(), // Use a real timeout in production!
	// 	},
	// )
	// // check(err)
	// if err != nil {
	// 	fmt.Println("Error 2")
	// 	fmt.Println(err)
	// 	return
	// }

	// tx, err = tx.Sign(network.TestNetworkPassphrase, kp_quest.(*keypair.Full))
	// // check(err)
	// if err != nil {
	// 	fmt.Println("Error 3")
	// 	fmt.Println(err)
	// 	return
	// }

	// txe, err := tx.Base64()
	// // check(err)
	// if err != nil {
	// 	fmt.Println("Error 4")
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(txe)

	// // submit transaction
	// resp, err := client.SubmitTransactionXDR(txe)
	// if err != nil {
	// 	fmt.Println("Failed to sent Transactions")
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(resp)

	// // Remove Signer
}
