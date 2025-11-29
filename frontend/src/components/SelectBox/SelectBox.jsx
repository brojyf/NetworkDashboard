import "./SelectBox.css";   

export default function SelectBox({ value, onChange }) {
  const options = ["Search Engine", "CDN", "Video Streaming", "Social", "Cloud", "AI"];

  return (
    <div className="select-wrapper">
      <label className="select-label">
        Category:
      </label>

      <select
        className="select-dropdown"
        value={value}
        onChange={(e) => onChange(e.target.value)}
      >
        <option value="" disabled>
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
