// content.js

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

  jobCards.forEach((card) => {
    // Check if the job title exists
    if (card.innerText.trim()) {
      // Check if the button already exists
      if (!card.querySelector(".custom-button")) {
        const button = document.createElement("button");
        button.innerText = "Get Job Details";
        button.className = "custom-button"; // Add a custom class for styling
        button.onclick = function () {
          const jobId = card
            .closest("[data-job-id]")
            .getAttribute("data-job-id");
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

        // Append the button to the job card
        card.appendChild(button);
      }
    }
  });
}

// Function to observe changes in the job listings
function observeJobListings(mutationsList) {
  // Check for added nodes
  mutationsList.forEach((mutation) => {
    if (mutation.type === "childList" && mutation.addedNodes.length > 0) {
      addJobButtons();
      hidePromotedJobs();
      hideAppliedJobs();
    }
  });
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
  }
}

// Wait for the page to fully load before running the script
window.addEventListener("load", initialize);
