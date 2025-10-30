import React, { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("连接中...");

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => console.log("✅ WebSocket 已连接");
    ws.onmessage = (event) => setMessage(event.data);
    ws.onclose = () => console.log("❌ WebSocket 已关闭");
    ws.onerror = (err) => console.error("WebSocket 错误:", err);

    return () => ws.close();
  }, []);

  return (
    <div style={{ textAlign: "center", marginTop: "50px", fontSize: "24px" }}>
      <h1>实时时间更新</h1>
      <p>{message}</p>
    </div>
  );
}

export default App;
