package lengthconv

func KMToM(km KM) M { return M(km * 1000) }

func MToDM(m M) DM { return DM(m * 10) }

func DMToKM(dm DM) KM { return KM(dm / 10000) }
