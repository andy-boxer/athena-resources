# Labeled Image deployment action

This action creates a request to the Labeled Image Deployment api server
The action will retrieve the monitor url returned by the api and then use that to monitor for success or failure.
The completion status is then returned so the workflow can make recovery decisions if required.

## Inputs

### `image_reg`

**Required** The FQDN of the service image to be deployed.

### `namespace`

**Required** The namespace into which to deploy the service.

### `deployment_kind`

**Required** The kind of the deployment. Default `"regular"`.

## Outputs

### `status`

The completion status of the deployment

## Example usage

```yaml
uses: ./.github/actions/lid-deployment
with:
  image_ref: '404210962762.dkr.ecr.eu-west-2.amazonaws.com/abx01-repository:hello-world-v0.0.2'
  namespace: 'test1-hello'
  deployment_kind: 'regular'
```
