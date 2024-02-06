package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

func main() {
	// Task 1: Generate redeem script
	preImage := "Btrust Builders"
	redeemScriptHex := generateRedeemScript(preImage)
	fmt.Println("Redeem Script (hex):", redeemScriptHex)

	// Task 2: Derive address from redeem script
	address := deriveAddress(redeemScriptHex)
	fmt.Println("Derived Address:", address)

	// Task 3: Construct transaction to send Bitcoins to the address
	txHex := constructTransaction(address)
	fmt.Println("Constructed Transaction (hex):", txHex)

	// Task 4: Construct transaction to spend from previous transaction
	spendingTxHex := constructSpendingTransaction(txHex, redeemScriptHex)
	fmt.Println("Constructed Spending Transaction (hex):", spendingTxHex)
}

// Task 1: Generate redeem script
func generateRedeemScript(preImage string) string {
	hash := sha256.Sum256([]byte(preImage))
	hashHex := hex.EncodeToString(hash[:])
	redeemScriptHex := "OP_SHA256 " + hashHex + " OP_EQUAL"
	return redeemScriptHex
}

// Task 2: Derive address from redeem script
func deriveAddress(redeemScriptHex string) string {
	redeemScript, err := hex.DecodeString(redeemScriptHex)
	if err != nil {
		fmt.Println("Error decoding redeem script hex:", err)
		return ""
	}

	p2sh, err := btcutil.NewAddressScriptHash(redeemScript, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Error deriving address from redeem script:", err)
		return ""
	}

	return p2sh.EncodeAddress()
}

// Task 3: Construct transaction to send Bitcoins to the address
func constructTransaction(address string) string {
	tx := wire.NewMsgTx(wire.TxVersion)
	tx.AddTxOut(wire.NewTxOut(100000, nil)) // Sending 0.001 BTC to the address

	// Serialize transaction
	var buf bytes.Buffer
	tx.Serialize(&buf)
	return hex.EncodeToString(buf.Bytes())
}

// Task 4: Construct transaction to spend from previous transaction
func constructSpendingTransaction(previousTxHex, redeemScriptHex string) string {
	previousTxBytes, err := hex.DecodeString(previousTxHex)
	if err != nil {
		fmt.Println("Error decoding previous transaction hex:", err)
		return ""
	}

	redeemScript, err := hex.DecodeString(redeemScriptHex)
	if err != nil {
		fmt.Println("Error decoding redeem script hex:", err)
		return ""
	}

	// Create input spending from previous transaction
	prevOut := wire.NewOutPoint(&chainhash.Hash{}, 0)
	txIn := wire.NewTxIn(prevOut, nil, nil)
	tx := wire.NewMsgTx(wire.TxVersion)
	tx.AddTxIn(txIn)

	// Add output to an address we control
	pkScript, err := txscript.PayToAddrScript(address)
	if err != nil {
		fmt.Println("Error creating output script:", err)
		return ""
	}
	txOut := wire.NewTxOut(90000, pkScript) // Sending 0.0009 BTC back to ourselves
	tx.AddTxOut(txOut)

	// Serialize transaction
	var buf bytes.Buffer
	tx.Serialize(&buf)
	return hex.EncodeToString(buf.Bytes())
}
