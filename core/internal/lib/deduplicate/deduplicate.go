package deduplicate

type Versioned interface {
	GetID() int
	GetVersion() int
}

func DeduplicateByVersion[T Versioned](items []T) []T {
	maxVersionMap := make(map[int]T)

	for _, item := range items {
		id := item.GetID()
		if existing, found := maxVersionMap[id]; found {
			if item.GetVersion() > existing.GetVersion() {
				maxVersionMap[id] = item
			}
		} else {
			maxVersionMap[id] = item
		}
	}

	deduplicated := make([]T, 0, len(maxVersionMap))
	for _, item := range maxVersionMap {
		deduplicated = append(deduplicated, item)
	}

	return deduplicated
}

func Deduplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]struct{})
	deduplicated := make([]T, 0)

	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = struct{}{}

			deduplicated = append(deduplicated, item)
		}
	}

	return deduplicated
}
