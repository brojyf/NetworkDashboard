package config

const (
	PingResultsCSVPath = "data/ping_results.csv"
	HopResultsCSVPath  = "data/hop_results.csv"

	TimestampColumn = 0
	SiteColumn      = 1
	SiteTypeColumn  = 2
	LatencyColumn   = 3
	JitterColumn    = 4

	MaxPointsPerSite = 5
	MinHopRowColumns = 5
)

var categoryToSiteType = map[string]string{
	"search_engine":   "SearchEngine",
	"ai":              "AI",
	"cdn":             "CDN",
	"social":          "Social",
	"cloud":           "Cloud",
	"video_streaming": "VideoStreaming",
}

func SiteTypeByCategory(category string) string {
	return categoryToSiteType[category]
}
