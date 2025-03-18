Contributing
============

## Syncing quickstart modules to the BSR

When adding a module for a quickstart that should be available on the BSR, there are a few prerequisite steps:

### Create the module in the BSR

Ensure that the module you plan on syncing has been created in the BSR. For example, if you were adding a `foo` module in the [tutorials](https://buf.build/quickstarts) organization, you would need to make sure that a new module is created there before you start syncing the protobuf files.

### Set the module up with the [buf-push-action](https://github.com/bufbuild/buf-push-action)

Add the new module to the [buf workflow](../workflows/buf.yaml), similarly to how the `buf-schema-registry` module is configured there.