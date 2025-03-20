Contributing
============

## Syncing quickstart modules to the BSR

When adding a module for a quickstart that should be available on the BSR, there are a few prerequisite steps:

### Create the module in the BSR

Ensure that the module you plan on syncing has been created in the BSR. For example, if you're adding a `foo` module in the [quickstarts](https://buf.build/quickstarts) organization, make sure that a new module is created there before you start syncing the Protobuf files.

### Set the module up with the [buf-push-action](https://github.com/bufbuild/buf-push-action)

Add the new module to a `buf.yaml` file in this directory, similar to what exists in the [buf-tour](https://github.com/bufbuild/buf-tour/blob/main/.github/workflows/buf.yaml) repo.