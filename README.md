# CloudflareDomainManager

CloudflareDomainManager is a small command-line app written in Go for managing DNS records for domains on Cloudflare.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your system.

### Steps

1. Clone the repository:
    ```sh
    git clone https://github.com/ferdzo/CloudflareDomainManager.git
    ```
2. Navigate to the project directory:
    ```sh
    cd CloudflareDomainManager
    ```
3. Build the application:
    ```sh
    go build -o cloudflareDomainManager
    ```

## Using the App

A simple command-line app that can be used to create, remove, modify, show, and perform other operations for DNS records on Cloudflare.

Usage:
```
[command]
```

Available Commands:
- **completion**: Generate the autocompletion script for the specified shell.
- **create**: Create a new DNS record.
- **delete**: Delete DNS records.
- **help**: Provide help about any command.
- **list**: List all DNS records in the zone with their IDs, types, names, contents, TTLs, and proxied status.
- **show**: Show the entire DNS record for the requested domain.
- **update**: Update a DNS record.

Flags:
- `-h, --help`: Help for this command.
- `-t, --toggle`: Help message for toggle.

Use `[command] --help` for more information about a command.

### Configuration

The credentials are stored in a simple `.env` file that should look something like this:
```
ZONE_ID = xxxxxxxxxxxxxxxxxxxxxxx
X_AUTH_KEY = xxxxxxxxxxxxxxxxxxxxxxxxxxx
X_AUTH_EMAIL = johndoe@doe.com
```

The Zone ID and the authorization key should be obtained from Cloudflare.

### Examples

```
cloudflareDomainManager show
```
This shows detailed DNS records.

```
cloudflareDomainManager create --type A --name test.ferdzo.xyz --content 1.1.1.1 --ttl 60 --proxied false
```
This command will create an A record with the name `test.ferdzo.xyz` that points to the IP address `1.1.1.1` and has a TTL of 1 minute (60 seconds).

```
cloudflareDomainManager delete asdasd632463asdhasdsad
```
Deletes a record with the DNS Record ID `asdasd632463asdhasdsad`.

## Error Handling

Common errors and troubleshooting steps:
- **Invalid credentials**: Ensure your `.env` file contains the correct `ZONE_ID`, `X_AUTH_KEY`, and `X_AUTH_EMAIL`.


## Contact Information

For support or questions, please contact [me](mailto:andrej@ferdzo.xyz).

