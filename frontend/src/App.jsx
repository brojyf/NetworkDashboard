import './App.css'
import { useState, useEffect } from "react";
import Header from "./components/Header/Header";
import Section from "./components/Section/Section"
import SelectBox from "./components/SelectBox/SelectBox"
import { mockData } from "./MockData"

function App() {

  const categoryMap = {
    "Search Engine": "searchEngine",
    "Video Streaming": "videoStreaming",
    "AI": "ai",
    "CDN": "cdn",
    "Social": "social",
    "Cloud": "cloud",
  };

  const [category, setCategory] = useState("");
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!category) return;

    async function fetchData() {
      setLoading(true);
      setError(null);
      setData(null);

      try {
        const res = await fetch(`http://localhost:8080/api/query?category=${encodeURIComponent(category)}`);

        // if (!res.ok) {
        //   throw new Error(`Server error: ${res.status}`);
        // }

        // const json = await res.json();
        // setData(json.data); 

      } catch (err) {
        // setError(err.message);
          const serverData = mockData[categoryMap[category]] || null;
          setData(serverData);
      } finally {
        setLoading(false);
      }
    }

    fetchData();
  }, [category]);

  return (
    <div>
      <Header />

      <SelectBox value={category} onChange={setCategory} />

      <div className="charts">

        {loading && <p>Loading...</p>}
        {error && <p style={{ color: "red" }}>Error: {error}</p>}
        {data && data.map((item, index) => (
          <Section key={index} data={item} />
        ))}
      </div>
    </div>
  )
}

export default App;
