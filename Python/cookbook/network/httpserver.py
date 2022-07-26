import asyncio

from fastapi import FastAPI
import uvicorn

app = FastAPI()


@app.get("/hello")
async def hello():
    for i in range(100):
        await asyncio.sleep(1)
        print(f"hello {i}")

    return "ok"


def main():
    uvicorn.run(
        app="cookbook.network.httpserver:app", host="0.0.0.0", port=8000, reload=True
    )


if __name__ == "__main__":
    main()
