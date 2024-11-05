document.getElementById("extractJobs").addEventListener("click", async () => {
  // Get the active tab in the current window
  const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });

  // Execute the script to extract job information
  const [{ result }] = await chrome.scripting.executeScript({
    target: { tabId: tab.id },
    function: extractJobInformation,
  });

  // Display the result in the output element
  document.getElementById("output").textContent = result
    ? JSON.stringify(result, null, 2)
    : "No jobs found.";
});
