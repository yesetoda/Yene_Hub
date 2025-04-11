package rate_limit_services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func IsRateLimited(ip string, limit int, windowSeconds int, upstashURL, upstashToken string) (bool, error) {
	key := fmt.Sprintf("rate_limit:%s", ip)

	// Step 1: Increment the count
	incrURL := fmt.Sprintf("%s/INCR/%s", upstashURL, key)
	req, _ := http.NewRequest("GET", incrURL, nil)
	req.Header.Set("Authorization", "Bearer "+upstashToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Result int `json:"result"`
	}
	json.Unmarshal(body, &result)

	count := result.Result

	// Step 2: Set expiry only if it's the first request
	if count == 1 {
		expireURL := fmt.Sprintf("%s/EXPIRE/%s/%d", upstashURL, key, windowSeconds)
		expireReq, _ := http.NewRequest("GET", expireURL, nil)
		expireReq.Header.Set("Authorization", "Bearer "+upstashToken)
		http.DefaultClient.Do(expireReq) // you can ignore the response here
	}

	return count > limit, nil
}
