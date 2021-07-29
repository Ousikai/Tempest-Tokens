from secrets import *
import requests
from stellar_sdk import Server, Keypair, TransactionBuilder, Network

# Set Server Address
server = Server("https://horizon-testnet.stellar.org")

# Get Account Secrets and Addresses
source = Keypair.from_secret(sq1_secret)
source_account = server.load_account(account_id=source.public_key)
base_fee = server.fetch_base_fee()
transaction = (
    TransactionBuilder(
        source_account=source_account,
        network_passphrase=Network.TESTNET_NETWORK_PASSPHRASE,
        base_fee=base_fee,
    )
    .add_text_memo("Hello, Stellar!")
    .append_payment_op(kirito_address, "10", "XLM")
    .build()
)
transaction.sign(source)
response = server.submit_transaction(transaction)
print(response)