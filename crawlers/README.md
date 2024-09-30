## Crawler Processing

#### Navigate to the crawler directory

```sh
cd ./
```

#### Create a virtual environment

```sh
python -m venv venv
```

#### Activate the virtual environment

**On macOS/Linux**

```sh
source venv/bin/activate
```

**On Windows**

```sh
venv\Scripts\activate
```

#### Install dependencies from requirements.txt

```sh
pip install -r ../requirements.txt
```

## Setup Crawler database

#### Setup Alembic

1. Initialize Alembic in your project:

   ```sh
   alembic init alembic
   ```

2. Configure your `alembic.ini` file and `env.py` to connect to your database.

3. Create a new migration script:

   ```sh
   alembic revision --autogenerate -m "Initial migration"
   ```

4. Apply the migration to your database:

   ```sh
   alembic upgrade head
   ```

#### Running Alembic

To run Alembic migrations, use the following command:

```sh
alembic upgrade head
```

This command will apply all pending migrations to your database.
