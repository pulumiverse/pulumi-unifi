# Unifi Resource Provider

The Unifi Resource Provider lets you manage [Unifi](https://www.ui.com) networks.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/unifi
```

or `yarn`:

```bash
yarn add @pulumiverse/unifi
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumiverse_unifi
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-unifi/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumiverse.Unifi
```

## Configuration

The following configuration points are available for the `unifi` provider:

- `unifi:apiKey` (environment: `UNIFI_API_KEY`) - the API key for `unifi`
- `unifi:region` (environment: `UNIFI_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/unifi/api-docs/).
