# Production TODO

## Server

- [ ] Replace the static bearer token with JWT/OIDC verification, key rotation, audience/issuer checks, and per-upload ACL enforcement for HEAD/PATCH/GET/DELETE.
- [ ] Persist upload/business metadata and idempotency keys in a database; do not use local `.info` files as the business source of truth.
- [ ] Add tenant/user quotas, rate limiting, concurrent-upload limits, and abuse controls.
- [ ] Add unfinished-upload expiration and a cleanup job. The local `filestore` does not clean abandoned files automatically.
- [ ] Add antivirus/content scanning and quarantine before publishing a completed file.
- [ ] Add MIME sniffing, archive-bomb checks, image/document validation, and filename/content-policy validation.
- [ ] Replace local storage with S3/GCS/Azure when horizontal scaling is required; use a distributed locking design compatible with the chosen storage.
- [ ] Add Prometheus metrics, tracing, structured audit logs, alerting, and storage-capacity monitoring.
- [ ] Add HTTPS or place the service behind a correctly configured reverse proxy; test forwarded headers and buffering behavior.
- [ ] Decide whether to disable download, termination, concatenation, or the experimental protocol according to product requirements.
- [ ] Implement a post-completion workflow: scan -> persist final record -> publish -> notify downstream systems.
- [ ] Add integration tests for interruption, offset conflict, duplicate requests, disk-full, server restart, and concurrent PATCH requests.

## Client

- [ ] Encrypt or protect checkpoint files if upload URLs are sensitive bearer capabilities.
- [ ] Add checksum-extension support after feature detection through OPTIONS.
- [ ] Add deferred upload length for streams whose total size is unknown.
- [ ] Add parallel upload through tus concatenation (`Upload-Concat: partial/final`) when the backend benefits from it.
- [ ] Add multi-file scheduling, bandwidth limits, pause/resume UI, and persistent upload history.
- [ ] Add custom retry classification, jitter, Retry-After support, proxy configuration, and TLS/mTLS configuration.
- [ ] Remove or terminate orphaned remote uploads when `-restart` creates a replacement upload.
