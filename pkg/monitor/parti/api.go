package parti

import (
	"encoding/json"
	"net/http"
)

// 465731

func CheckPublicApi(userId string) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api-backend.parti.com/parti_v2/profile/get_livestream_channel_info/"+userId, nil)
	if err != nil {
		return false, err
	}

	req.Header = http.Header{
		"accept":             []string{"application/json"},
		"accept-language":    []string{"en-US,en;q=0.9"},
		"cache-control":      []string{"no-cache"},
		"pragma":             []string{"no-cache"},
		"sec-ch-ua":          []string{"\"Chromium\";v=\"124\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"124\""},
		"sec-ch-ua-mobile":   []string{"?0"},
		"sec-ch-ua-platform": []string{"\"Windows\""},
		"sec-fetch-dest":     []string{"empty"},
		"sec-fetch-mode":     []string{"cors"},
		"sec-fetch-site":     []string{"same-origin"},
		"user-agent":         []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, err
	}

	var response ApiResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return false, err
	}

	return response.IsStreamingLiveNow, nil
}
