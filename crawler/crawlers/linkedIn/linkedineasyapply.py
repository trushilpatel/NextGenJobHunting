import time
import random
from urllib.parse import urlparse
import pyautogui
import traceback
from selenium.common.exceptions import TimeoutException
from selenium.common.exceptions import StaleElementReferenceException
from selenium.webdriver.common.by import By

from itertools import product
import logging

from services.crawled_job_service import CrawledJobService
import os

logger = logging.getLogger(__name__)
# Set up logging
log_directory = os.path.join(os.getcwd(), "logs")
os.makedirs(log_directory, exist_ok=True)
log_file = os.path.join(log_directory, "linkedin.log")

logging.basicConfig(
    format="%(asctime)s %(levelname)-8s %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S",
    filename=log_file,
    level=logging.INFO,
)


class LinkedinEasyApply:
    def __init__(self, parameters, driver):
        self.crawled_job_service = CrawledJobService()
        self.browser = driver
        self.email = parameters["email"]
        self.password = parameters["password"]
        self.disable_lock = parameters["disableAntiLock"]
        self.company_blacklist = parameters.get("companyBlacklist", []) or []
        self.title_blacklist = parameters.get("titleBlacklist", []) or []
        self.poster_blacklist = parameters.get("posterBlacklist", []) or []
        self.positions = parameters.get("positions", [])
        self.locations = parameters.get("locations", [])
        self.residency = parameters.get("residentStatus", [])
        self.base_search_url = self.get_base_search_url(parameters)
        self.seen_jobs = []
        self.file_name = "../output_"
        self.unprepared_questions_file_name = "../unprepared_questions"
        self.output_file_directory = parameters["outputFileDirectory"]
        self.resume_dir = parameters["uploads"]["resume"]
        if "coverLetter" in parameters["uploads"]:
            self.cover_letter_dir = parameters["uploads"]["coverLetter"]
        else:
            self.cover_letter_dir = ""
        self.checkboxes = parameters.get("checkboxes", [])
        self.university_gpa = parameters["universityGpa"]
        self.salary_minimum = parameters["salaryMinimum"]
        self.notice_period = int(parameters["noticePeriod"])
        self.languages = parameters.get("languages", [])
        self.experience = parameters.get("experience", [])
        self.personal_info = parameters.get("personalInfo", [])
        self.eeo = parameters.get("eeo", [])
        self.experience_default = int(self.experience["default"])

    def login(self):
        try:
            self.browser.get("https://www.linkedin.com")
            time.sleep(random.uniform(5, 10))
            # self.browser.find_element(By.ID, "username").send_keys(self.email)
            # self.browser.find_element(By.ID, "password").send_keys(self.password)
            # self.browser.find_element(By.CSS_SELECTOR, ".btn__primary--large").click()
            time.sleep(random.uniform(5, 10))
        except TimeoutException:
            raise Exception("Could not login!")

    def security_check(self):
        current_url = self.browser.current_url
        page_source = self.browser.page_source

        if "/checkpoint/challenge/" in current_url or "security check" in page_source:
            input(
                "Please complete the security check and press enter on this console when it is done."
            )
            time.sleep(random.uniform(5.5, 10.5))

    def start_applying(self):
        searches = list(product(self.positions, self.locations))
        random.shuffle(searches)

        page_sleep = 0
        minimum_time = 60 * 15  # minimum time bot should run before taking a break
        minimum_page_time = time.time() + minimum_time

        for position, location in searches:
            location_url = "&location=" + location
            job_page_number = -1

            logger.info("Starting the search for " + position + " in " + location + ".")

            try:
                while True:
                    page_sleep += 1
                    job_page_number += 1
                    logger.info("Going to job page " + str(job_page_number))
                    self.next_job_page(position, location_url, job_page_number)
                    time.sleep(random.uniform(1.5, 3.5))
                    logger.info("Starting the application process for this page...")
                    self.apply_jobs(location)
                    logger.info(
                        "Job applications on this page have been successfully completed."
                    )

                    time_left = minimum_page_time - time.time()
                    if time_left > 0:
                        logger.info("Sleeping for " + str(time_left) + " seconds.")
                        time.sleep(time_left)
                        minimum_page_time = time.time() + minimum_time
                    if page_sleep % 5 == 0:
                        sleep_time = random.randint(
                            60, 90
                        )  # Changed from 500, 900 {seconds}
                        logger.info("Sleeping for " + str(sleep_time / 60) + " minutes.")
                        time.sleep(sleep_time)
                        page_sleep += 1
            except:
                traceback.print_exc()
                pass

            time_left = minimum_page_time - time.time()
            if time_left > 0:
                logger.info("Sleeping for " + str(time_left) + " seconds.")
                time.sleep(time_left)
                minimum_page_time = time.time() + minimum_time
            if page_sleep % 5 == 0:
                sleep_time = random.randint(500, 900)
                logger.info("Sleeping for " + str(sleep_time / 60) + " minutes.")
                time.sleep(sleep_time)
                page_sleep += 1

    def apply_jobs(self, location):
        no_jobs_text = ""
        try:
            no_jobs_element = self.browser.find_element(
                By.CLASS_NAME, "jobs-search-two-pane__no-results-banner--expand"
            )
            no_jobs_text = no_jobs_element.text
        except:
            pass
        if "No matching jobs found" in no_jobs_text:
            raise Exception("No more jobs on this page.")

        if "unfortunately, things are" in self.browser.page_source.lower():
            raise Exception("No more jobs on this page.")

        job_results_header = ""
        maybe_jobs_crap = ""
        job_results_header = self.browser.find_element(
            By.CLASS_NAME, "jobs-search-results-list__text"
        )
        maybe_jobs_crap = job_results_header.text

        if "Jobs you may be interested in" in maybe_jobs_crap:
            raise Exception("Nothing to do here, moving forward...")

        try:
            job_results = self.browser.find_element(
                By.CLASS_NAME, "jobs-search-results-list"
            )
            self.scroll_slow(job_results)
            self.scroll_slow(job_results, step=300, reverse=True)

            job_list = self.browser.find_elements(
                By.CLASS_NAME, "scaffold-layout__list-container"
            )[0].find_elements(By.CLASS_NAME, "jobs-search-results__list-item")
            if len(job_list) == 0:
                raise Exception("No job class elements found in page")
        except Exception as e:
            logger.error(f"Error getting job list: {e}")
            raise Exception("No more jobs on this page.")

        if len(job_list) == 0:
            raise Exception("No more jobs on this page.")

        for job_tile in job_list:
            job_title, company, link = ("", "", "")

            try:
                # patch to incorporate new 'verification' crap by LinkedIn
                # job_title = job_tile.find_element(By.CLASS_NAME, 'job-card-list__title').text # original code
                job_title_element = job_tile.find_element(
                    By.CLASS_NAME, "job-card-list__title"
                )
                job_title = job_title_element.find_element(By.TAG_NAME, "strong").text

                link = (
                    job_tile.find_element(By.CLASS_NAME, "job-card-list__title")
                    .get_attribute("href")
                    .split("?")[0]
                )
                job_id = None
                try:
                    job_id = int(link.rstrip("/").split("/")[-1])
                except ValueError:
                    logger.error(f"Invalid job ID in link: {link}")
                    continue
            except:
                pass
            try:
                company = job_tile.find_element(
                    By.CLASS_NAME, "job-card-container__primary-description"
                ).text
            except:
                pass

            max_retries = 3
            retries = 0
            while retries < max_retries:
                try:
                    job_el = job_tile.find_element(
                        By.CLASS_NAME, "job-card-list__title"
                    )
                    job_el.click()
                    break
                except StaleElementReferenceException:
                    retries += 1
                    continue
                except Exception as e:
                    break

            time.sleep(random.uniform(5, 10))

            try:
                # scroll job details to the bottom and the to the up
                self.scroll_element("jobs-search__job-details--container", end=1600)
                self.scroll_element(
                    "jobs-search__job-details--container",
                    end=1600,
                    step=400,
                    reverse=True,
                )

                # Get the current website's base URL
                current_url = self.browser.current_url
                base_url = "{0.scheme}://{0.netloc}".format(urlparse(current_url))

                jobdetails = self.get_job_details()

                self.crawled_job_service.add_crawled_job(
                    job_id=job_id,
                    job_data=str(jobdetails),
                    platform_url=link,
                )
            except Exception as e:
                logger.info(e)
                logger.info(f"Could not apply to the job in {company}")
                traceback.print_exc()

            self.seen_jobs += link

    def scroll_element(self, element_name, end=1600, step=100, reverse=False):
        try:
            element = self.browser.find_element(By.CLASS_NAME, element_name)
            self.scroll_slow(element, end=end, step=step, reverse=reverse)
        except Exception as e:
            logger.error(f"Error scrolling element {element_name}: {e}")
            pass

    def scroll_slow(
        self, scrollable_element, start=0, end=3600, step=100, reverse=False
    ):
        if reverse:
            start, end = end, start
            step = -step

        for i in range(start, end, step):
            self.browser.execute_script(
                "arguments[0].scrollTo(0, {})".format(i), scrollable_element
            )
            time.sleep(random.uniform(1.0, 2.6))

    def avoid_lock(self):
        if self.disable_lock:
            return

        pyautogui.keyDown("ctrl")
        pyautogui.press("esc")
        pyautogui.keyUp("ctrl")
        time.sleep(1.0)
        pyautogui.press("esc")

    def get_base_search_url(self, parameters):
        remote_url = ""
        lessthanTenApplicants_url = ""

        if parameters.get("remote"):
            remote_url = "&f_WT=2"
        else:
            remote_url = ""
            # TO DO: Others &f_WT= options { WT=1 onsite, WT=2 remote, WT=3 hybrid, f_WT=1%2C2%2C3 }

        if parameters["lessthanTenApplicants"]:
            lessthanTenApplicants_url = "&f_EA=true"

        level = 1
        experience_level = parameters.get("experienceLevel", [])
        experience_url = "f_E="
        for key in experience_level.keys():
            if experience_level[key]:
                experience_url += "%2C" + str(level)
            level += 1

        distance_url = "?distance=" + str(parameters["distance"])

        job_types_url = "f_JT="
        job_types = parameters.get("jobTypes", [])
        # job_types = parameters.get('experienceLevel', [])
        for key in job_types:
            if job_types[key]:
                job_types_url += "%2C" + key[0].upper()

        date_url = ""
        dates = {
            "all time": "",
            "month": "&f_TPR=r2592000",
            "week": "&f_TPR=r604800",
            "24 hours": "&f_TPR=r86400",
        }
        date_table = parameters.get("date", [])
        for key in date_table.keys():
            if date_table[key]:
                date_url = dates[key]
                break

        easy_apply_url = ""

        extra_search_terms = [
            distance_url,
            remote_url,
            lessthanTenApplicants_url,
            job_types_url,
            experience_url,
        ]
        extra_search_terms_str = (
            "&".join(term for term in extra_search_terms if len(term) > 0)
            + easy_apply_url
            + date_url
        )

        return extra_search_terms_str

    def next_job_page(self, position, location, job_page):
        self.browser.get(
            "https://www.linkedin.com/jobs/search/"
            + self.base_search_url
            + "&keywords="
            + position
            + location
            + "&start="
            + str(job_page * 25)
        )

        self.avoid_lock()

    def get_inner_text(self, selector, by=By.CLASS_NAME):
        """Helper function to get the inner text of an element with error handling."""
        try:
            element = self.browser.find_element(by, selector)
            return element.get_attribute("innerText").strip()
        except Exception as e:
            # traceback.print_exc()
            return None  # Return None if the element is not found

    def get_job_details(self):
        """Retrieve and print job details."""
        # Define selectors
        # all_job_details_selector = "jobs-search__job-details--wrapper"
        company_name_selector = "job-details-jobs-unified-top-card__company-name"
        primary_description_selector = (
            "job-details-jobs-unified-top-card__primary-description-container"
        )
        hirer_info_selector = "hirer-card__hirer-information"
        hirer_link_selector = "div.hirer-card__hirer-information a.app-aware-link"
        job_insight_selector = "job-details-jobs-unified-top-card__job-insight"
        job_description_selector = "jobs-description-content__text"
        company_description_selector = "jobs-company__company-description"
        applicants_selector = 'div[data-view-name="premium-job-applicant-insights"]'

        # Get text details using the helper function
        # all_job_details_text = self.get_inner_text(
        #     all_job_details_selector, By.CLASS_NAME
        # )
        company_name = self.get_inner_text(company_name_selector, By.CLASS_NAME)
        primary_description_text = self.get_inner_text(
            primary_description_selector, By.CLASS_NAME
        )
        hirer_name_text = self.get_inner_text(hirer_info_selector, By.CLASS_NAME)
        linkedin_profile_link = self.get_inner_text(
            hirer_link_selector, By.CSS_SELECTOR
        )
        job_insight_text = self.get_inner_text(job_insight_selector, By.CLASS_NAME)
        job_description_text = self.get_inner_text(
            job_description_selector, By.CLASS_NAME
        )
        company_description_text = self.get_inner_text(
            company_description_selector, By.CLASS_NAME
        )
        applicants_text = self.get_inner_text(applicants_selector, By.CSS_SELECTOR)

        # Print the collected information
        # print(f"All Job Details: {all_job_details_text}")

        job_details = {
            "company_name": company_name,
            "primary_description": primary_description_text,
            "hirer_name": hirer_name_text,
            "linkedin_profile_link": linkedin_profile_link,
            "job_insight": job_insight_text,
            "job_description": job_description_text,
            "company_description": company_description_text,
            "applicants": applicants_text,
        }
        return job_details
