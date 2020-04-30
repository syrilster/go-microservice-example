[![Actions Status](https://github.com/syrilster/go-microservice-example/workflows/Build%20and%20test%20Go/badge.svg)](https://github.com/syrilster/go-microservice-example/actions)

[![codecov](https://codecov.io/gh/syrilster/go-microservice-example/branch/master/graph/badge.svg)](https://codecov.io/gh/syrilster/go-microservice-example)

# go microservice example

This is an API layer built to demonstrate service to service call in Golang. This is a currency conversion service which calls the
currency-exchange (Java) service to get the currency day rates for the conversion.

## Project Structure

- [config](./config) -  configuration files for the service and infrastructure.
    - [dev](./config/dev) - configuration overrides for the **development** environment.
    - [prod](./config/prod) - configuration overrides for the **production** environment.
    - [test](./config/test) - configuration overrides for the **test** environment.
- [internal](./internal) - internal go packages that are specific to this service and are non exportable.
- [test](./test) - test related artifacts such as: smoke tests, test data, etc.
    
# K8's issues
* Error when executing kubectl get all or any other command. This is because the keys to your old cluster is cached.
   ```
    Error:
    Unable to connect to the server: x509: certificate signed by unknown authority
    
    Fix:
    gcloud container clusters get-credentials YOURCLUSTERHERE --zone YOURCLUSTERZONEHERE
    ```
# References
* [Enabling rolling update in k8's](https://medium.com/platformer-blog/enable-rolling-updates-in-kubernetes-with-zero-downtime-31d7ec388c81)
* [K8's service to service communication](https://dev.to/azure/how-to-access-your-kubernetes-applications-using-services-5626)
* [Github actions example](https://brunopaz.dev/blog/building-a-basic-ci-cd-pipeline-for-a-golang-application-using-github-actions)
* [Github actions GoLang](https://presstige.io/p/Using-GitHub-Actions-with-Go-2ca9744b531f4f21bdae9976d1ccbb58)
