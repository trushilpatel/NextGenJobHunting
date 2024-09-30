from sqlalchemy import Column, Integer, String, Text, TIMESTAMP, func
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class CrawledJob(Base):
    __tablename__ = "crawled_job"

    id = Column(Integer, primary_key=True)
    job_id = Column(String(255), nullable=False)
    job_data = Column(Text, nullable=False)
    platform_url = Column(String(255), nullable=False)
    status = Column(String(50), nullable=False, default="new")
    created_at = Column(TIMESTAMP, server_default=func.now())
