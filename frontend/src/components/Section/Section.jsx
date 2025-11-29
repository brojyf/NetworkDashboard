import LineChart from "../Chart/Chart";
import Traceroute from "../Traceroute/Traceroute";
import "./Section.css"

export default function Section({data}) {
    const hops = Array.isArray(data?.hops) ? data.hops : [];
    const latency = data?.latency;
    const jitter = data?.jitter;
    const website = data?.website || "Unknown";

    const hasLatency = latency && Array.isArray(latency.data) && Array.isArray(latency.labels);
    const hasJitter = jitter && Array.isArray(jitter.data) && Array.isArray(jitter.labels);
    const hasHops = Array.isArray(hops);
    const canRender = hasLatency && hasJitter && hasHops;
    const hasHopContent = hasHops && hops.length > 0;

    return (
        <section className="section-card">
            <h1>{website}</h1>
            {canRender ? (
                <>
                    <div className="charts-row">
                        <LineChart graphData={latency} />
                        <LineChart graphData={jitter} />
                    </div>
                    <div className="traceroute-block">
                        {hasHopContent ? (
                            <Traceroute data={hops} />
                        ) : (
                            <div className="section-placeholder">
                                Traceroute is still retrieving data. Please try again later.
                            </div>
                        )}
                    </div>
                </>
            ) : (
                <div className="section-placeholder">
                    Server is retrieving data. Please try again later.
                </div>
            )}
        </section>
    );
}
