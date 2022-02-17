package jaxvanity

type worker struct {
	gen IGenerator
}

// Work generates bitcoin wallet and pushes through channel
func (w *worker) Work(compressed bool) (result IWallet, erri error) {
	wallet, err := w.gen.Generate(compressed)
	if err != nil {
		erri = err
		return
	}
	result = wallet
	return
}
