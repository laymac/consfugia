	query := datastore.NewQuery("Wallet").Filter("Address =", address_wallet)
	var wallet Wallet
	if err := client.Get(ctx, query.KeysOnly(), &wallet); err != nil {
		return err
	}  
