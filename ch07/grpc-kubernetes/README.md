


kubectl apply -f server/grpc-prodinfo-server.yaml

kubectl apply -f client/grpc-prodinfo-client-job.yaml 

kubectl get pods 

Check for : 'grpc-productinfo-server*' and 'grpc-productinfo-client*' pods and make sure they are running and completed state respectively. 

kubectl logs grpc-productinfo-server-587b894b7c-276ll
    2019/09/07 22:25:05 New product added - ID : 577c03e2-d1be-11e9-bcb6-46a2b71e2706, Name : Sumsung S10
    2019/09/07 22:25:05 New product retrieved - ID : value:"577c03e2-d1be-11e9-bcb6-46a2b71e2706"


    kubectl logs grpc-productinfo-client-sqs4r
        2019/09/07 22:55:27 Product ID: 95c5ef97-d1c2-11e9-bcb6-46a2b71e2706 added successfully
        2019/09/07 22:55:27 Product: %!(EXTRA string=id:"95c5ef97-d1c2-11e9-bcb6-46a2b71e2706" name:"Sumsung S10" description:"Samsung Galaxy S10 is the latest smart phone, launched in February 2019" price:700 )


