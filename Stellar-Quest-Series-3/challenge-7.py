from secrets import *
import requests
from stellar_sdk import Server, Keypair, TransactionBuilder, Network

# Change
sq_address  = config_address
sq_secret   = config_secret
print(sq_address)

# Fund Via Friendbot
url = 'https://friendbot.stellar.org'
response = requests.get(url, params={'addr': sq_address})
print(response)

# Get sq_jwt
sq_jwt  = kelly_jwt 
max_len = len(sq_jwt)

# Set Server Address
server = Server("https://horizon-testnet.stellar.org")

def get_index_str(index_int):
    if index_int < 10:
        return ("0" + str(index_int))
    else:
        return str(index_int)

# Create Op
transaction = (
    TransactionBuilder(
        source_account = server.load_account(account_id=sq_address), 
        network_passphrase=Network.TESTNET_NETWORK_PASSPHRASE, 
        base_fee=100000) 
)
begin_parser = 0
end_parser=0
index_int=0

# Add Each Op
while (begin_parser < max_len):

    # Get Key
    end_parser += 62
    data_name = get_index_str(index_int) + sq_jwt[begin_parser:end_parser]

    # Get Value
    begin_parser = end_parser
    end_parser += 64
    data_value = sq_jwt[begin_parser:end_parser]

    # Append Operation
    transaction.append_manage_data_op(data_name, data_value) 

    # Set Up Next Loop
    begin_parser = end_parser
    index_int += 1

envelope = transaction.build()
envelope.sign(sq_secret)
response = server.submit_transaction(envelope)
print("\nTransaction hash: {}".format(response["hash"]))
