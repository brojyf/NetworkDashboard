package model

func strPtr(s string) *string {
	return &s
}

var MockData = ServerData{
	SearchEngine: []WebsiteData{
		{
			Website: "www.google.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Jitter (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
		{
			Website: "www.bing.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Jitter (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
	},

	AI: []WebsiteData{
		{
			Website: "www.chatgpt.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
		{
			Website: "claud.ai",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
	},

	CDN: []WebsiteData{
		{
			Website: "www.cloudflare.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-ggateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
		{
			Website: "www.akamai.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
	},

	Social: []WebsiteData{
		{
			Website: "www.facebook.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},

		{
			Website: "x.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
	},

	Cloud: []WebsiteData{
		{
			Website: "www.icloud.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
		{
			Website: "drive.google.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
	},

	VideoStreaming: []WebsiteData{
		{
			Website: "www.youtube.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},

		{
			Website: "www.netflix.com",
			Latency: MetricBlock{
				Title:  "Latency (ms)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{120, 90, 150, 80, 130},
			},
			Jitter: MetricBlock{
				Title:  "Package Loss (%)",
				Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
				Data:   []int{1, 3, 5, 0, 0},
			},
			Hops: []Hop{
				{Hop: 1, IP: strPtr("192.168.1.1"), Hostname: strPtr("router.home"), Latency: []interface{}{1.12, 0.98, 1.05}},
				{Hop: 2, IP: strPtr("10.22.0.1"), Hostname: strPtr("isp-gateway.local"), Latency: []interface{}{9.8, 10.1, 9.9}},
				{Hop: 3, IP: nil, Hostname: nil, Latency: []interface{}{"*", "*", "*"}},
				{Hop: 4, IP: strPtr("142.250.68.14"), Hostname: strPtr("google.com"), Latency: []interface{}{22.4, 24.1, 23.5}},
			},
		},
	},
}
