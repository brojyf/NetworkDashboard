// mockTraceroute.js
export const tracerouteData = {
  target: "google.com",
  hops: [
    {
      hop: 1,
      ip: "192.168.1.1",
      hostname: "router.home",
      latency: [1.12, 0.98, 1.05]
    },
    {
      hop: 2,
      ip: "10.22.0.1",
      hostname: "isp-gateway.local",
      latency: [9.8, 10.1, 9.9]
    },
    {
      hop: 3,
      ip: null,
      hostname: null,
      latency: ["*", "*", "*"]
    },
    {
      hop: 4,
      ip: "142.250.68.14",
      hostname: "google.com",
      latency: [22.4, 24.1, 23.5]
    }
  ]
};
