package key

import (
	"crypto/rand"

	"github.com/azer/logger"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/database/localdb"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

// Manager is local key-manager service (KMS) containing the account's private key
// and re-encryption keys that owned by the account.
type manager struct {
	ownerKey      *Key
	localDatabase localdb.Database
	log           *logger.Logger
}

func NewKeyManager(ownerKey *Key, localDatabase localdb.Database) Manager {
	kms := &manager{
		ownerKey:      ownerKey,
		localDatabase: localDatabase,
		log:           logger.New("kms"),
	}
	kms.log.Info("Node key loaded.", logger.Attrs{"address": ownerKey.EthereumAddress.Hex()})
	return kms
}

// NodeKey returns the owner key.
func (kms *manager) NodeKey() *Key {
	return kms.ownerKey
}

// DecryptExternalData looks for re-encryption key, then tries to decrypt.
func (kms *manager) DecryptExternalData(data *common.EncryptedData) (*common.Data, error) {
	// TODO: Implement key.Manager.DecryptExternalData
	return nil, nil
}

func (kms *manager) Encrypt(payload string) ([]byte, error) {
	publicKey := kms.ownerKey.ECIESPrivate.PublicKey
	return ecies.Encrypt(rand.Reader, &publicKey, []byte(payload), nil, nil)
}

func (kms *manager) Decrypt(encryptedPayload []byte) (string, error) {
	privateKey := kms.ownerKey.ECIESPrivate
	payload, err := privateKey.Decrypt(encryptedPayload, nil, nil)
	return string(payload), err
}

func (kms *manager) SignEthTx(tx *types.Transaction) (*types.Transaction, error) {
	return types.SignTx(tx, types.EIP155Signer{}, kms.ownerKey.PrivateKey)
}

func (kms *manager) SignBDBTx(tx *txn.Transaction) error {
	return tx.Sign([]*txn.KeyPair{kms.ownerKey.DeriveBigchainDBKeyPair()})
}
