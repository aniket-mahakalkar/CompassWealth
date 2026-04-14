# Deployment Guide to Google Cloud Run

This guide will walk you through building a Docker image for your application and deploying it to Google Cloud Run.

## Prerequisites

1.  **Google Cloud Account**: Ensure you have a GCP account.
2.  **gcloud CLI**: Install and initialize the [Google Cloud SDK](https://cloud.google.com/sdk/docs/install).
3.  **Docker**: Ensure Docker is installed and running on your machine (alternatively, use Cloud Build).

## Step 1: Configure gcloud

Run the following commands to set up your environment:

```powershell
# Login to Google Cloud
gcloud auth login

# Set your project ID
gcloud config set project [YOUR_PROJECT_ID]

# Enable required services
gcloud services enable run.googleapis.com containerregistry.googleapis.com cloudbuild.googleapis.com
```

## Step 2: Build and Deploy using Cloud Build (Recommended)

Cloud Build allows you to build the Docker image in the cloud without needing Docker installed locally.

```powershell
# Build the image and deploy to Cloud Run
gcloud builds submit --tag gcr.io/[YOUR_PROJECT_ID]/compass-wealth

# Deploy to Cloud Run
gcloud run deploy compass-wealth `
  --image gcr.io/[YOUR_PROJECT_ID]/compass-wealth `
  --platform managed `
  --region us-central1 `
  --allow-unauthenticated
```

## Step 3: Local Build (Optional)

If you prefer to build locally and then push:

```powershell
# Build locally
docker build -t gcr.io/[YOUR_PROJECT_ID]/compass-wealth .

# Push to Artifact Registry/Container Registry
docker push gcr.io/[YOUR_PROJECT_ID]/compass-wealth

# Deploy
gcloud run deploy compass-wealth --image gcr.io/[YOUR_PROJECT_ID]/compass-wealth --region us-central1
```

## Important Notes

### Automatic Migrations
The server is configured to automatically run database migrations on startup. Every time you deploy or the container restarts, it will ensure that the tables (like `Account`) are created correctly.

### Persistence
This project currently uses **SQLite** (`compass_wealth.db`). 

> [!WARNING]
> Cloud Run is **stateless** and its filesystem is ephemeral. This means any changes made to the `compass_wealth.db` file will be **lost** when the container restarts.

For a production environment, you should consider:
1.  **Cloud SQL**: Use a managed database like PostgreSQL or MySQL.
2.  **Litestream**: If you want to stick with SQLite, use [Litestream](https://litestream.io/) to replicate the database to Google Cloud Storage.
