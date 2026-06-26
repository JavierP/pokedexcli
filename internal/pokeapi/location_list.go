package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(s *string) (LocationsResp, error) {
	url := ""
	if s == nil {
		url = baseURL + "/location-area"
	} else {
		url = *s
	}
	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationsResp{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationsResp{}, err
		}
		return locationsResp, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationsResp{}, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Printf("error closing body: %v\n", err)
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResp{}, err
	}
	if res.StatusCode > 299 {
		return LocationsResp{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}
	c.cache.Add(url, body)
	locationsResp := LocationsResp{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return LocationsResp{}, err
	}
	return locationsResp, nil
}
