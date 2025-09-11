import java.io.StringReader;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.KeyPair;
import java.security.PrivateKey;
import java.security.spec.MGF1ParameterSpec;
import java.util.Base64;
import javax.crypto.Cipher;
import javax.crypto.SecretKey;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.OAEPParameterSpec;
import javax.crypto.spec.PSource;
import javax.crypto.spec.SecretKeySpec;

import org.bouncycastle.openssl.PEMKeyPair;
import org.bouncycastle.openssl.PEMParser;
import org.bouncycastle.openssl.jcajce.JcaPEMKeyConverter;

public class DecryptScript {

    /*
     * CHANGE THESE VALUES TO YOURS
     */

    private static final String ENCRYPTED_KEY_B64 = "THIS_IS_THE_KEY_RETURNED_IN_THE_S3_HEADERS";
    private static final String IV_B64 = "THIS_IS_THE_IV_RETURNED_IN_THE_S3_HEADERS";
    private static final String CIPHERTEXT_FILE = "/path/to/encrypted/file";
    private static final String PRIVATE_KEY_FILE = "/path/to/private/key.pem";

    public static void main(final String[] args) throws Exception {
        // Load BC
        Security.addProvider(new BouncyCastleProvider());

        final byte[] encryptedKey = Base64.getDecoder().decode(ENCRYPTED_KEY_B64);
        final byte[] iv = Base64.getDecoder().decode(IV_B64);
        final byte[] cipherPlusTag = Files.readAllBytes(Path.of(CIPHERTEXT_FILE));

        final String privateKeyString = Files.readString(Path.of(PRIVATE_KEY_FILE));

        /* 2. RSA-decrypt (unwrap) the AES data key */
        final PrivateKey rsaPriv = buildEncryptionKeyPair(privateKeyString);

        final Cipher rsa = Cipher.getInstance("RSA/ECB/OAEPWithSHA-256AndMGF1Padding");
        final OAEPParameterSpec oaep = new OAEPParameterSpec(
            "SHA-256",
            "MGF1",
            MGF1ParameterSpec.SHA256,
            PSource.PSpecified.DEFAULT);
        rsa.init(Cipher.DECRYPT_MODE, rsaPriv, oaep);

        final byte[] rawAesKey = rsa.doFinal(encryptedKey);
        final SecretKey aesKey = new SecretKeySpec(rawAesKey, "AES");

        /* 3. AES/GCM decrypt the payload */
        final Cipher gcm = Cipher.getInstance("AES/GCM/NoPadding");
        final GCMParameterSpec gcmSpec = new GCMParameterSpec(128, iv);
        gcm.init(Cipher.DECRYPT_MODE, aesKey, gcmSpec);

        final byte[] plaintext = gcm.doFinal(cipherPlusTag);

        /* 4. Persist plaintext to disk */
        final String outFile = CIPHERTEXT_FILE + ".json";
        writeFile(outFile, plaintext);
        System.out.println("Decrypted data written to: " + outFile);
    }

    private static PrivateKey buildEncryptionKeyPair(final String privateKey) {
        try (PEMParser pemParser = new PEMParser(new StringReader(privateKey))) {
            JcaPEMKeyConverter converter = (new JcaPEMKeyConverter()).setProvider("BC");
            PrivateKeyInfo privateKeyInfo = (PrivateKeyInfo) pemParser.readObject();
            return converter.getPrivateKey(privateKeyInfo);
        } catch (Exception e) {
            throw new IllegalArgumentException("invalid private key", e);
        }
    }

    /**
     * Write bytes to a file, creating or replacing it.
     */
    private static void writeFile(final String path, final byte[] data) throws Exception {
        Files.write(Path.of(path), data);
    }
}