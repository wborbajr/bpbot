# Generate the private key for the client side (CA)
openssl genrsa -out ca.key 4096
# Generate the self-signed certificate with the CA primary key
openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 365 -out ca.crt
# Generate the private key for the server side
openssl genrsa -out server.key 4096
# Generate a certficate Signing Request
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
# Use CA.crt and CA.key to sign the certificate for the server side
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365