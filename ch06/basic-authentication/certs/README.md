#Creating private RSA key

To generate RSA key using OpenSSL tool, we need to use `genrsa` command like below,

```shell
$ openssl genrsa<1> -out server.key<2> 2048<3>
Generating RSA private key, 2048 bit long modulus
....................................................+++
.............................................+++
e is 65537 (0x10001)
```

1. Specifies which algorithm to use to create the key. OpenSSL supports creating keys with a different algorithm like
 RSA, DSA, and ECDSA. All types are practical for use in all scenarios. For example, for web server keys commonly uses RSA. In our case, we need to generate RSA type key.
2. Specifies the name of the generated key. Can have any name with `.key` as extension.
3. Specifies the size of the key. The default size for RSA keys is only 512 bits, which is not secure because an
 intruder can use brute force to recover your private key. So we use a 2048-bit RSA key which is considered to be secure.

Here you can also add a passphrase to the key. So you need the passphrase whenever you need to use the key. In this example, we are not going to add a passphrase to the key.

So now we successfully created our private key(`server.key`) and we are going to use this in our gRPC server. Let’s create a self-signed public certificate to distribute among our clients. 

# Creating public key/certificate
Once we have the private key, we need to create the certificate. In this example, we are going to create a self-signed certificate. In other words, there is no certificate authority(CA). Typically in a production deployment, either you will use a public certificate authority or an enterprise-level certificate authority to sign your public certificate. So any client who trusts the certificate authority can verify.

Since in this example, our TLS server is for our own testing purposes, we probably don’t want to go to a certificate authority(CA) for a publicly trusted certificate.

Let’s execute the following command to generate a self-signed public certificate. Certificate generation is an interactive process during which you’ll be asked to enter information which is going to be incorporated with the certificate.

```shell
$ openssl req -new -x509<1> -sha256<2> -key server.key<3> \
              -out server.crt<4> -days 3650<5>
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields, there will be a default value 
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:US
State or Province Name (full name) []:California
Locality Name (eg, city) []:Mountain View
Organization Name (eg, company) []:O’Reilly Media, Inc
Organizational Unit Name (eg, section) []:Publishers
Common Name (eg, fully qualified hostname[]:localhost
Email Address []:webmaster@localhost
```

1. Specifies the format of the public certificate. X.509 is a standard format which is used in many Internet protocols
, including TLS/SSL.
2. Specifies the secure hash algorithm.
3. Specifies private key(`server.key`) file location which we generated before.
4. Specifies the name of the generated certificate. Can have any name with `.crt` as extension.
5. Specifies the lifetime of the certificate to 365 days.

[NOTE]
=====
The most important question among the questions asked when generating the certificate, is the “Common Name” which is composed of the host, domain, or IP address of the server related to the certificate. This name is used during the verification and if the host name doesn’t match the common name a warning is raised.
=====

# Creating CA and self-signed certificates
Let’s create a Certificate Authority and self-signed certificate for our example. This is similar to what we have done to create the server key and self-signed certificate in the previous section. To generate RSA key using OpenSSL tool, execute the following command,

````shell script
$ openssl genrsa -aes256 -out ca.key 4096
````

Here we create a new private key with a password for the CA. Now we can create the root CA certificate with a validity of two years using the SHA256 hash algorithm.

````shell script
$ openssl req -new -x509 -sha256 -days 730 -key ca.key -out ca.crt
````

As in the server certificate creation earlier, Certificate generation is an interactive process. You can make them blank by entering `.`. But you need to give a name for Common Name as mentioned before.

So now we created both the private key and self-signed certificate of our Certificate Authority. We can verify the root certificate using below command,

````shell script
$ openssl x509 -noout -text -in ca.crt
````

We can check the validity period for 02 years and the issuer and subject should both be set to the value passed for “Common Name” because this is a root certificate and it is self-signed.

The next step is to create a server private key and certificate. Unlike the previous section, we need to get the certificate signed by our new Certificate Authority(CA). 

# Creating server key and certificate
To create server private key, we need to follow the same steps as in the section `Creating private RSA key`. So we are going to use the same key for this scenario.

Once we have the server private key, we can proceed to create a Certificate Signing Request (CSR). This is a formal request asking a CA to sign a certificate, and it contains the public key of the entity requesting the certificate and some information about the entity. Like certificate creation, CSR creation process is also an interactive process during which you’ll be providing the elements of the certificate distinguished name.

Execute the following command to create a certificate signing request.
````shell script
$ openssl req -new -sha256 -key server.key -out server.csr
````
After a CSR is generated, we can sign the request and generate the certificate using our own CA certificate. Normally, the CA and the certificate requester are two different companies who don’t want to share their private keys. That’s why we need this intermediate step for certificate creation.

Let’s execute the following command to use our root CA to sign the CSR and create server certificate. You’ll be prompted to enter the root CA’s password.
````shell script
$ openssl x509 -req -days 365 -sha256 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 1 -out server.crt
````
Now we have created server key(server.key) and server certificate(server.crt). We can use them to enable mutual TLS in server side later. Let’s create client key and certificate,

# Creating client key and certificate
Generating the client certificate is very similar to creating the server certificate. We need to execute the following commands to create a private key, create a certificate signing request and create a certificate for client application.
````shell script
$ openssl genrsa -out client.key 2048
$ openssl req -new -key client.key -out client.csr
$ openssl x509 -req -days 365 -sha256 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 2 -out client.crt
````
