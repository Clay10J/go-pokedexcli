package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// req, err := http.NewRequest("GET", fullURL, nil)
	// if err != nil {
	// 	return LocationAreasResponse{}, err
	// }

	// resp, err := c.httpClient.Do(req)
	// if err != nil {
	// 	return LocationAreasResponse{}, err
	// }

	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(dat, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResponse, nil
	}

	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResponse, nil
}

func (c *Client) ListPokemonInLocationArea(name *string) (LocationAreasSpecificResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint + fmt.Sprintf("/%s", *name)

	// req, err := http.NewRequest("GET", fullURL, nil)
	// if err != nil {
	// 	return LocationAreasSpecificResponse{}, err
	// }

	// resp, err := c.httpClient.Do(req)
	// if err != nil {
	// 	return LocationAreasSpecificResponse{}, err
	// }

	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationAreasSpecificResponse := LocationAreasSpecificResponse{}
		err := json.Unmarshal(dat, &locationAreasSpecificResponse)
		if err != nil {
			return LocationAreasSpecificResponse{}, err
		}

		return locationAreasSpecificResponse, nil
	}

	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationAreasSpecificResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasSpecificResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasSpecificResponse{}, err
	}

	locationAreasSpecificResponse := LocationAreasSpecificResponse{}
	err = json.Unmarshal(dat, &locationAreasSpecificResponse)
	if err != nil {
		return LocationAreasSpecificResponse{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasSpecificResponse, nil
}
