import * as pulumi from "@pulumi/pulumi";
import * as danubedata from "@danubedata/pulumi";

// Test 1: Create an SSH key
const testSshKey = new danubedata.SshKey("test-key", {
    name: "pulumi-test-key",
    publicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIExampleTestKey pulumi-test@danubedata.ro",
});

// Test 2: Create a storage bucket
const testBucket = new danubedata.StorageBucket("test-bucket", {
    name: "pulumi-test-bucket",
    region: "fsn1",
    versioningEnabled: false,
});

// Export outputs to verify resources were created
export const sshKeyId = testSshKey.id;
export const sshKeyName = testSshKey.name;
export const bucketEndpoint = testBucket.endpointUrl;
export const bucketName = testBucket.minioBucketName;
