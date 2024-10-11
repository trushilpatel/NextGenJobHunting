import logging
import os
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from dotenv import load_dotenv
from models import crawled_job
from sqlalchemy.dialects.postgresql import insert

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
            # Create an insert statement with an on_conflict_do_nothing clause
            insert_stmt = insert(crawled_job.CrawledJob).values(
                job_id=job_id, job_data=job_data, platform_url=platform_url, status=status
            ).on_conflict_do_nothing(index_elements=['job_id'])

            # Execute the statement
            self.session.execute(insert_stmt)
            self.session.commit()
            logger.info(f"Job {job_id} added successfully or already exists.")
        except Exception as e:
            self.session.rollback()
            logger.error(f"Failed to add job {job_id}. Error: {e}")
        finally:
            self.session.close()


    def get_crawled_job_by_id(self, job_id):
        try:
            job = self.session.query(crawled_job.CrawledJob).filter_by(job_id=job_id).first()
            if job:
                return job
            else:
                return None
        except Exception as e:
            logger.error(f"Failed to retrieve job {job_id}. Error: {e}")
            return None
        finally:
            self.session.close()