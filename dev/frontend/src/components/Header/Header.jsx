import "./Header.css";

export default function Header() {
    return (
        <header className="header">
          <div className="content">
            <h1 className="title">Network Dashboard</h1>
            <div className="metric-tags">
              <span className="pill">Latency</span>
              <span className="pill">Jitter</span>
              <span className="pill">Traceroute</span>
            </div>
          </div>
        </header>
    );
}
