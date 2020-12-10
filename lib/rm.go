package lib

func (h *Handler) CalcOneRm() error {
	h.OneRm = h.Kg * (1 + (float64(h.Rep) / float64(h.Correction)))

	return nil
}

func (h *Handler) CalcRm() (float64, error) {
	kg := h.OneRm / (1 + (float64(h.TargetRm) / float64(40)))

	return kg, nil
}
