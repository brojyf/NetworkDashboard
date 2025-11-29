package Handler

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

const dataFilePath = "data/ping_results.csv"
const hopFilePath = "data/hop_results.csv"
const timeColumn = 0
const siteColumn = 1
const typeColumn = 2
const latencyColumn = 3
const jitterColumn = 4
const pick = 5

var categoryMap = map[string]string{
	"search_engine":   "SearchEngine",
	"ai":              "AI",
	"cdn":             "CDN",
	"social":          "Social",
	"cloud":           "Cloud",
	"video_streaming": "VideoStreaming",
}

type WebsiteData struct {
	Website string      `json:"website"`
	Latency MetricBlock `json:"latency"`
	Jitter  MetricBlock `json:"jitter"`
	Hops    []Hop       `json:"hops"`
}
type MetricBlock struct {
	Title  string    `json:"title"`
	Labels []string  `json:"labels"`
	Data   []float64 `json:"data"`
}
type Hop struct {
	Hop      int           `json:"hop"`
	IP       *string       `json:"ip"`
	Hostname *string       `json:"hostname"`
	Latency  []interface{} `json:"latency"`
}

type siteAccumulator struct {
	labels  []string
	latency []float64
	jitter  []float64
}

func Handler(c *gin.Context) {
	category := c.Query("category")
	hops := loadHopData()
	a, b := getTargetRows(categoryMap[category])
	c.JSON(200, mapToWebsiteData(a, b, hops))
}

func mapToWebsiteData(a, b [][]string, hops map[string][]Hop) []WebsiteData {
	return []WebsiteData{
		convertOneSite(a, hops),
		convertOneSite(b, hops),
	}
}

func convertOneSite(rows [][]string, hops map[string][]Hop) WebsiteData {
	if len(rows) == 0 {
		return WebsiteData{}
	}

	site := rows[0][siteColumn]
	acc := siteAccumulator{}
	for _, row := range rows {
		acc.labels = append(acc.labels, row[timeColumn])
		if v, err := strconv.ParseFloat(row[latencyColumn], 64); err == nil {
			acc.latency = append(acc.latency, v)
		}
		if v, err := strconv.ParseFloat(row[jitterColumn], 64); err == nil {
			acc.jitter = append(acc.jitter, v)
		}
	}

	return WebsiteData{
		Website: site,
		Latency: MetricBlock{
			Title:  "Latency (ms)",
			Labels: acc.labels,
			Data:   acc.latency,
		},
		Jitter: MetricBlock{
			Title:  "Jitter (ms)",
			Labels: acc.labels,
			Data:   acc.jitter,
		},
		Hops: hops[site],
	}
}

func getTargetRows(t string) ([][]string, [][]string) {
	// 1. Read
	file, err := os.Open(dataFilePath)
	if err != nil {
		log.Println("open file:", err)
		return nil, nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("close file:", err)
		}
	}(file)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("read csv:", err)
		return nil, nil
	}
	data := records[1:]

	// 2. Filter
	var filtered [][]string
	for _, row := range data {
		if len(row) > typeColumn && row[typeColumn] == t {
			filtered = append(filtered, row)
		}
	}
	if len(filtered) == 0 {
		return [][]string{}, [][]string{}
	}

	// 3. Top 5
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i][timeColumn] > filtered[j][timeColumn]
	})
	var siteA, siteB string
	var resultA, resultB [][]string
	countsA, countsB := 0, 0

	for _, row := range filtered {
		site := row[siteColumn]
		if siteA == "" {
			siteA = site
		}
		if site == siteA && countsA < pick {
			resultA = append(resultA, row)
			countsA++
			continue
		}
		if siteB == "" && site != siteA {
			siteB = site
		}
		if site == siteB && countsB < pick {
			resultB = append(resultB, row)
			countsB++
			continue
		}
		if countsA >= pick && countsB >= pick {
			break
		}
	}

	return reverseRows(resultA), reverseRows(resultB)
}

func reverseRows(rows [][]string) [][]string {
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i]
	}
	return rows
}

func loadHopData() map[string][]Hop {
	file, err := os.Open(hopFilePath)
	if err != nil {
		log.Println("open hop file:", err)
		return map[string][]Hop{}
	}
	defer func(file *os.File) {
		if cerr := file.Close(); cerr != nil {
			log.Println("close hop file:", cerr)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("read hop csv:", err)
		return map[string][]Hop{}
	}
	if len(records) <= 1 {
		return map[string][]Hop{}
	}

	result := make(map[string][]Hop)
	for _, row := range records[1:] {
		if len(row) < 5 {
			continue
		}

		site := row[0]
		hopVal, err := strconv.Atoi(row[1])
		if err != nil {
			continue
		}
		ip := pointerIfNotEmpty(row[2])
		hostname := pointerIfNotEmpty(row[3])

		var latency []interface{}
		if v, err := strconv.ParseFloat(row[4], 64); err == nil {
			latency = append(latency, v)
		} else {
			latency = []interface{}{}
		}

		result[site] = append(result[site], Hop{
			Hop:      hopVal,
			IP:       ip,
			Hostname: hostname,
			Latency:  latency,
		})
	}

	for site := range result {
		sort.Slice(result[site], func(i, j int) bool {
			return result[site][i].Hop < result[site][j].Hop
		})
	}

	return result
}

func pointerIfNotEmpty(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}
