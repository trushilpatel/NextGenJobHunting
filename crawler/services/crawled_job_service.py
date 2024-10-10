import os
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from dotenv import load_dotenv
from models import crawled_job


class CrawledJobService:
    def __init__(self):
        # Load environment variables
        load_dotenv()

        self.CRAWLER_DB_HOST = os.getenv("CRAWLER_DB_HOST")
        self.CRAWLER_DB_PORT = os.getenv("CRAWLER_DB_PORT")
        self.CRAWLER_DB_USERNAME = os.getenv("CRAWLER_DB_USERNAME")
        self.CRAWLER_DB_PASSWORD = os.getenv("CRAWLER_DB_PASSWORD")
        self.CRAWLER_DB_NAME = os.getenv("CRAWLER_DB_NAME")
        self.CRAWLER_DB_DRIVER = os.getenv("CRAWLER_DB_DRIVER")
        self.DATABASE_URL = f"{self.CRAWLER_DB_DRIVER}://{self.CRAWLER_DB_USERNAME}:{self.CRAWLER_DB_PASSWORD}@{self.CRAWLER_DB_HOST}:{self.CRAWLER_DB_PORT}/{self.CRAWLER_DB_NAME}"

        self.check_for_missing_env_variables()
        # Create a database engine
        engine = create_engine(self.DATABASE_URL)
        Session = sessionmaker(bind=engine)
        self.session = Session()

    def check_for_missing_env_variables(self):
        # Check for missing environment variables
        required_vars = {
            "CRAWLER_DB_HOST": self.CRAWLER_DB_HOST,
            "CRAWLER_DB_PORT": self.CRAWLER_DB_PORT,
            "CRAWLER_DB_USERNAME": self.CRAWLER_DB_USERNAME,
            "CRAWLER_DB_PASSWORD": self.CRAWLER_DB_PASSWORD,
            "CRAWLER_DB_NAME": self.CRAWLER_DB_NAME,
            "CRAWLER_DB_DRIVER": self.CRAWLER_DB_DRIVER,
        }

        for var_name, var_value in required_vars.items():
            if not var_value:
                raise ValueError(
                    f"Environment variable {var_name} is missing or empty."
                )
    
    def add_crawled_job(self, job_id, job_data, platform_url, status="new"):
        try:
            new_job = crawled_job.CrawledJob(
                job_id=job_id, job_data=job_data, platform_url=platform_url, status=status
            )
            self.session.add(new_job)
            self.session.commit()
            print(f"Job {job_id} added successfully.")
        except Exception as e:
            self.session.rollback()
            print(f"Failed to add job {job_id}. Error: {e}")
        finally:
            self.session.close()
