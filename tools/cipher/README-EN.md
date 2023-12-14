# How to Use the cipher Encryption and Decryption Script #

## Brief Introduction ##

This script integrates the RSA and GCM encryption and decryption methods.

 *  GCM symmetric encryption and decryption requires 16-bit, 24-bit, or 32-bit keys and 16-bit nonce for background sensitive data encryption and decryption.
 *  For asymmetric RSA encryption and decryption, public and private keys need to be generated. For details about how to generate public and private keys, see [Deployment guide](../../doc/deployment/deployment-guide-EN.md)	For encrypting and decrypting sensitive data during network transmission, see Step 4 "Preparing Certificates" in the deployment guide.

## How to Use ##

 *  Configuration`go`Language environment variable, run:`go env -w GOOS=linux/winodws`, used to generate`linux`/`windows`Executable file under
 *  Use the`go build ./main.go`Generated`windows`/`linux`After the executable file is executed, the encryption and decryption operations can be performed after the execution parameters are entered.

## Parameter Description ##

|   Field name   |                     Description                      |                                                        Parameter Requirements                                                         |
|:--------------:|:----------------------------------------------------:|:-------------------------------------------------------------------------------------------------------------------------------------:|
|     -mode      |               Encryption or decryption               |                                                             encode/decode                                                             |
|      -str      | Plaintext or ciphertext for encryption or decryption |                                                                  --                                                                   |
|      -key      |        Key for GCM encryption and decryption.        | GCM encryption and decryption are mandatory. RSA encryption and decryption are optional. The length must be 16, 24, or 32 characters. |
|     -nonce     |       Nonce for GCM encryption and decryption        |       GCM encryption and decryption are mandatory, and RSA encryption and decryption are optional. The length is 16 characters.       |
| -cipher-method |           Encryption and Decryption Method           |                                                                GCM/RSA                                                                |
| -private-cert  |    RSA private key for encryption and decryption     |                                                     Path in Linux/Windows format                                                      |
|  -public-cert  |     Public key for RSA encryption and decryption     |                                 Path in Linux or Windows format, generated based on the private key.                                  |



