package service

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	internalconfig "github.com/brojyf/NetworkDashboard/internal/config"
)

type QueryService struct {
	pingFilePath string
	hopFilePath  string
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

type siteSeriesAccumulator struct {
	labels  []string
	latency []float64
	jitter  []float64
}

func NewQueryService() *QueryService {
	return &QueryService{
		pingFilePath: internalconfig.PingResultsCSVPath,
		hopFilePath:  internalconfig.HopResultsCSVPath,
	}
}

func (s *QueryService) QueryByCategory(category string) []WebsiteData {
	siteType := internalconfig.SiteTypeByCategory(category)
	hopsBySite := s.loadHopsBySite()
	firstSiteRows, secondSiteRows := s.loadRecentRowsForTwoSites(siteType)

	return []WebsiteData{
		s.buildWebsiteData(firstSiteRows, hopsBySite),
		s.buildWebsiteData(secondSiteRows, hopsBySite),
	}
}

func (s *QueryService) loadRecentRowsForTwoSites(siteType string) ([][]string, [][]string) {
	records, err := s.readAllCSVRows(s.pingFilePath)
	if err != nil || len(records) <= 1 {
		return [][]string{}, [][]string{}
	}

	var filteredRows [][]string
	for _, row := range records[1:] {
		if len(row) <= internalconfig.JitterColumn {
			continue
		}
		if row[internalconfig.SiteTypeColumn] == siteType {
			filteredRows = append(filteredRows, row)
		}
	}
	if len(filteredRows) == 0 {
		return [][]string{}, [][]string{}
	}

	sort.Slice(filteredRows, func(i, j int) bool {
		left := parseTimestampOrClock(filteredRows[i][internalconfig.TimestampColumn])
		right := parseTimestampOrClock(filteredRows[j][internalconfig.TimestampColumn])
		return left.After(right)
	})

	var firstSite, secondSite string
	var firstSiteRows, secondSiteRows [][]string
	firstCount, secondCount := 0, 0

	for _, row := range filteredRows {
		site := row[internalconfig.SiteColumn]
		if firstSite == "" {
			firstSite = site
		}

		if site == firstSite && firstCount < internalconfig.MaxPointsPerSite {
			firstSiteRows = append(firstSiteRows, row)
			firstCount++
			continue
		}

		if secondSite == "" && site != firstSite {
			secondSite = site
		}

		if site == secondSite && secondCount < internalconfig.MaxPointsPerSite {
			secondSiteRows = append(secondSiteRows, row)
			secondCount++
			continue
		}

		if firstCount >= internalconfig.MaxPointsPerSite && secondCount >= internalconfig.MaxPointsPerSite {
			break
		}
	}

	return reverseRows(firstSiteRows), reverseRows(secondSiteRows)
}

func (s *QueryService) buildWebsiteData(rows [][]string, hopsBySite map[string][]Hop) WebsiteData {
	if len(rows) == 0 {
		return WebsiteData{}
	}

	site := rows[0][internalconfig.SiteColumn]
	acc := siteSeriesAccumulator{}
	for _, row := range rows {
		if len(row) <= internalconfig.JitterColumn {
			continue
		}

		timestamp := parseTimestampOrClock(row[internalconfig.TimestampColumn])
		acc.labels = append(acc.labels, timestamp.Format("15:04"))

		if v, err := strconv.ParseFloat(row[internalconfig.LatencyColumn], 64); err == nil {
			acc.latency = append(acc.latency, v)
		}
		if v, err := strconv.ParseFloat(row[internalconfig.JitterColumn], 64); err == nil {
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
		Hops: hopsBySite[site],
	}
}

func (s *QueryService) loadHopsBySite() map[string][]Hop {
	records, err := s.readAllCSVRows(s.hopFilePath)
	if err != nil || len(records) <= 1 {
		return map[string][]Hop{}
	}

	hopsBySite := make(map[string][]Hop)
	for _, row := range records[1:] {
		if len(row) < internalconfig.MinHopRowColumns {
			continue
		}

		hopNumber, err := strconv.Atoi(row[1])
		if err != nil {
			continue
		}

		var latency []interface{}
		if value, err := strconv.ParseFloat(row[4], 64); err == nil {
			latency = append(latency, value)
		} else {
			latency = []interface{}{}
		}

		site := row[0]
		hopsBySite[site] = append(hopsBySite[site], Hop{
			Hop:      hopNumber,
			IP:       stringPtrIfNotEmpty(row[2]),
			Hostname: stringPtrIfNotEmpty(row[3]),
			Latency:  latency,
		})
	}

	for site := range hopsBySite {
		sort.Slice(hopsBySite[site], func(i, j int) bool {
			return hopsBySite[site][i].Hop < hopsBySite[site][j].Hop
		})
	}

	return hopsBySite
}

func (s *QueryService) readAllCSVRows(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println("open csv file:", err)
		return nil, err
	}
	defer func(file *os.File) {
		if cerr := file.Close(); cerr != nil {
			log.Println("close csv file:", cerr)
		}
	}(file)

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println("read csv file:", err)
		return nil, err
	}
	return records, nil
}

func reverseRows(rows [][]string) [][]string {
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i]
	}
	return rows
}

func stringPtrIfNotEmpty(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func parseTimestampOrClock(value string) time.Time {
	if t, err := time.Parse(time.RFC3339, value); err == nil {
		return t
	}
	if t, err := time.Parse("15:04", value); err == nil {
		now := time.Now()
		return time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location())
	}
	return time.Time{}
}
