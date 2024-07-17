
from fastapi import FastAPI
import uvicorn

app = FastAPI()

@app.get("/one")
def foo():
    return {"message": "hello"}


if __name__ == '__main__':
	uvicorn.run(app, port=8081, host="0.0.0.0")
