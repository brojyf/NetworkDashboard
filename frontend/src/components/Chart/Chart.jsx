import { useEffect, useRef } from "react";
import Chart from "chart.js/auto";

export default function LineChart({graphData}) {
  const chartRef = useRef(null);
  const chartInstance = useRef(null);

  useEffect(() => {
    if (chartInstance.current)
        chartInstance.current.destroy();

    chartInstance.current = new Chart(chartRef.current, {
      type: "line",
      data: {
        labels: graphData.labels,
        datasets: [
          {
            label: graphData.title,
            data: graphData.data,
            borderWidth: 2,
            tension: 0.3,
          },
        ],
      },
      options: {
        responsive: true,
        maintainAspectRatio: false, 
      },
    });
  }, 
    [graphData.data, graphData.labels, graphData.title]
  );

  return (
    <div className="chart-box">
      <canvas ref={chartRef} />
    </div>
  );
}
