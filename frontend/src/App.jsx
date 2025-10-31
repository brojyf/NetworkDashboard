import './App.css'
import { useState } from "react";
import Header from "./components/Header/Header";
import Section from "./components/Section/Section"
import SelectBox from "./components/SelectBox/SelectBox"
import { mockData } from "./MockData.js"

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
  const data = mockData[categoryMap[category]] || null;

  return (
    <div>
      <div className="header">
        <Header />
      </div>

      <SelectBox value={category} onChange={setCategory} />

      <div className="charts">
        {data && data.map((item, index) => (
        <Section key={index} data={item} />
      ))}
      </div>
      
    </div>
  )
}

export default App
