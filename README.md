# hckblk
BLK Hackathon 2020 prototype project on Smart Contract

# Useful links
https://docs.aws.amazon.com/blockchain-templates/latest/developerguide/blockchain-templates-hyperledger.html

# Useful commands
### Create 2 new blocks
`peer chaincode invoke -o orderer.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel -n mydapp -c '{"Args":["createCar", "CAR02", "Volkswagen", "Passat", "Black", "Anushka"]}'`

`peer chaincode invoke -o orderer.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel -n mydapp -c '{"Args":["createCar", "CAR03", "Volkswagen", "Polo", "Blue", "Virat"]}'`


### query the blocks
`peer chaincode query -C mychannel -n mydapp -c '{"Args":["queryCar", "CAR02"]}'`

`peer chaincode query -C mychannel -n mydapp -c '{"Args":["queryCar", "CAR03"]}'`


### update the CAR02 block (identified by Car ID)
`peer chaincode invoke -o orderer.example.com:7050 --tls --cafile $ORDERER_CA -C mychannel -n mydapp -c '{"Args":["changeCarOwner", "CAR02", "Jennifer"]}'`
### query the block post update
`peer chaincode query -C mychannel -n mydapp -c '{"Args":["queryCar", "CAR02"]}'`


# DEMO steps
Currently there are 6 blocks and 7 transactions in the blockchain

We will create 2 new entries

Query the chain about both of them. This will not add any block.

Update one of the blocks, which actually creates a new block in the chain and does not actually update the old entry.

Will verify post updating whether the updated data is visible or not