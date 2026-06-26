package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationArea{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("error closing body: %v\n", err)
		}
	}()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationResp := LocationArea{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
