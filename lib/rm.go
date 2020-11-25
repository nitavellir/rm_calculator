package lib

func (h *Handler) CalcRm() error {
	if h.HardMode {
		h.OneRm = h.Kg * (1 + (float64(h.Rep) / float64(30)))
	} else {
		h.OneRm = h.Kg * (1 + (float64(h.Rep) / float64(40)))
	}
	return nil
}
