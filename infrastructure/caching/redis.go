package caching

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/sync/singleflight"
)

// sfGroup is used to ensure only one loader call is active per key at any time.
var sfGroup singleflight.Group

// SetRedisValue sets a value in Redis by key.
func SetRedisValue(key, value string) error {
	fmt.Println("Setting value in Redis for key:", key, "with value:", value)
	upstashToken := os.Getenv("REDIS_TOKEN")
	upstashURL := os.Getenv("REDIS_URL")

	url := fmt.Sprintf("%s/SET/%s/%s", upstashURL, key, value)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+upstashToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Set Response: %s\n", string(body))
	return nil
}

// GetRedisValue retrieves a value from Redis by key.
// Returns an empty string if not present (cache miss).
func GetRedisValue(key string) (string, error) {
	upstashURL := os.Getenv("REDIS_URL")
	url := upstashURL + "/GET/" + key

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("REDIS_TOKEN"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("redis error: %s", res.Status)
	}

	var result struct {
		Result *string `json:"result"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}

	fmt.Println("Redis result:", result, result.Result)
	if result.Result == nil {
		return "", nil // Cache miss
	}
	return *result.Result, nil
}
// GetOrSetRedisValue checks the cache for a key, and if present returns it immediately.
// In the background, it refreshes the cache using the provided loader function.
// If the key is not present, it loads the data using singleflight, caches it, and returns it.
func GetOrSetRedisValue(key string, loader func() (string, error)) (string, error) {
    // First, check the cache.
    if value, err := GetRedisValue(key); err != nil {
        return "", err
    } else if value != "" {
        // Found in cache, return it immediately.
        // Trigger background refresh of the cache.
        go func() {
            loadedValue, err := loader()
            if err != nil {
                log.Println("Background loader failed for key:", key, "error:", err)
                return
            }
            if err := SetRedisValue(key, loadedValue); err != nil {
                log.Println("Failed to update redis cache for key:", key, "error:", err)
            }
        }()
        return value, nil
    }

    // If not found in cache, use singleflight to load the data.
    result, err, _ := sfGroup.Do(key, func() (interface{}, error) {
        // Loader function to load data from the backend.
        loadedValue, err := loader()
        if err != nil {
            return "", err
        }
        // Update the cache synchronously.
        if err := SetRedisValue(key, loadedValue); err != nil {
            log.Println("Failed to update redis cache for key:", key, "error:", err)
        }
        return loadedValue, nil
    })
    if err != nil {
        return "", err
    }
    return result.(string), nil
}


func DeleteRedisValue(key string) error {
    upstashToken := os.Getenv("REDIS_TOKEN")
    upstashURL := os.Getenv("REDIS_URL")
    
    // Build a URL that instructs Redis to delete the key.  
    // Adjust the endpoint (DEL) based on your Upstash implementation.
    url := fmt.Sprintf("%s/DEL/%s", upstashURL, key)
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        return err
    }
    req.Header.Add("Authorization", "Bearer "+upstashToken)
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    // Consume and discard the body.
    _, err = io.ReadAll(resp.Body)
    return err
}