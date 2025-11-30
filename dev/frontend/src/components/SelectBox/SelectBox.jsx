import "./SelectBox.css";

export default function SelectBox({ value, onChange, onRefresh, isRefreshing }) {
  const options = ["Search Engine", "CDN", "Video Streaming", "Social", "Cloud", "AI"];
  const disableRefresh = !value || isRefreshing;

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

      <button
        type="button"
        className="refresh-button"
        onClick={onRefresh}
        disabled={disableRefresh}
      >
        Refresh
      </button>
    </div>
  );
}
