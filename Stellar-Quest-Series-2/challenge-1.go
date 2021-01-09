package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main_c1() {

	// Use the default pubnet client
	kp, _ := keypair.Parse("MASTER_ACCOUNT_SECRET_KEY")
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
		Destination: "SERIES_2_PUBLIC_KEY",
		Amount:      "5000",
	}

	var txnMemo [32]byte
	hexBytes, err := hex.DecodeString("e3366fcb087bdb2381b7069a19405b748da831c18145eba25654d1092e93ef37")
	if err != nil {
		log.Fatalln(err)
	}
	copy(txnMemo[:], hexBytes)

	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sourceAccount,
			IncrementSequenceNum: true,
			Operations:           []txnbuild.Operation{&op},
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(), // Use a real timeout in production!
			Memo:                 txnbuild.MemoHash(txnMemo),
			// Memo:                 txnbuild.MemoText(str),
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
	// resp, err := client.SubmitTransactionXDR(txe)
	// if err != nil {
	// 	fmt.Println("Failed to sent Transactions")
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(resp)
}
