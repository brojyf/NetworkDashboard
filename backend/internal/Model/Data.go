package model

type ServerData struct {
	SearchEngine   []WebsiteData `json:"searchEngine"`
	AI             []WebsiteData `json:"ai"`
	CDN            []WebsiteData `json:"cdn"`
	Social         []WebsiteData `json:"social"`
	Cloud          []WebsiteData `json:"cloud"`
	VideoStreaming []WebsiteData `json:"videoStreaming"`
}

type WebsiteData struct {
	Website     string      `json:"website"`
	Latency     MetricBlock `json:"latency"`
	PackageLost MetricBlock `json:"package_lost"`
	Hops        []Hop       `json:"hops"`
}

type MetricBlock struct {
	Title  string   `json:"title"`
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

type Hop struct {
	Hop      int           `json:"hop"`
	IP       *string       `json:"ip"`
	Hostname *string       `json:"hostname"`
	Latency  []interface{} `json:"latency"`
}
