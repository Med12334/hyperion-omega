package main

import (
"encoding/json"
"fmt"
"net/http"
"sync"
"time"
)

type ComputeWorkload struct {
JobId         string `json:job_id`
TenantId      string `json:tenant_id`
VcpuRequested int    `json:vcpu_requested`
GpuRequired   bool   `json:gpu_required`
Region        string `json:region`
}

type NodeAllocation struct {
NodeId       string    `json:node_id`
AssignedJob  string    `json:assigned_job`
ScheduledAt  time.Time `json:scheduled_at`
}

type ClusterScheduler struct {
mu          sync.Mutex
allocations []NodeAllocation
nodePool    []string
}

func NewClusterScheduler() *ClusterScheduler {
return &ClusterScheduler{
nodePool: []string{"node-eu-west-gpu-01", "node-us-east-cpu-02", "node-ap-south-flex-03"},
}
}

func (s *ClusterScheduler) ScheduleHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
return
}

var workload ComputeWorkload
if err := json.NewDecoder(r.Body).Decode(&workload); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

s.mu.Lock()
defer s.mu.Unlock()

targetNode := s.nodePool[len(s.allocations)%len(s.nodePool)]
alloc := NodeAllocation{
NodeId:      targetNode,
AssignedJob: workload.JobId,
ScheduledAt: time.Now(),
}
s.allocations = append(s.allocations, alloc)

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(alloc)
}

func main() {
scheduler := NewClusterScheduler()
http.HandleFunc("/api/v1/schedule", scheduler.ScheduleHandler)
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte(`{"status":"UP","scheduler":"Go-Raft-Scheduler"}`))
})

fmt.Println("[GO SCHEDULER] Layer 4 & 5 Cluster Scheduler listening on port 8090...")
http.ListenAndServe(":8090", nil)
}
