package helm

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	DefaultHelmRepositoryURL = "https://charts.bitnami.com/bitnami"
)

type (
	File struct {
		APIVersion string                   `yaml:"apiVersion" json:"apiVersion"`
		Entries    map[string][]serviceInfo `yaml:"entries" json:"entries"`
		Generated  string                   `yaml:"generated" json:"generated"`
	}

	Annotations struct {
		Category string `yaml:"category" json:"category"` // category
	}

	serviceInfo struct {
		Annotations    *Annotations `yaml:"annotations" json:"annotations,omitempty"`
		Name           string       `yaml:"name" json:"name"`               // chart name
		AppVersion     string       `yaml:"appVersion" json:"appVersion"`   // app version
		CreatedAt      string       `yaml:"created" json:"createdAt"`       // created date
		Deprecated     bool         `yaml:"deprecated" json:"deprecated"`   // is chart deprecated
		Description    string       `yaml:"description" json:"description"` // description
		Home           string       `yaml:"home" json:"home"`               // base url
		Sources        []string     `yaml:"sources" json:"sources"`         // repository info
		Urls           []string     `yaml:"urls" json:"urls"`               // tarball url
		PackageVersion string       `yaml:"version" json:"packageVersion"`  // pacakge version
		Icon           string       `yaml:"icon" json:"icon,omitempty"`     // icon url
	}

	chartList struct {
		Name string      `json:"name"`
		Data serviceInfo `json:"data"`
		Repo string      `json:"repo"`
	}
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

	_, err := url.ParseRequestURI(repo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, err := searchRepo(repo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	chartList := jsonToChartList(repo, result.Entries)
	sort.Slice(chartList, func(i, j int) bool {
		return chartList[i].Name < chartList[j].Name
	})

	c.JSON(http.StatusOK, chartList)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// _ = json.NewEncoder(w).Encode(chartList)

	// return nil
}

func searchRepo(repoPath string) (File, error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	url, err := url.ParseRequestURI(repoPath)
	if err != nil {
		return File{}, errors.Wrap(err, fmt.Sprintf("invalid helm chart URL: %s", repoPath))
	}

	/* get repository index.yaml content */
	url.Path = path.Join(url.Path, "index.yaml")
	resp, err := client.Get(url.String())
	if err != nil {
		return File{}, errors.Wrap(err, "failed to get index file content")
	}

	var file File
	err = yaml.NewDecoder(resp.Body).Decode(&file)
	if err != nil {
		return File{}, errors.Wrap(err, "failed to decode index file content")
	}

	return file, nil
}

func jsonToChartList(repo string, data map[string][]serviceInfo) (cl []chartList) {
	for k, v := range data {
		cl = append(cl, chartList{
			Name: k,
			Data: v[0],
			Repo: repo,
		})
	}
	return cl
}
