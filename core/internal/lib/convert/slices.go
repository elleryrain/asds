package convert

func Slice[SI ~[]EI, SO ~[]EO, EI, EO any](in SI, fn func(EI) EO) SO {
	if in == nil {
		return nil
	}

	out := make(SO, 0, len(in))

	for _, item := range in {
		out = append(out, fn(item))
	}

	return out
}
