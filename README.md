# go microservice example

This is an API layer built to demonstrate service to service call in Golang. This is a currency conversion service which calls the
currency-exchange (Java) service to get the currency day rates for the conversion.

# K8's issues
* Error when executing kubectl get all or any other command. This is because the keys to your old cluster is cached.
   ```
    Error:
    Unable to connect to the server: x509: certificate signed by unknown authority
    
    Fix:
    gcloud container clusters get-credentials YOURCLUSTERHERE --zone YOURCLUSTERZONEHERE
    ```
