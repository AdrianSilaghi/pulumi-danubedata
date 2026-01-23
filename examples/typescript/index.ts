import * as pulumi from "@pulumi/pulumi";
import * as danubedata from "@danubedata/pulumi";

// Create an SSH key for VPS authentication
const sshKey = new danubedata.SshKey("my-key", {
    name: "deployment-key",
    publicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIExample... user@example.com",
});

// Create a VPS instance
const webServer = new danubedata.Vps("web-server", {
    name: "web-server",
    image: "ubuntu-24.04",
    datacenter: "fsn1",
    resourceProfile: "vps-starter",
    authMethod: "ssh_key",
    sshKeyId: sshKey.id,
});

// Create a Redis cache for session storage
const sessionCache = new danubedata.Cache("session-cache", {
    name: "session-cache",
    cacheProvider: "redis",
    resourceProfile: "cache-starter",
    datacenter: "fsn1",
});

// Create a PostgreSQL database
const appDatabase = new danubedata.Database("app-db", {
    name: "app-database",
    databaseProvider: "postgresql",
    resourceProfile: "db-starter",
    databaseName: "myapp",
    username: "appuser",
    datacenter: "fsn1",
});

// Create an S3-compatible storage bucket for assets
const assetsBucket = new danubedata.StorageBucket("assets", {
    name: "app-assets",
    region: "fsn1",
    versioningEnabled: true,
});

// Create access keys for the storage bucket
const storageKey = new danubedata.StorageAccessKey("assets-key", {
    name: "app-assets-key",
});

// Export the outputs
export const vpsPublicIp = webServer.publicIp;
export const cacheEndpoint = sessionCache.endpoint;
export const cachePort = sessionCache.port;
export const databaseEndpoint = appDatabase.endpoint;
export const bucketEndpoint = assetsBucket.endpointUrl;
