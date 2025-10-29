import './App.css'
import Header from "./components/Header/Header";
import Section from "./components/Section/Section"

function App() {

  const data =  {
    website: "google.com",
    latency_data: {
      title: "Latency (ms)",
      labels: ["Mon", "Tue", "Wed", "Thu", "Fri"],
      data: [120, 90, 150, 80, 130]
    },
    package_lost_data: {
      title: "Package Loss (%)",
      labels: ["Mon", "Tue", "Wed", "Thu", "Fri"],
      data: [1, 3, 5, 0, 0]
    },
    hop_data: {
      count: 4,
    }
  }

  return (
      <div>
        <div className="header">
          <Header />
        </div>
        <Section
          website={data.website}
          lat={data.latency_data}
          packl={data.package_lost_data}
         />
      </div>
  )
}

export default App
