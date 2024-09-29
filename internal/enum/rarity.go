package enum

type Rarity string

const (
	RarityN    Rarity = "N"
	RarityR    Rarity = "R"
	RaritySR   Rarity = "SR"
	RaritySSR  Rarity = "SSR"
	RarityASSR Rarity = "ASSR"
)

func (r Rarity) String() string {
	return string(r)
}

func (r Rarity) isValid() bool {
	switch r {
	case RarityN, RarityR, RaritySR, RaritySSR, RarityASSR:
		return true
	default:
		return false
	}
}
