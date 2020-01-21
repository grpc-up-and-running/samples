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

So now we successfully created our private key(`server.key`) and we are going to use this in our gRPC server. Let’s create a self-signed public certificate to distribute among our clients. 

# Generate public key/certificate
Once we have the private key, we need to create the certificate. In this example, we are going to create a self-signed certificate. In other words, there is no certificate authority(CA). Typically in a production deployment, either you will use a public certificate authority or an enterprise-level certificate authority to sign your public certificate. So any client who trusts the certificate authority can verify.

Since in this example, our TLS server is for our own testing purposes, we probably don’t want to go to a certificate authority(CA) for a publicly trusted certificate.

Let’s execute the following command to generate a self-signed public certificate. Certificate generation is an interactive process during which you’ll be asked to enter information which is going to be incorporated with the certificate.

```shell script
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

# Convert server/client keys to pem format
In order to secure java application, we need to provide key store(.pem file). We can easily convert the server and client keys using following command 

```shell script
$ openssl pkcs8 -topk8 -inform pem -in server.key -outform pem -nocrypt -out server.pem
$ openssl pkcs8 -topk8 -inform pem -in client.key -outform pem -nocrypt -out client.pem
```