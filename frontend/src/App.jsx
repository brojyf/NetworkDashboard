import './App.css'
import Header from "./components/Header/Header";
import Section from "./components/Section/Section"

function App() {

  const data =  {
    website: "google.com",
    latency: {
      title: "Latency (ms)",
      labels: ["Mon", "Tue", "Wed", "Thu", "Fri"],
      data: [120, 90, 150, 80, 130]
    },
    package_lost: {
      title: "Package Loss (%)",
      labels: ["Mon", "Tue", "Wed", "Thu", "Fri"],
      data: [1, 3, 5, 0, 0]
    },
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
  }

  return (
      <div>
        <div className="header">
          <Header />
        </div>

        <h1>Selected: Search Engine.</h1>

        <Section
          data={data}
         />
        <Section
          data={data}
         />
      </div>
  )
}

export default App
