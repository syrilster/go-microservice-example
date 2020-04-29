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
* https://medium.com/platformer-blog/enable-rolling-updates-in-kubernetes-with-zero-downtime-31d7ec388c81
* https://dev.to/azure/how-to-access-your-kubernetes-applications-using-services-5626