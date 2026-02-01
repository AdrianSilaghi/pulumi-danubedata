"""Example DanubeData infrastructure with Pulumi (Python)"""
import pulumi
import pulumi_danubedata as danubedata

# Create an SSH key for VPS authentication
ssh_key = danubedata.SshKey("my-key",
    name="deployment-key",
    public_key="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIExample... user@example.com")

# Create a VPS instance
web_server = danubedata.Vps("web-server",
    name="web-server",
    image="ubuntu-24.04",
    datacenter="fsn1",
    resource_profile="nano_shared",
    auth_method="ssh_key",
    ssh_key_id=ssh_key.id)

# Create a Redis cache for session storage
session_cache = danubedata.Cache("session-cache",
    name="session-cache",
    cache_provider="redis",
    resource_profile="micro",
    datacenter="fsn1")

# Create a PostgreSQL database
app_database = danubedata.Database("app-db",
    name="app-database",
    engine="postgresql",
    resource_profile="small",
    database_name="myapp",
    datacenter="fsn1")

# Create an S3-compatible storage bucket for assets
assets_bucket = danubedata.StorageBucket("assets",
    name="app-assets",
    region="fsn1",
    versioning_enabled=True)

# Create access keys for the storage bucket
storage_key = danubedata.StorageAccessKey("assets-key",
    name="app-assets-key")

# Export the outputs
pulumi.export("vps_public_ip", web_server.public_ip)
pulumi.export("cache_endpoint", session_cache.endpoint)
pulumi.export("cache_port", session_cache.port)
pulumi.export("database_endpoint", app_database.endpoint)
pulumi.export("bucket_endpoint", assets_bucket.endpoint_url)
