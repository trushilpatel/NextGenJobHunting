// content.js

let throttleTimeout;
const processedJobIds = new Set(); // Store processed job IDs
let isProcessing = false; // Flag to indicate if the script is currently processing

// Function to hide promoted jobs
function hidePromotedJobs() {
  const promotedJobs = document.querySelectorAll(
    ".job-card-container[data-promotion-type]"
  );
  promotedJobs.forEach((job) => {
    job.style.display = "none";
  });
}

// Function to hide applied jobs
function hideAppliedJobs() {
  const appliedJobs = document.querySelectorAll(
    '.job-card-container[data-job-state="applied"]'
  );
  appliedJobs.forEach((job) => {
    job.style.display = "none";
  });
}

// Function to add buttons to each job listing
function addJobButtons() {
  const jobCards = document.querySelectorAll(".job-card-list__title");
  const fragment = document.createDocumentFragment(); // Create a document fragment
  const newJobIds = []; // Array to keep track of newly added job IDs

  jobCards.forEach((card) => {
    // Get job ID from the parent container
    const jobId = card.closest("[data-job-id]").getAttribute("data-job-id");

    // Check if the job title exists and hasn't been processed yet
    if (card.innerText.trim() && !processedJobIds.has(jobId)) {
      // Check if the button already exists
      if (!card.querySelector(".custom-button")) {
        const button = document.createElement("button");
        button.innerText = "Get Job Details";
        button.className = "custom-button"; // Add a custom class for styling
        button.onclick = function () {
          const jobTitle = card.innerText;
          const company = card
            .closest(".artdeco-entity-lockup__content")
            .querySelector(
              ".job-card-container__primary-description"
            ).innerText;
          const location = card
            .closest(".artdeco-entity-lockup__content")
            .querySelector(".job-card-container__metadata-item span").innerText;

          const jobDetails = {
            JobId: jobId,
            Title: jobTitle,
            Company: company,
            Location: location,
            URL:
              window.location.origin +
              card
                .closest("[data-job-id]")
                .querySelector("a")
                .getAttribute("href"),
          };

          console.log(JSON.stringify(jobDetails, null, 2));
        };

        fragment.appendChild(button); // Append the button to the fragment
        card.appendChild(fragment); // Append the fragment to the card

        // Mark this job ID as processed
        processedJobIds.add(jobId);
      }
    }
  });
}

// Function to observe changes in the job listings with throttling
function observeJobListings(mutationsList) {
  // If throttle timeout is active, do not proceed
  if (throttleTimeout) {
    return;
  }

  throttleTimeout = setTimeout(() => {
    requestAnimationFrame(() => {
      // Only proceed if not currently processing
      if (isProcessing) return;

      isProcessing = true; // Set processing flag

      // Call the function to add buttons
      addJobButtons();
      hidePromotedJobs();
      hideAppliedJobs();

      // After adding buttons, re-check for any new job cards added during processing
      const jobCards = document.querySelectorAll(".job-card-list__title");
      jobCards.forEach((card) => {
        const jobId = card.closest("[data-job-id]").getAttribute("data-job-id");
        // If it's a new job, add it to the processed set and add the button
        if (card.innerText.trim() && !processedJobIds.has(jobId)) {
          processedJobIds.add(jobId); // Mark as processed
        }
      });

      isProcessing = false; // Reset processing flag
    });
    throttleTimeout = null; // Reset the throttle timeout
  }, 500); // 2 seconds delay
}

// Initial setup function
function initialize() {
  hidePromotedJobs();
  hideAppliedJobs();
  addJobButtons();

  const jobListContainer = document.querySelector(
    ".scaffold-layout__list-container"
  );
  if (jobListContainer) {
    const observer = new MutationObserver(observeJobListings);
    observer.observe(jobListContainer, { childList: true, subtree: true });

    // Clean up observer when no longer needed
    window.addEventListener("unload", () => {
      observer.disconnect();
    });
  }
}

// Wait for the page to fully load before running the script
window.addEventListener("load", initialize);
