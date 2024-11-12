package scrape

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	clientID     = "your-client-id"
	clientSecret = "your-client-secret"
)

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func getClientCredentialsToken() (string, error) {
	// Prepare data for request
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")

	// Create a POST request to the Twitch OAuth token endpoint
	req, err := http.NewRequest("POST", "https://id.twitch.tv/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var oauthResponse OAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&oauthResponse); err != nil {
		return "", err
	}

	return oauthResponse.AccessToken, nil
}

/*func main() {
	token, err := getClientCredentialsToken()
	if err != nil {
		fmt.Println("Error obtaining token:", err)
		return
	}

	fmt.Println("OAuth Token:", token)
	// Use this token in API requests as needed
}
*/
