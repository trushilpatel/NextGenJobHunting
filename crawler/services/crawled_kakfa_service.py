import logging
from kafka import KafkaProducer
import json
import os

logger = logging.getLogger(__name__)
# Set up logging
log_directory = os.path.join(os.getcwd(), "logs")
os.makedirs(log_directory, exist_ok=True)
log_file = os.path.join(log_directory, "kafka.log")

logging.basicConfig(
    format="%(asctime)s %(levelname)-8s %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S",
    filename=log_file,
    level=logging.INFO,
)

class CrawledKafkaService:
    def __init__(self):
        # Kafka configuration from environment variables
        self.KAFKA_HOST = os.getenv("KAFKA_HOST", "localhost")
        self.KAFKA_PORT = os.getenv("KAFKA_PORT", "9092")
        self.KAFKA_TOPIC = os.getenv("KAFKA_TOPIC", "crawled_job")

        # Combine host and port
        self.KAFKA_SERVER = f"{self.KAFKA_HOST}:{self.KAFKA_PORT}"

        # Initialize Kafka producer
        self.producer = KafkaProducer(
            bootstrap_servers=self.KAFKA_SERVER,
            value_serializer=lambda v: json.dumps(v).encode('utf-8')  # Serialize JSON
        )

    def send_crawled_job(self, crawled_job):
        try:
            self.producer.send(self.KAFKA_TOPIC, crawled_job)
            self.producer.flush()  # Ensure the message is sent
            logger.info(f"Message sent to topic {self.KAFKA_TOPIC}")

        except Exception as e:
            logger.error(f"Failed to send message: {e}")


if __name__ == "__main__":
    service = CrawledKafkaService()
    test_job = {
        "title": "Software Engineer",
        "company": "Tech Company",
        "location": "Remote",
        "description": "An exciting role in a dynamic company.",
        "date_posted": "2023-10-01"
    }
    service.send_crawled_job(test_job)
    logger.info("Test job sent to Kafka successfully.")