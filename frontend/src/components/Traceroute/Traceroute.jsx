import { useState } from "react";
import "./Traceroute.css"; 

export default function Traceroute({data}) { 

  const [showTable, setShowTable] = useState(false);

  const hopCount = data.length;
  const hopSequence = data
    .map((hop) => hop.hostname || "-")
    .join(" → ");

  return ( 
    <div className="traceroute-container"> 

      <div className="hop-header">
        <div className="hop-summary"> 
          <span>🔁 Total Hops: </span> 
          <span className="hop-number">{hopCount}</span> 
        </div>

        <button
          className="toggle-btn"
          onClick={() => setShowTable(!showTable)}
        >
          {showTable ? "Hide" : "More"}
        </button>
      </div>

      <div className="hop-path">
        {hopSequence}
      </div>

      {showTable && (
        <table className="traceroute-table"> 
          <thead> 
            <tr>
              <th>Hop</th>
              <th>IP Address</th>
              <th>Hostname</th>
              <th>Latency (ms)</th>
            </tr>
          </thead>
          <tbody>
            {data.map((hop, index) => (
              <tr key={index}>
                <td>{hop.hop}</td>
                <td>{hop.ip || "-"}</td>
                <td>{hop.hostname || "-"}</td>
                <td>
                  {Array.isArray(hop.latency)
                    ? hop.latency.join(" / ")
                    : hop.latency}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
