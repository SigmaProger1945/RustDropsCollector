package scrape

type TwitchDropsInventory struct {
	Data TwitchDropsInventoryData `json:"data"`
}

type TwitchDropsInventoryData struct {
	CurrentUser TwitchDropsInventoryCurrentUser `json:"currentUser"`
}

type TwitchDropsInventoryCurrentUser struct {
	Inventory TwitchDropsInventoryCurrentUserInventory `json:"inventory"`
}

type TwitchDropsInventoryCurrentUserInventory struct {
	DropCampaignsInProgress []TwitchDropsInventoryDropsInProgress `json:"dropCampaignsInProgress"`
}

type TwitchDropsInventoryDropsInProgress struct {
	Name           string                                  `json:"name"`
	Status         string                                  `json:"status"`
	Game           TwitchDropsInventoryDropsInProgressGame `json:"game"`
	TimeBasedDrops []TwitchDropsInventoryTimeBasedDrops    `json:"timeBasedDrops"`
}

type TwitchDropsInventoryTimeBasedDrops struct {
	Name                   string                                  `json:"name"`
	RequiredMinutesWatched int                                     `json:"requiredMinutesWatched"`
	Self                   TwitchDropsInventoryDropsInProgressSelf `json:"self"`
}

type TwitchDropsInventoryDropsInProgressSelf struct {
	CurrentMinutesWatched int  `json:"currentMinutesWatched"`
	IsClaimed             bool `json:"isClaimed"`
}

type TwitchDropsInventoryDropsInProgressGame struct {
	Name string `json:"name"`
}
