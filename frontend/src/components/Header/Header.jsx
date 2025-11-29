import "./Header.css";

export default function Header() {
    return (
        <header className="header">
          <div className="content">
            <h1 className="title">Network Dashboard</h1>
            <p className="subtitle">Latency • Jitter • Traceroute</p>
          </div>
        </header>
    );
}
