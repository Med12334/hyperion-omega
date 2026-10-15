package io.hyperion.omega.governance;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.time.Instant;
import java.util.HexFormat;
import java.util.Map;
import java.util.UUID;

@SpringBootApplication
@RestController
@RequestMapping("/api/v1/finops")
@CrossOrigin(origins = "*")
public class FinOpsGovernanceApplication {

    private static final String GENESIS_HASH = "0000000000000000000000000000000000000000000000000000000000000000";

    public static void main(String[] args) {
        SpringApplication.run(FinOpsGovernanceApplication.class, args);
    }

    public record AuditRequest(String tenantId, String clusterRegion, String operationType) {}

    @PostMapping("/audit-seal")
    public Map<String, Object> sealMerkleAudit(@RequestBody AuditRequest req) throws Exception {
        UUID auditId = UUID.randomUUID();
        String merkleHash = computeSha256(GENESIS_HASH + "|" + req.tenantId() + "|" + req.operationType());

        return Map.of(
            "auditId", auditId,
            "tenantId", req.tenantId(),
            "clusterRegion", req.clusterRegion(),
            "operationType", req.operationType(),
            "previousHash", GENESIS_HASH,
            "merkleHash", merkleHash,
            "timestamp", Instant.now()
        );
    }

    private String computeSha256(String raw) throws Exception {
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        byte[] hash = digest.digest(raw.getBytes(StandardCharsets.UTF_8));
        return HexFormat.of().formatHex(hash);
    }
}
