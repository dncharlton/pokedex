package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationAreasResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status codeL %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreasResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}

func (c *Client) ListLocationEncounters(location string) (LocationEncounterResp, error) {
	endpoint := "/location-area/" + location
	fullURL := baseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationEncounterResp := LocationEncounterResp{}
		err := json.Unmarshal(dat, &locationEncounterResp)
		if err != nil {
			return LocationEncounterResp{}, err
		}
		return locationEncounterResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationEncounterResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationEncounterResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationEncounterResp{}, fmt.Errorf("bad status codeL %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationEncounterResp{}, err
	}

	locationEncounterResp := LocationEncounterResp{}
	err = json.Unmarshal(dat, &locationEncounterResp)

	if err != nil {
		return LocationEncounterResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationEncounterResp, nil
}