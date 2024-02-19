import asyncio
import logging
from datetime import timedelta

from temporalio import workflow
from temporalio.client import Client
from temporalio.worker import Worker

with workflow.unsafe.imports_passed_through():
   from activities import (SampleModelInput, TrainModelInput, sample_model,
                           train_model)

async def main():
  logging.basicConfig(level=logging.INFO)

  client = await Client.connect("localhost:7233")

  worker = Worker(
      client,
      task_queue="ml",
      activities=[train_model, sample_model],
  )
  logging.info("running python worker")
  await worker.run()

if __name__ == "__main__":
    asyncio.run(main())
