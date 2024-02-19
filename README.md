# Neosync Temporal Feb 24 Meetup Demos

## General Setup

1. Run `docker compose up -d` to stand up the Temporal suite of services as well as two postgres databases that will be used as our "prod" and "stage" data sources.
2. See the [data gen readme](./data-gen/README.md) for info on how to run this demo.
3. See the [ml data gen readme](./ml-data-gen/README.md) for info on how to run this demo.

## Data Gen Demo

This demo will generate 100 rows of random data and insert it into a `public.users` table in a postgres database.

## ML Data Gen Demo

This second demo will take the data we created in the first demo and train a machine learning model on it.
From there, we can sample that model to generate new data that is statistically consistent with the source input.
