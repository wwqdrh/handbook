"""modify account table

Revision ID: 90a28d4b85da
Revises: 3baee1caec30
Create Date: 2020-12-21 13:26:48.332788

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = '90a28d4b85da'
down_revision = '3baee1caec30'
branch_labels = None
depends_on = None


def upgrade():
    op.add_column('account', sa.Column('last_transaction_date', sa.DateTime))


def downgrade():
    op.drop_column('account', 'last_transaction_date')
