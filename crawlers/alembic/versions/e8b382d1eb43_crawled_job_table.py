"""crawled_job table

Revision ID: e8b382d1eb43
Revises: 
Create Date: 2024-09-30 13:48:18.807805

"""

from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = "e8b382d1eb43"
down_revision: Union[str, None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade():
    # Create the Crawled_jobs table
    op.execute(
        """
    CREATE TABLE Crawled_jobs (
        id SERIAL PRIMARY KEY,
        job_id VARCHAR(255) NOT NULL,
        job_data TEXT NOT NULL,
        platform_url VARCHAR(255) NOT NULL,
        status VARCHAR(50) NOT NULL DEFAULT 'new',
        created_at TIMESTAMP DEFAULT NOW()
    );
    """
    )


def downgrade():
    # Drop the Crawled_jobs table
    op.execute("DROP TABLE IF EXISTS Crawled_jobs;")
