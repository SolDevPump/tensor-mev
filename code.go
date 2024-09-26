package main

import (
)

func main() {
	logger := logger.Logger{}

	cfg := constants.ReadJson()
	spew.Dump(cfg)

	tensor, _ := constants.NewTradeSettings(
		cfg.WalletKey,
		cfg.MintAddress,
		cfg.RPCConnection,
		cfg.MinPrice*float64(solana.LAMPORTS_PER_SOL),
		cfg.MaxPrice*float64(solana.LAMPORTS_PER_SOL),
		cfg.CycleDelay,
		cfg.Iteration_delay,
		cfg.BundleTipSize,
	)

	jitoClient, err := searcher_client.NewNoAuth(
		context.TODO(),
		jito_go.NewYork.BlockEngineURL,
		tensor.RPCConnection,
		tensor.RPCConnection,
		nil,
	)
	if err != nil {
		log.Fatalf("Error generating jito client: %s\n", err.Error())
	}

	for {
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
		processTransaction(context.TODO(), logger, tensorApi, tensor, jitoClient)
		logger.Info("All transactions processed.")
		time.Sleep(time.Duration(tensor.CycleDelay * float64(time.Second)))
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
	}
}

func processTransaction(ctx context.Context, logger logger.Logger, tensorApi *api.NftService, tensor *constants.TradeSettings, jitoClient *searcher_client.Client) error {

	// Generate transaction arrays

	// Fetch recent blockhash
	logger.Info("Fetching blockhash to pass in jito bundle")
	blockhash, err := tensor.RPCConnection.GetRecentBlockhash(ctx, rpc.CommitmentProcessed)
	if err != nil {
		logger.Error(fmt.Sprintf("Error fetching blockhash: %s", err.Error()))
		blockhash = nil // Ensure blockhash is nil if fetch fails
	}

}

// Refactored processTransactions as a separate function
func processTransactions(ctx context.Context, transactions []*solana.Transaction, walletPrivateKey []byte, rpcConnection *rpc.Client, logger logger.Logger) time.Time {

}

func generateTransactionArray(ctx context.Context, logger logger.Logger, tensorApi *api.NftService, tensor *constants.TradeSettings, listPrice float64, txAmount int) []*solana.Transaction {

}
