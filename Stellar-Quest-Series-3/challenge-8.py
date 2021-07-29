from secrets import *
import json
import requests
from stellar_sdk import Server, Keypair, TransactionBuilder, Network, TransactionEnvelope

#----- Step 0: Set Up -----#
sq_address  = config_address
sq_secret   = config_secret

# Fund Via Friendbot
url = 'https://friendbot.stellar.org'
response = requests.get(url, params={'addr': sq_address})
print("Friend Funding Response:")
print(response)

# # Get Server Urls
anchor_url = 'https://testanchor.stellar.org'
network = Network.TESTNET_NETWORK_PASSPHRASE

# # Add Trustline to MULT
tx= (
	TransactionBuilder(
		source_account = server.load_account(account_id=sq_address), 
		network_passphrase=Network.TESTNET_NETWORK_PASSPHRASE, 
		base_fee=100) 
		.append_change_trust_op("MULT","GDLD3SOLYJTBEAK5IU4LDS44UMBND262IXPJB3LDHXOZ3S2QQRD5FSMM") #you need to trust the asset in order to receive it
		.build()
)
tx.sign(sq_secret)
response = server.submit_transaction(tx)
print("\nTransaction hash: {}".format(response["hash"]))

#----- Step 1: Get a JWT Token -----#
# Get Request from the Challenge Transaction
data = requests.get(anchor_url + '/auth?account=' + sq_address)
response = data.json()
tx_xdr = response['transaction']

# Create a Transaction Envelope from the xdr and sign it
tx_env = TransactionEnvelope.from_xdr(tx_xdr, network)
tx_env.sign(sq_secret)

# Convert the signed Transaction Envelope to a new xdr to POST
tx_xdr1 = tx_env.to_xdr()

# Reuse of the response variable to store the new xdr
response["transaction"] = tx_xdr1

# POST request with the signed transaction
token_req = requests.post(anchor_url + '/auth', response)

# From the response we need the token
token_json = token_req.json()
ch8_jwt = token_json["token"]

#----- Step 2: Use a GET Request on the SEP-12 Customer Endpoint to get list of data for fake KYC -----#
headers = {'Authorization': 'Bearer {}'.format(ch8_jwt)}
data = requests.get('https://testanchor.stellar.org/kyc/customer', headers=headers)

# # ----- Step 3: Use a PUT request on the SEP-12 customer endpoint to submit that data ----- #
data = requests.put('https://testanchor.stellar.org/kyc/customer', headers=headers, data=kyc_info)
customer_id = data.json()['id']
data = requests.get('https://testanchor.stellar.org/kyc/customer?id=' + customer_id, headers=headers)

# Step 4: Use a GET request on the sep-06 deposit endpoint to despoit of he MULT to your address
data = requests.get('https://testanchor.stellar.org/sep6/info', headers=headers)
mult_data = {
    'account': '{}'.format(sq_address), 
    'amount' : '100'
}
data = requests.get('https://testanchor.stellar.org/sep6/deposit?asset_code=MULT&type=bank_account', headers=headers, data=mult_data)
print(data.json())


