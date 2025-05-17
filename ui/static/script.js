document.addEventListener("DOMContentLoaded", () => {
  const packsInput = document.getElementById("packs");
  const amountInput = document.getElementById("amount");
  const submitBtn = document.getElementById("submitBtn");
  const resultEl = document.getElementById("result");

  let apiUrl;

  init();

  async function init() {
    try {
      apiUrl = await loadConfig();
      submitBtn.addEventListener("click", handleSubmit);
    } catch (error) {
      showError(`Error loading config: ${error.message}`);
      submitBtn.disabled = true;
    }
  }

  async function loadConfig() {
    const res = await fetch("config.json");
    if (!res.ok) throw new Error(`Failed to load config.json: ${res.statusText}`);

    const config = await res.json();
    if (!config.apiUrl) throw new Error("apiUrl not found in config.json");

    return config.apiUrl;
  }

  async function handleSubmit() {
    const packSizes = parsePackSizes(packsInput.value);
    const amount = parseAmount(amountInput.value);

    if (!packSizes || amount === null) return;

    const payload = { pack_sizes: packSizes, amount };

    try {
      const response = await fetch(`${apiUrl}/shipment-packs`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorData = await safeJson(response);
        showError(`Error: ${errorData.error || response.statusText}`);
        return;
      }

      const data = await response.json();
      showResult(data);
    } catch (err) {
      showError(`Network error: ${err.message}`);
    }
  }

  function parsePackSizes(input) {
    if (!input.trim()) {
      alert("Pack Sizes cannot be empty");
      return null;
    }

    const sizes = input.split(",").map(s => parseInt(s.trim())).filter(n => !isNaN(n));
    if (sizes.length === 0) {
      alert("Please enter valid pack sizes");
      return null;
    }

    return sizes;
  }

  function parseAmount(input) {
    const value = parseInt(input.trim());
    if (isNaN(value) || value <= 0) {
      alert("Amount must be a positive number");
      return null;
    }

    return value;
  }

  async function safeJson(response) {
    try {
      return await response.json();
    } catch {
      return {};
    }
  }

  function showResult(data) {
    resultEl.textContent = JSON.stringify(data, null, 2);
  }

  function showError(message) {
    resultEl.textContent = message;
  }
});
