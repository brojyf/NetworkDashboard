import "./App.css";
import { useState, useEffect } from "react";
import Header from "./components/Header/Header";
import Section from "./components/Section/Section";
import SelectBox from "./components/SelectBox/SelectBox";


const categoryMap = {
    "Search Engine": "search_engine",
    "Video Streaming": "video_streaming",
    AI: "ai",
    CDN: "cdn",
    Social: "social",
    Cloud: "cloud",
};

function App() {

    const [category, setCategory] = useState("");
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);

    useEffect(() => {
        if (!category) return;

        const apiCategory = categoryMap[category];

        (async () => {
            setLoading(true);
            setError(null);
            setData(null);

            try {
                const res = await fetch(
                    `http://localhost:8080/api/query?category=${encodeURIComponent(
                        apiCategory,
                    )}`,
                );

                const json = await res.json();
                setData(json);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        })();
    }, [category]);

    return (
        <div>
            <Header />

            <SelectBox value={category} onChange={setCategory} />

            <div className="charts">
                {loading && <p>Loading...</p>}
                {error && <p style={{ color: "red" }}>Error: {error}</p>}

                {data &&
                    data.map((item, index) => <Section key={index} data={item} />)}
            </div>
        </div>
    );
}

export default App;
