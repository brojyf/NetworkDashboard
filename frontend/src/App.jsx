import "./App.css";
import { useState, useEffect } from "react";
import Header from "./components/Header/Header";
import Section from "./components/Section/Section";
import SelectBox from "./components/SelectBox/SelectBox";


const categoryMap = {
    "Search Engine": "search_engine",
    "Video Streaming": "video_streaming",
    "AI": "ai",
    "CDN": "cdn",
    "Social": "social",
    "Cloud": "cloud",
};

function App() {

    const [category, setCategory] = useState("");
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [refreshCounter, setRefreshCounter] = useState(0);

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
    }, [category, refreshCounter]);

    const handleRefresh = () => {
        if (!category) return;
        setRefreshCounter((count) => count + 1);
    };

    const isDataArray = Array.isArray(data);
    const safeData = isDataArray ? data : [];

    return (
        <div className="container">
            <Header />

            <div className="select-box-card">
                <SelectBox
                    value={category}
                    onChange={setCategory}
                    onRefresh={handleRefresh}
                    isRefreshing={loading}
                />
            </div>

            <div className="charts">
                {!category && (
                    <div className="placeholder-box">
                        <p className="placeholder">Select a category to load data</p>
                    </div>
                )}

                {category && (
                    <>
                        {loading && <p>Loading...</p>}

                        {!loading && error && (
                            <p style={{ color: "red" }}>Error: {error}</p>
                        )}

                        {!loading && !error && (!data || !isDataArray) && (
                            <div className="placeholder-box">
                                <p className="placeholder">
                                    Server is still loading. Try again later.
                                </p>
                            </div>
                        )}

                        {!loading && !error && isDataArray && data.length === 0 && (
                            <div className="placeholder-box">
                                <p className="placeholder">No data available</p>
                            </div>
                        )}

                        {!loading && !error && isDataArray && data.length > 0 && (
                            safeData.map((item, index) => (
                                <Section key={index} data={item} />
                            ))
                        )}
                    </>
                )}
            </div>

        </div>
    );
}

export default App;
