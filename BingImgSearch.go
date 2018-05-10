package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bing/{query}", GetBingImages).Methods("GET")
	router.HandleFunc("/bing/{query}/{pageSize}", GetBingImages).Methods("GET")
	router.HandleFunc("/bing/{query}/{pageSize}/{offSet}", GetBingImages).Methods("GET")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

//GetBingImages returns nothing
func GetBingImages(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := params["query"]
	pageSize, err := strconv.Atoi(params["pageSize"])
	if err != nil {
		pageSize = 25
	}
	offSet, err := strconv.Atoi(params["offSet"])
	if err != nil {
		offSet = 0
	}
	imageResponse, resperr := GetPage(query, pageSize, offSet)
	if resperr != nil {
		json.NewEncoder(w).Encode(&resperr)
	} else {
		json.NewEncoder(w).Encode(&imageResponse)
	}
}

//GetPage returns ImageResult
func GetPage(searchQuery string, pagesize int, offSet int) (*ImageResult, error) {
	apiKey := "3cbc13236b2c4ee684fb2de927c2f685"

	url := fmt.Sprintf("https://api.cognitive.microsoft.com/bing/v7.0/images/search?q=%s&mkt=en-US&SafeSearch=strict&aspect=all&count=%d&offset=%d", searchQuery, pagesize, offSet)
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
	// for _, img := range data.Value {
	// 	fmt.Println("Image Url: ", img.ContentURL)
	// }
	fmt.Println("Served request with following Results")
	fmt.Println("Query: ", searchQuery)
	fmt.Println("PageSize: ", pagesize)
	fmt.Println("Offset: ", offSet)
	fmt.Println("Results Expected: ", data.TotalEstimatedMatches)
	fmt.Println("Next Offset: ", data.NextOffset)
	return &data, nil
}

//BingImageResult returns ImageResult
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
