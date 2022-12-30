package helm

import (
	"log"

	"github.com/gin-gonic/gin"
)

const (
	DefaultHelmRepositoryURL = "https://charts.bitnami.com/bitnami"
)

//	@BasePath	/api

// helm handler example
//	@Summary	get charts exmaple
//	@Schemes
//	@Description	get chart list exmaple
//	@Tags			Helm
//	@Accept			json
//	@Produce		json
//	@Success		200	{json}	get chart list
//	@Router			/user [get]
func GetChartsHandler(c *gin.Context) {
	log.Printf("[INFO] [message: GET CHART LIST START]")

	repo := c.Query("repo")
	if repo == "" {
		repo = DefaultHelmRepositoryURL
	}

	// _, err := url.ParseRequestURI(repo)
	// if err != nil {

	// }

	// if err != nil {
	// 	return httperror.BadRequest(100, fmt.Sprintf("Provided URL '%s' is not valid", repo), err)
	// 	// return &httperror.HandlerError{StatusCode: http.StatusBadRequest, Message: "Bad request", Err: errors.Wrap(err, fmt.Sprintf("provided URL %q is not valid", repo))}
	// }

	// result, err := searchRepo(repo)
	// if err != nil {
	// 	return httperror.BadRequest(100, "Search failed", err)
	// }

	// chartList := jsonToChartList(repo, result.Entries)
	// sort.Slice(chartList, func(i, j int) bool {
	// 	return chartList[i].Name < chartList[j].Name
	// })

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// _ = json.NewEncoder(w).Encode(chartList)

	// return nil
}
