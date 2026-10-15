from fastapi import FastAPI
from pydantic import BaseModel
from typing import List
from cortex.autoscaler import PredictiveAutoscaler
from cortex.rca_debugger import CausalRcaDebugger

app = FastAPI(title="Hyperion-Omega AI Cortex & Causal RCA Plane", version="1.0.0")

autoscaler = PredictiveAutoscaler()
rca = CausalRcaDebugger()

class ForecastRequest(BaseModel):
    current_rps: float
    active_gpu_jobs: int

class RcaRequest(BaseModel):
    error_logs: List[str]

@app.get("/health")
def health_check():
    return {"status": "UP", "engine": "Python-AI-Cortex-v1"}

@app.post("/api/v1/ai/forecast")
def forecast(req: ForecastRequest):
    return autoscaler.forecast_workload(req.current_rps, req.active_gpu_jobs)

@app.post("/api/v1/ai/rca")
def analyze_rca(req: RcaRequest):
    return rca.analyze_failure_cascade(req.error_logs)
