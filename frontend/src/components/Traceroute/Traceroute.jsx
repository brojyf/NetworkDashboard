import { tracerouteData } from "./mockTraceroute"; 
import "./Traceroute.css"; 

export default function Traceroute() { 
  const hopCount = tracerouteData.hops.length; 
  const hopSequence = tracerouteData.hops 
    .map((hop) => hop.hostname || "-") .join(" → "); 
  return ( 
    <div className="traceroute-container"> 
      <div className="hop-summary"> 
        <span>🔁 Total Hops: </span> 
        <span className="hop-number">{hopCount}</span> 
      </div> 
      <div className="hop-path"> 
        {hopSequence} 
      </div> 
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
          {tracerouteData.hops.map((hop, index) => ( 
            <tr key={index}> 
              <td>{hop.hop}</td> 
              <td>{hop.ip || "-"}</td>
              <td>{hop.hostname || "-"}</td> 
              <td> {Array.isArray(hop.latency) 
                ? hop.latency.join(" / ") 
                : hop.latency} 
              </td> 
            </tr> 
          ))} 
        </tbody> 
      </table> 
    </div> 
  ); 
}