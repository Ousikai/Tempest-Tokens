from secrets import *
import requests
from stellar_sdk import Server, Keypair, TransactionBuilder, Network

# Set Server Address
server = Server("https://horizon-testnet.stellar.org")

# Get Account Secrets and Addresses
kirito_keypair = Keypair.from_secret(kirito_secret)
# destination = Keypair.random()
destination = oscar_address

# Create Account
source_account = server.load_account(account_id=kirito_keypair.public_key)
transaction = TransactionBuilder(
    source_account=source_account,
    network_passphrase=Network.TESTNET_NETWORK_PASSPHRASE,
    base_fee=100) \
    .append_create_account_op(destination=destination, starting_balance="1.5") \
    .build()
transaction.sign(kirito_keypair)
response = server.submit_transaction(transaction)
# print("Transaction hash: {}".format(response["hash"]))
print(response)
