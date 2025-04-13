package dadata

const (
	baseURL               = "https://suggestions.dadata.ru"
	findByIdPartyEndpoint = "/suggestions/api/4_1/rs/findById/party"
)

type request struct {
	Query string `json:"query"`
}

type FindByInnResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	Data SuggestionData `json:"data"`
}

type SuggestionData struct {
	Name    SuggestionDataName    `json:"name"`
	OKPO    string                `json:"okpo"`
	OGRN    string                `json:"ogrn"`
	KPP     string                `json:"kpp"`
	INN     string                `json:"inn"`
	Address SuggestionDataAddress `json:"address"`
}

type SuggestionDataName struct {
	Short        string `json:"short"`
	ShortWithOpf string `json:"short_with_opf"`
	FullWithOpf  string `json:"full_with_opf"`
}

type SuggestionDataAddress struct {
	UnrestrictedValue string                    `json:"unrestricted_value"`
	Data              SuggestionDataAddressData `json:"data"`
}

type SuggestionDataAddressData struct {
	TaxOffice string `json:"tax_office"`
}
