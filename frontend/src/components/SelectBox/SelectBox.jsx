import React from "react";

export default function SelectBox({ value, onChange }) {
  const options = ["Search Engine", "CDN", "Video Steaming", "Social", "Cloud", "AI"];

  return (
    <div style={{ marginBottom: "16px" }}>
      <label style={{ fontWeight: "bold", marginRight: "8px" }}>
        Category:
      </label>

      <select
        value={value}
        onChange={(e) => onChange(e.target.value)}
        style={{
          padding: "8px 12px",
          fontSize: "16px",
          borderRadius: "8px",
          border: "1px solid #ccc",
          outline: "none",
        }}
      >
        {/* 默认空选项 */}
        <option value="" disabled selected>
          Choose one category
        </option>

        {options.map((opt) => (
          <option key={opt} value={opt}>
            {opt}
          </option>
        ))}
      </select>
    </div>
  );
}
