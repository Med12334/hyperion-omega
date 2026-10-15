from cortex.autoscaler import PredictiveAutoscaler
from cortex.rca_debugger import CausalRcaDebugger

def test_predictive_autoscaler_prewarms_gpus():
    scaler = PredictiveAutoscaler()
    res = scaler.forecast_workload(1200.0, 4)
    assert res["autoscaler_action"] == "PREWARM_GPU_NODES"
    assert res["recommended_gpu_allocation"] == 6

def test_rca_debugger_detects_oom():
    debugger = CausalRcaDebugger()
    res = debugger.analyze_failure_cascade(["Container SIGKILL OOMKilled encountered on pod-03"])
    assert res["severity"] == "CRITICAL"
    assert "OOM" in res["root_cause_summary"]
