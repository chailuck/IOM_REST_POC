## AM Secret
- API Key : 0emB2YuLZVzwSKtTGcERbvG6ptWo5hmZ

## Configuration
The service uses ETCD for configuration management. Key configuration parameters:

```yaml
api:
  port: 8080
  version: v1
  timeout: 30s

etcd:
  endpoints:
    - localhost:2379
  prefix: /api-manager

auth:
  enabled: true
  jwt_secret: your-secret-key
```

## Installation

```bash
# Clone repository
git clone https://gitlab.com/ft25/iom/poc/am.git

# Configure GitLab private repository access
# Replace <your_gitlab_token> with your actual GitLab personal access token
git config --global url."https://oauth2:<your_gitlab_token>@gitlab.com".insteadOf "https://gitlab.com"

# Install dependencies
# For private repositories in go.mod:
# gitlab.com/ft25/iom/framework => gitlab.com/ft25/iom/framework.git
# gitlab.com/ft25/iom/model => gitlab.com/ft25/iom/uat/model.git
go mod download

```

### Managing Dependencies
To update specific dependencies:

```bash
# Update all dependencies
go get -u ./...

# Update specific dependency
go get -u gitlab.com/ft25/iom/framework
go get -u gitlab.com/ft25/iom/model

# Tidy up go.mod and go.sum
go mod tidy
```


## Usage

```bash
# Run service
./api-manager

# With custom config
./api-manager --config=/path/to/config.yaml
```

## API Documentation
Swagger UI is available at: http://localhost:8080/swagger/index.html

## Testing

```bash
# Run unit tests
go test ./...

# Run integration tests
go test -tags=integration ./...
```

## Contributing
1. Fork the repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Create Pull Request

## License
Copyright (c) 2024 True Corporation