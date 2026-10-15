package io.hyperion.omega.governance;

import org.junit.jupiter.api.Test;
import java.util.Map;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertEquals;

class FinOpsGovernanceApplicationTest {

    @Test
    void shouldSealMerkleHashSuccessfully() throws Exception {
        FinOpsGovernanceApplication app = new FinOpsGovernanceApplication();
        FinOpsGovernanceApplication.AuditRequest req = new FinOpsGovernanceApplication.AuditRequest(
            "tenant-alpha", "eu-west-1", "DEPLOY_CLUSTER_NODE"
        );

        Map<String, Object> result = app.sealMerkleAudit(req);
        assertEquals("tenant-alpha", result.get("tenantId"));
        assertNotNull(result.get("merkleHash"));
        assertEquals(64, ((String) result.get("merkleHash")).length());
    }
}
