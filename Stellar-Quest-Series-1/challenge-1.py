from secrets import *
import requests
from stellar_sdk import Server, Keypair, TransactionBuilder, Network

# Set Server Address
server = Server("https://horizon-testnet.stellar.org")

# Get Account Secrets and Addresses
kirito_keypair = Keypair.from_secret(kirito_secret)
destination = Keypair.random()

# Create Accoutn
source_account = server.load_account(account_id=kirito_keypair.public_key)
transaction = TransactionBuilder(
    source_account=source_account,
    network_passphrase=Network.TESTNET_NETWORK_PASSPHRASE,
    base_fee=100) \
    .append_create_account_op(destination="GCEEQPRC5SXWC552SEP3D2XWSYIQWDATY7CAX2T2Y7VNZ2JQSMHLMFH2", starting_balance="1000") \
    .build()
transaction.sign(kirito_keypair)
response = server.submit_transaction(transaction)
print(response)
