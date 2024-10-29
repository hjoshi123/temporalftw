# TemporalFTW

TemporalFTW is a demo project for the Denver DevOps meetup. This project simulates the use of temporal through a simple banking 
application where users can create accounts, deposit and withdraw money. This sample project demonstrates the following concepts of temporal:
* Workflow - Parent and Child workflows
* Activity - Local and Activity
* Signals
* Query

## Prerequisites

* Go 1.22 or later
* Docker
* Docker Compose
* Enough resources to run the Temporal service and the sample application
* AtlasGo binary for migrations. Run `curl -sSf https://atlasgo.sh | sh -s -- -y` to install it.

## Running the project

Ideally, running the project is as simple as running `docker compose --env-file .env up -d --build` in the root directory of the project. 
This will start the Temporal service and the sample application. But before running the docker compose we need to create an env file, so that
the application and worker can read the env configurations. The configuration definition is located in `internal/config/configuration.go`. 
The `.env` file should be like the one shown below:
```shell
DB_PORT=5432
DB_HOST=postgres
DB_NAME=temporalLoan
DB_USER=dev
DB_PASSWORD=dev124342
LOG_LEVEL=DEBUG
ENVIRONMENT=development
PORT=8081
TEMPORAL_HOST_PORT=temporal:7233
TEMPORAL_HOST=temporal
TEMPORAL_PORT=7233
```

## Building the project

To build the project and contribute to it along with go you would also need `sqlboiler` and `atlas` installed. These run the migrations and
create the models in `pkg/models/postgres`. To build the project, run the following commands:
```shell
git clone git@github.com:hjoshi123/temporalftw.git && cd temporalftw
go mod download

# Create migration file
atlas migrate new --edit <migration_name>

# Apply the migration
atlas migrate apply -u postgres://dev:dev12342@localhost:5432/temporalLoan?sslmode=disable

# Create models. Configure the database connection in the sqlboiler.toml file
sqlboiler psql
```

## Architecture

Before diving into temporal, let's briefly see the directory structure of the project.

* `cmd` - Contains the main entry point of the application and the temporal worker.
* `internal` - Contains the internal logic of the application like logging, db connections which are not exposed
* `pkg` - Contains the public logic of the application like the API handlers, Temporal workflows, activities, etc.
* `migrations` - Contains the database migration scripts to run on any environment or machine where the application is deployed.

### Temporal

Both temporal workflows and activities are located in the `pkg/workflows` and `pkg/activities` directory. There are 3 workflows used 
out of which the main parent workflow is `TransactionWorkflow` which is responsible for creating child workflows `TransactionApproveWorkflow`
and `TransactionRejectWorkflow`.
These child workflows are triggered through the use of signals from the parent workflow. Once someone manually approves or rejects the transaction,
it signals the workflow (by using workflowID). 