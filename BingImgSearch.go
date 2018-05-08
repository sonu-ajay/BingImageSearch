package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ImageResult struct {
	Type            string `json:"_type"`
	Instrumentation struct {
		Type string `json:"_type"`
	} `json:"instrumentation"`
	ReadLink              string `json:"readLink"`
	WebSearchURL          string `json:"webSearchUrl"`
	TotalEstimatedMatches int    `json:"totalEstimatedMatches"`
	NextOffset            int    `json:"nextOffset"`
	Value                 []struct {
		WebSearchURL       string    `json:"webSearchUrl"`
		Name               string    `json:"name"`
		ThumbnailURL       string    `json:"thumbnailUrl"`
		DatePublished      time.Time `json:"datePublished"`
		ContentURL         string    `json:"contentUrl"`
		HostPageURL        string    `json:"hostPageUrl"`
		ContentSize        string    `json:"contentSize"`
		EncodingFormat     string    `json:"encodingFormat"`
		HostPageDisplayURL string    `json:"hostPageDisplayUrl"`
		Width              int       `json:"width"`
		Height             int       `json:"height"`
		Thumbnail          struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"thumbnail"`
		ImageInsightsToken string `json:"imageInsightsToken"`
		InsightsMetadata   struct {
			RecipeSourcesCount      int `json:"recipeSourcesCount"`
			BestRepresentativeQuery struct {
				Text         string `json:"text"`
				DisplayText  string `json:"displayText"`
				WebSearchURL string `json:"webSearchUrl"`
			} `json:"bestRepresentativeQuery"`
			PagesIncludingCount int `json:"pagesIncludingCount"`
			AvailableSizesCount int `json:"availableSizesCount"`
		} `json:"insightsMetadata"`
		ImageID     string `json:"imageId"`
		AccentColor string `json:"accentColor"`
	} `json:"value"`
	QueryExpansions []struct {
		Text         string `json:"text"`
		DisplayText  string `json:"displayText"`
		WebSearchURL string `json:"webSearchUrl"`
		SearchLink   string `json:"searchLink"`
		Thumbnail    struct {
			ThumbnailURL string `json:"thumbnailUrl"`
		} `json:"thumbnail"`
	} `json:"queryExpansions"`
	PivotSuggestions []struct {
		Pivot       string        `json:"pivot"`
		Suggestions []interface{} `json:"suggestions"`
	} `json:"pivotSuggestions"`
	RelatedSearches []struct {
		Text         string `json:"text"`
		DisplayText  string `json:"displayText"`
		WebSearchURL string `json:"webSearchUrl"`
		SearchLink   string `json:"searchLink"`
		Thumbnail    struct {
			ThumbnailURL string `json:"thumbnailUrl"`
		} `json:"thumbnail"`
	} `json:"relatedSearches"`
}

func main() {
	GetPage()
}

func GetPage() (*ImageResult, error) {
	apiKey := "BING_API_KEY"

	url := fmt.Sprintf("https://api.cognitive.microsoft.com/bing/v7.0/images/search?q=Plants&mkt=en-US&SafeSearch=strict&aspect=all&count=25&offset=25")
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("ocp-apim-subscription-key", apiKey)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(response.Body)
	var data ImageResult
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	for _, img := range data.Value {
		fmt.Println("Image Url: ", img.ContentURL)
	}
	fmt.Println("Next Data Offset: ", data.NextOffset)
	return &data, nil
}
