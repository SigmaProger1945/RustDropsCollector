package scrape

type PersistedQuery struct {
	Version    int    `json:"version"`
	Sha256Hash string `json:"sha256Hash"`
}

type Extensions struct {
	PersistedQuery PersistedQuery `json:"persistedQuery"`
}

type RecommendationContext struct {
	Platform       string `json:"platform"`
	ClientApp      string `json:"clientApp"`
	Location       string `json:"location"`
	ReferrerDomain string `json:"referrerDomain"`
	ViewportHeight int    `json:"viewportHeight"`
	ViewportWidth  int    `json:"viewportWidth"`
}

type Input struct {
	SectionInputs         []string              `json:"sectionInputs"`
	RecommendationContext RecommendationContext `json:"recommendationContext"`
	WithFreeformTags      bool                  `json:"withFreeformTags"`
}

type VariablesPersonalSections struct {
	Input                       Input `json:"input"`
	CreatorAnniversariesFeature bool  `json:"creatorAnniversariesFeature"`
}

type VariablesInventory struct {
	FetchRewardCampaigns bool `json:"fetchRewardCampaigns"`
}

type GQLRequest struct {
	OperationName string      `json:"operationName"`
	Variables     interface{} `json:"variables"`
	Extensions    Extensions  `json:"extensions"`
}
