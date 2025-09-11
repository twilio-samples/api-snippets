#!/usr/bin/env python3
"""
Performs RSA-OAEP decryption of AES key, then AES-GCM decryption of payload.
"""

import base64
import os
from pathlib import Path
from typing import Union

from cryptography.hazmat.primitives import hashes, serialization
from cryptography.hazmat.primitives.asymmetric import padding
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives.asymmetric.rsa import RSAPrivateKey


class DecryptScript:
    """
    CHANGE THESE VALUES TO YOURS
    """

    ENCRYPTED_KEY_B64 = "THIS_IS_THE_KEY_RETURNED_IN_THE_S3_HEADERS"
    IV_B64 = "THIS_IS_THE_IV_RETURNED_IN_THE_S3_HEADERS"
    CIPHERTEXT_FILE = "/path/to/encrypted/file"
    PRIVATE_KEY_FILE = "/path/to/your/private_key.pem"

    @staticmethod
    def main() -> None:
        """Main decryption process"""
        try:
            # Decode base64 values
            encrypted_key = base64.b64decode(DecryptScript.ENCRYPTED_KEY_B64)
            iv = base64.b64decode(DecryptScript.IV_B64)

            # Read encrypted file
            cipher_plus_tag = Path(DecryptScript.CIPHERTEXT_FILE).read_bytes()

            # Read private key
            private_key_string = Path(DecryptScript.PRIVATE_KEY_FILE).read_text()

            # 2. RSA-decrypt (unwrap) the AES data key
            rsa_private_key = DecryptScript._build_encryption_key_pair(private_key_string)

            # RSA/ECB/OAEPWithSHA-256AndMGF1Padding equivalent
            raw_aes_key = rsa_private_key.decrypt(
                encrypted_key,
                padding.OAEP(
                    mgf=padding.MGF1(algorithm=hashes.SHA256()),
                    algorithm=hashes.SHA256(),
                    label=None
                )
            )

            # 3. AES/GCM decrypt the payload
            # In AES-GCM, the authentication tag is typically appended to the ciphertext
            # We need to separate the ciphertext from the tag (last 16 bytes)
            ciphertext = cipher_plus_tag[:-16]
            tag = cipher_plus_tag[-16:]

            # Create AES-GCM cipher
            cipher = Cipher(
                algorithms.AES(raw_aes_key),
                modes.GCM(iv, tag),
                backend=default_backend()
            )
            decryptor = cipher.decryptor()

            # Decrypt the payload
            plaintext = decryptor.update(ciphertext) + decryptor.finalize()

            # 4. Persist plaintext to disk
            out_file = DecryptScript.CIPHERTEXT_FILE + ".json"
            DecryptScript._write_file(out_file, plaintext)
            print(f"Decrypted data written to: {out_file}")

        except Exception as e:
            print(f"Decryption failed: {e}")
            raise

    @staticmethod
    def _build_encryption_key_pair(private_key_string: str) -> RSAPrivateKey:
        """
        Parse private key from PEM string
        """
        try:
            private_key = serialization.load_pem_private_key(
                private_key_string.encode('utf-8'),
                password=None,
                backend=default_backend()
            )
            return private_key
        except Exception as ex:
            raise ValueError(f"Invalid private key: {ex}")

    @staticmethod
    def _write_file(path: Union[str, Path], data: bytes) -> None:
        """
        Write bytes to a file, creating or replacing it.
        """
        Path(path).write_bytes(data)


if __name__ == "__main__":
    DecryptScript.main()
