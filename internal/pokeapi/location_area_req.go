package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageURL != nil {
		fullUrl = *pageURL
	}

	val, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(val, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}	
		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return 	LocationAreasResp{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullUrl, data)

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}

func (c *Client) ListPokemonInArea(locationName string) (LocationAreaNameResp, error) {
	endpoint := "/location-area/"
	fullUrl := baseUrl + endpoint + locationName

	val, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreaNameResp := LocationAreaNameResp{}
		err := json.Unmarshal(val, &locationAreaNameResp)
		if err != nil {
			return LocationAreaNameResp{}, err
		}	
		return locationAreaNameResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return 	LocationAreaNameResp{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaNameResp{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaNameResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaNameResp{}, err
	}

	c.cache.Add(fullUrl, data)

	locationAreaNameResp := LocationAreaNameResp{}
	err = json.Unmarshal(data, &locationAreaNameResp)
	if err != nil {
		return LocationAreaNameResp{}, err
	}

	return locationAreaNameResp, nil
}