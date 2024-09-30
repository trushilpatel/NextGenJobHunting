import os
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from dotenv import load_dotenv
from models import CrawledJobs, Base

# Load environment variables
load_dotenv()

CRAWLER_DB_HOST = os.getenv("CRAWLER_DB_HOST")
CRAWLER_DB_PORT = os.getenv("CRAWLER_DB_PORT")
CRAWLER_DB_USERNAME = os.getenv("CRAWLER_DB_USERNAME")
CRAWLER_DB_PASSWORD = os.getenv("CRAWLER_DB_PASSWORD")
CRAWLER_DB_NAME = os.getenv("CRAWLER_DB_NAME")
CRAWLER_DB_DRIVER = os.getenv("CRAWLER_DB_DRIVER")
DATABASE_URL = f"{CRAWLER_DB_DRIVER}://{CRAWLER_DB_USERNAME}:{CRAWLER_DB_PASSWORD}@{CRAWLER_DB_HOST}:{CRAWLER_DB_PORT}/{CRAWLER_DB_NAME}"

# Create a database engine
engine = create_engine(DATABASE_URL)
Session = sessionmaker(bind=engine)


class DBService:
    def __init__(self):
        self.session = Session()

    def add_crawled_job(self, job_id, job_data, platform_url):
        new_job = CrawledJobs(
            job_id=job_id, job_data=job_data, platform_url=platform_url
        )
        self.session.add(new_job)
        self.session.commit()
        print(f"Job {job_id} added successfully.")
        self.session.close()


def check_for_missing_env_variables():
    # Check for missing environment variables
    required_vars = {
        "CRAWLER_DB_HOST": CRAWLER_DB_HOST,
        "CRAWLER_DB_PORT": CRAWLER_DB_PORT,
        "CRAWLER_DB_USERNAME": CRAWLER_DB_USERNAME,
        "CRAWLER_DB_PASSWORD": CRAWLER_DB_PASSWORD,
        "CRAWLER_DB_NAME": CRAWLER_DB_NAME,
        "CRAWLER_DB_DRIVER": CRAWLER_DB_DRIVER,
    }

    for var_name, var_value in required_vars.items():
        if not var_value:
            raise ValueError(f"Environment variable {var_name} is missing or empty.")
