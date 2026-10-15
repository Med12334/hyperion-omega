//! Hyperion-Omega Layer 2 & 3: Rust Zero-Trust Edge Proxy & eBPF Syscall Filter

#[derive(Debug, Clone, PartialEq)]
pub enumEdgeVerdict {
    AllowPassThrough,
    DropRateLimited,
    QuarantineKernelEscape,
}

#[derive(Debug, Clone)]
pub struct PacketFrame {
    pub source_ip: String,
    pub path: String,
    pub syscall_hint: String,
    pub requests_per_sec: u32,
}

pub struct EdgeSentinelEngine {
    pub max_rps_threshold: u32,
}

impl EdgeSentinelEngine {
    pub fn new(max_rps_threshold: u32) -> Self {
        Self { max_rps_threshold }
    }

    pub fn inspect_frame(&self, frame: &PacketFrame) -> enumEdgeVerdict {
        let syscall = frame.syscall_hint.to_lowercase();
        if syscall.contains("ptrace_attach")
            || syscall.contains("docker.sock")
            || syscall.contains("cap_sys_admin")
        {
            return enumEdgeVerdict::QuarantineKernelEscape;
        }

        if frame.requests_per_sec > self.max_rps_threshold {
            return enumEdgeVerdict::DropRateLimited;
        }

        enumEdgeVerdict::AllowPassThrough
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn should_quarantine_ebpf_container_escape_syscall() {
        let engine = EdgeSentinelEngine::new(1000);
        let malicious_frame = PacketFrame {
            source_ip: String::from("192.0.2.99"),
            path: String::from("/api/v1/compute/exec"),
            syscall_hint: String::from("sys_ptrace_attach_host_ns"),
            requests_per_sec: 15,
        };
        assert_eq!(
            engine.inspect_frame(&malicious_frame),
            enumEdgeVerdict::QuarantineKernelEscape
        );
    }

    #[test]
    fn should_drop_volumetric_ddos_spike() {
        let engine = EdgeSentinelEngine::new(500);
        let flood_frame = PacketFrame {
            source_ip: String::from("198.51.100.42"),
            path: String::from("/api/v1/schedule"),
            syscall_hint: String::from("sys_read"),
            requests_per_sec: 1200,
        };
        assert_eq!(
            engine.inspect_frame(&flood_frame),
            enumEdgeVerdict::DropRateLimited
        );
    }
}
