package scrape

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ScrapeInventory() error {
	payload := []GQLRequest{
		{
			OperationName: "PersonalSections",
			Variables: VariablesPersonalSections{
				Input: Input{
					SectionInputs: []string{"RECS_FOLLOWED_SECTION", "RECOMMENDED_SECTION"},
					RecommendationContext: RecommendationContext{
						Platform:       "web",
						ClientApp:      "twilight",
						Location:       "inventory",
						ReferrerDomain: "www.twitch.tv",
						ViewportHeight: 969,
						ViewportWidth:  732,
					},
					WithFreeformTags: false,
				},
				CreatorAnniversariesFeature: false,
			},
			Extensions: Extensions{
				PersistedQuery: PersistedQuery{
					Version:    1,
					Sha256Hash: "25796b747fe97ab22412774edefcc4186b487f98ffa8d437e098c7e1968a6fa3",
				},
			},
		},
		{
			OperationName: "Inventory",
			Variables: VariablesInventory{
				FetchRewardCampaigns: true,
			},
			Extensions: Extensions{
				PersistedQuery: PersistedQuery{
					Version:    1,
					Sha256Hash: "09acb7d3d7e605a92bdfdcc465f6aa481b71c234d8686a9ba38ea5ed51507592",
				},
			},
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	twitch_cookies := []Cookies{
		{Name: "auth_token", Value: "k5n5honrk6xq5nyt5448y552nugtbu"},
		{Name: "persistent", Value: "715125887::nknvfwmmwmpmhd1yojtvmng6ceejux"},
	}

	req, err := http.NewRequest("POST", "https://gql.twitch.tv/gql", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req = AddCookies(req, twitch_cookies)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth yhu7mg9whbpp47bwfrvlwmj5e1es52")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	/*body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}
	fmt.Println(string(body))*/
	avalibleStreams, err := ScrapeRustDropsPage()
	if err != nil {
		return err
	}
	//fmt.Println(avalibleStreams)
	err = UnmarshallResponse(resp, avalibleStreams)
	if err != nil {
		return err
	}
	return nil
}

func UnmarshallResponse(resp *http.Response, avalibleStreams []string) error {
	var twitchDropsInventory []TwitchDropsInventory
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(b), &twitchDropsInventory)
	if err != nil {
		return err
	}
	dropsInProgress := twitchDropsInventory[1].Data.CurrentUser.Inventory.DropCampaignsInProgress
	var avalibleStreamsList []string

	var lastWord string
	var isOnce = true
	for _, dropInProgress := range dropsInProgress {
		//fmt.Println(dropInProgress.Name, dropInProgress.Status)
		if dropInProgress.Status != "EXPIRED" {
			timeBasedDrops := dropInProgress.TimeBasedDrops
			if !isOnce {
				break
			}
			for _, timeBasedDrop := range timeBasedDrops {
				if !isOnce {
					break
				}
				self := timeBasedDrop.Self
				if !self.IsClaimed {
					inventoryDropWords := strings.Fields(timeBasedDrop.Name)
					firstWord := inventoryDropWords[0]
					for _, avalibleStream := range avalibleStreams {
						avalibleStreamWords := strings.LastIndex(avalibleStream, "/")
						lastWord = avalibleStream[avalibleStreamWords+1:]
						//fmt.Println(avalibleStream)
						avalibleStreamsList = append(avalibleStreamsList, avalibleStream)
					}
					isOnce = false
					if firstWord == lastWord {

					}
				}
				fmt.Println(avalibleStreamsList)
				/*err := ScrapeStream(avalibleStreamsList[0])
				if err != nil {
					return err
					}*/
				//fmt.Printf("Name: %s\nRequired watch: %dmin\nWatched: %dmin\nIsCollected: %v", timeBasedDrop.Name, timeBasedDrop.RequiredMinutesWatched, self.CurrentMinutesWatched, self.IsClaimed)
			}
		}
	}
	return nil
}
