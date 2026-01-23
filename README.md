# DanubeData Pulumi Provider

A Pulumi provider for managing [DanubeData](https://danubedata.ro) cloud infrastructure resources.

## Installation

### Node.js (TypeScript/JavaScript)

```bash
npm install @danubedata/pulumi
```

### Python

```bash
pip install pulumi_danubedata
```

### Go

```go
import "github.com/AdrianSilaghi/pulumi-danubedata/sdk/go/danubedata"
```

### .NET

```bash
dotnet add package DanubeData.Pulumi
```

## Configuration

Set your API token:

```bash
export DANUBEDATA_API_TOKEN="your-api-token"
```

Or configure in your Pulumi program:

```typescript
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config("danubedata");
// Set api_token in Pulumi.yaml or via `pulumi config set danubedata:apiToken --secret`
```

## Example Usage

### TypeScript

```typescript
import * as danubedata from "@danubedata/pulumi";

// Create an SSH key
const sshKey = new danubedata.SshKey("my-key", {
    name: "my-ssh-key",
    publicKey: "ssh-ed25519 AAAA... user@example.com",
});

// Create a VPS instance
const vps = new danubedata.Vps("my-server", {
    name: "web-server",
    image: "ubuntu-24.04",
    datacenter: "fsn1",
    resourceProfile: "vps-starter",
    authMethod: "ssh_key",
    sshKeyId: sshKey.id,
});

// Create a Redis cache
const cache = new danubedata.Cache("my-cache", {
    name: "app-cache",
    cacheProvider: "redis",
    resourceProfile: "cache-starter",
    datacenter: "fsn1",
});

// Create a PostgreSQL database
const database = new danubedata.Database("my-db", {
    name: "app-database",
    databaseProvider: "postgresql",
    resourceProfile: "db-starter",
    databaseName: "myapp",
    username: "appuser",
});

// Create an S3-compatible storage bucket
const bucket = new danubedata.StorageBucket("my-bucket", {
    name: "app-assets",
    region: "fsn1",
    versioningEnabled: true,
});

// Create a serverless container
const serverless = new danubedata.Serverless("my-app", {
    name: "api-service",
    deploymentType: "image",
    imageUrl: "ghcr.io/myorg/myapp:latest",
    port: 8080,
    minInstances: 0,
    maxInstances: 10,
});

// Export outputs
export const vpsIp = vps.publicIp;
export const cacheEndpoint = cache.endpoint;
export const databaseEndpoint = database.endpoint;
export const bucketEndpoint = bucket.endpointUrl;
export const serverlessUrl = serverless.url;
```

### Python

```python
import pulumi
import pulumi_danubedata as danubedata

# Create an SSH key
ssh_key = danubedata.SshKey("my-key",
    name="my-ssh-key",
    public_key="ssh-ed25519 AAAA... user@example.com")

# Create a VPS instance
vps = danubedata.Vps("my-server",
    name="web-server",
    image="ubuntu-24.04",
    datacenter="fsn1",
    resource_profile="vps-starter",
    auth_method="ssh_key",
    ssh_key_id=ssh_key.id)

# Create a Redis cache
cache = danubedata.Cache("my-cache",
    name="app-cache",
    cache_provider="redis",
    resource_profile="cache-starter",
    datacenter="fsn1")

# Export outputs
pulumi.export("vps_ip", vps.public_ip)
pulumi.export("cache_endpoint", cache.endpoint)
```

## Available Resources

| Resource | Description |
|----------|-------------|
| `Vps` | Virtual Private Server instances |
| `Serverless` | Serverless containers with scale-to-zero |
| `Cache` | Redis, Valkey, or Dragonfly cache instances |
| `Database` | MySQL, PostgreSQL, or MariaDB databases |
| `StorageBucket` | S3-compatible object storage buckets |
| `StorageAccessKey` | Access keys for storage buckets |
| `SshKey` | SSH keys for VPS authentication |
| `Firewall` | Network firewall rules |
| `VpsSnapshot` | VPS snapshots for backup/recovery |

## Available Data Sources

| Data Source | Description |
|-------------|-------------|
| `getVpsImages` | List available VPS operating system images |
| `getCacheProviders` | List cache providers (redis, valkey, dragonfly) |
| `getDatabaseProviders` | List database providers (mysql, postgresql, mariadb) |
| `getSshKeys` | List SSH keys in your account |
| `getVpss` | List all VPS instances |
| `getDatabases` | List all database instances |
| `getCaches` | List all cache instances |
| `getFirewalls` | List all firewalls |
| `getServerlessContainers` | List serverless containers |
| `getStorageBuckets` | List storage buckets |
| `getStorageAccessKeys` | List storage access keys |
| `getVpsSnapshots` | List VPS snapshots |

## Development

### Prerequisites

- Go 1.24+
- Pulumi CLI
- Node.js 18+ (for TypeScript SDK)
- Python 3.8+ (for Python SDK)

### Building

```bash
# Build the provider
make build

# Generate SDKs
make generate

# Install for local development
make dev
```

### Testing

```bash
make test
```

## License

Apache-2.0
