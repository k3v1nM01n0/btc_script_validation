package main

import (
    "bytes"
    "encoding/hex"
    "fmt"
    "github.com/btcsuite/btcd/wire"
)

func main() {
    // Raw transaction hex string
    rawTxHex := "020000000001010ccc140e766b5dbc884ea2d780c5e91e4eb77597ae64288a42575228b79e234900000000000000000002bd37060000000000225120245091249f4f29d30820e5f36e1e5d477dc3386144220bd6f35839e94de4b9cae81c00000000000016001416d31d7632aa17b3b316b813c0a3177f5b6150200140838a1f0f1ee607b54abf0a3f55792f6f8d09c3eb7a9fa46cd4976f2137ca2e3f4a901e314e1b827c3332d7e1865ffe1d7ff5f5d7576a9000f354487a09de44cd00000000"

    // Decode raw transaction hex
    rawTxBytes, err := hex.DecodeString(rawTxHex)
    if err != nil {
        fmt.Println("Error decoding raw transaction hex:", err)
        return
    }

    // Deserialize raw transaction bytes
    var tx wire.MsgTx
    err = tx.Deserialize(bytes.NewReader(rawTxBytes))
    if err != nil {
        fmt.Println("Error deserializing raw transaction:", err)
        return
    }

    // Print decoded transaction details
    fmt.Println("Version:", tx.Version)
    fmt.Println("Locktime:", tx.LockTime)
    fmt.Println("Number of inputs:", len(tx.TxIn))
    fmt.Println("Number of outputs:", len(tx.TxOut))
    for i, txIn := range tx.TxIn {
        fmt.Printf("Input %d: Previous output hash: %s, Index: %d\n", i, txIn.PreviousOutPoint.Hash.String(), txIn.PreviousOutPoint.Index)
    }
    for i, txOut := range tx.TxOut {
        fmt.Printf("Output %d: Value: %d, Script: %x\n", i, txOut.Value, txOut.PkScript)
    }
}

