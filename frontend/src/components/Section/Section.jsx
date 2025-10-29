import LineChart from "../Chart/Chart";
import Traceroute from "../Traceroute/Traceroute";
import "./Section.css"

export default function Section({website, lat, packl, tr}) {
    return (
        <div>
            <h1>{website}</h1>
            <div className="charts-row">
                <LineChart
                  graphData={lat}
                 />
                <LineChart
                  graphData={packl}
                 />
            </div>
            <div>
                <Traceroute />
            </div>
        </div>
    );
}