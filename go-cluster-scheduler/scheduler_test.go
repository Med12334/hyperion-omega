package main

import (
"bytes"
"encoding/json"
"net/http"
"net/http/httptest"
"testing"
)

func TestScheduleHandler(t *testing.T) {
scheduler := NewClusterScheduler()
server := httptest.NewServer(http.HandlerFunc(scheduler.ScheduleHandler))
defer server.Close()

workload := ComputeWorkload{
JobId:         "job-ai-train-991",
TenantId:      "tenant-enterprise-01",
VcpuRequested: 32,
GpuRequired:   true,
Region:        "eu-west-1",
}

body, _ := json.Marshal(workload)
resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(body))
if err != nil {
t.Fatalf("Failed to post schedule request: %v", err)
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
t.Errorf("Expected status 200, got %d", resp.StatusCode)
}

var alloc NodeAllocation
json.NewDecoder(resp.Body).Decode(&alloc)
if alloc.AssignedJob != "job-ai-train-991" {
t.Errorf("Expected assigned job 'job-ai-train-991', got %s", alloc.AssignedJob)
}
}
