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
from sqlalchemy.dialects import postgresql


def upgrade() -> None:
    # Create the ENUM type for status if it doesn't exist
    # Check if the table exists before dropping it
    conn = op.get_bind()
    if conn.dialect.has_table(conn, "crawled_jobs"):
        op.drop_table("crawled_jobs")

    status_enum = postgresql.ENUM("new", "in_queue", "completed", name="status_enum")
    if conn.dialect.has_type(conn, "status_enum"):
        status_enum.drop(conn, checkfirst=True)

    # Create the crawled_jobs table
    op.create_table(
        "crawled_job",
        sa.Column("id", sa.Integer, primary_key=True, autoincrement=True),
        sa.Column("job_id", sa.String(length=255), nullable=False),
        sa.Column("job_data", sa.Text(), nullable=False),
        sa.Column("platform_url", sa.String(length=255), nullable=False),
        sa.Column("status", status_enum, nullable=False, server_default="new"),
        sa.Column("created_at", sa.TIMESTAMP(), server_default=sa.func.now()),
    )


def downgrade() -> None:
    # Drop the crawled_jobs table
    op.drop_table("crawled_job")

    # Check if the ENUM type exists before dropping it
    conn = op.get_bind()
    if conn.dialect.has_type(conn, "status_enum"):
        status_enum = postgresql.ENUM(
            "new", "in_queue", "completed", name="status_enum"
        )
        status_enum.drop(op.get_bind(), checkfirst=True)
