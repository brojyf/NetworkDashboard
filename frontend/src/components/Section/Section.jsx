import LineChart from "../Chart/Chart";
import Traceroute from "../Traceroute/Traceroute";
import "./Section.css"

export default function Section({data}) {
    return (
        <div>
            <h1>{data.website}</h1>
            <div className="charts-row">
                <LineChart
                  graphData={data.latency}
                 />
                <LineChart
                  graphData={data.jitter}
                 />
            </div>
            <div>
                <Traceroute
                  data={data.hops}
                 />
            </div>
        </div>
    );
}
