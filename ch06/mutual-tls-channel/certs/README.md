# Generate private RSA key

To generate RSA key using OpenSSL tool, we need to use `genrsa` command like below,

```shell script
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

So now we successfully created our private key(`server.key`) and we are going to use this in our gRPC server.

# Generate CA and self-signed certificates
Let’s create a Certificate Authority and self-signed certificate for our example. This is similar to what we have done to create the server key and self-signed certificate in the previous section. To generate RSA key using OpenSSL tool, execute the following command,

```shell script
$ openssl genrsa -aes256 -out ca.key 4096
```

Here we create a new private key with a password for the CA. Now we can create the root CA certificate with a validity of two years using the SHA256 hash algorithm.

```shell script
$ openssl req -new -x509 -sha256 -days 730 -key ca.key -out ca.crt
```

As in the server certificate creation earlier, Certificate generation is an interactive process. You can make them blank by entering `.`. But you need to give a name for Common Name as mentioned before.

So now we created both the private key and self-signed certificate of our Certificate Authority. We can verify the root certificate using below command,

```shell script
$ openssl x509 -noout -text -in ca.crt
```

We can check the validity period for 02 years and the issuer and subject should both be set to the value passed for “Common Name” because this is a root certificate and it is self-signed.

The next step is to create a server private key and certificate. Unlike the previous section, we need to get the certificate signed by our new Certificate Authority(CA). 

# Generate server certificate
Once we have the server private key, we can proceed to create a Certificate Signing Request (CSR). This is a formal request asking a CA to sign a certificate, and it contains the public key of the entity requesting the certificate and some information about the entity. Like certificate creation, CSR creation process is also an interactive process during which you’ll be providing the elements of the certificate distinguished name.

Execute the following command to create a certificate signing request.
```shell script
$ openssl req -new -sha256 -key server.key -out server.csr
```
After a CSR is generated, we can sign the request and generate the certificate using our own CA certificate. Normally, the CA and the certificate requester are two different companies who don’t want to share their private keys. That’s why we need this intermediate step for certificate creation.

Let’s execute the following command to use our root CA to sign the CSR and create server certificate. You’ll be prompted to enter the root CA’s password.
```shell script
$ openssl x509 -req -days 365 -sha256 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 1 -out server.crt
```
Now we have created server key(server.key) and server certificate(server.crt). We can use them to enable mutual TLS in server side later. Let’s create client key and certificate,

# Generate client key and certificate
Generating the client certificate is very similar to creating the server certificate. We need to execute the following commands to create a private key, create a certificate signing request and create a certificate for client application.
```shell script
$ openssl genrsa -out client.key 2048
$ openssl req -new -key client.key -out client.csr
$ openssl x509 -req -days 365 -sha256 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 2 -out client.crt
```

# Convert server/client keys to pem format
In order to secure java application, we need to provide key store(.pem file). We can easily convert the server and client keys using following command 

```shell script
$ openssl pkcs8 -topk8 -inform pem -in server.key -outform pem -nocrypt -out server.pem
$ openssl pkcs8 -topk8 -inform pem -in client.key -outform pem -nocrypt -out client.pem
```