source myScripts/./export_script.sh
CORE_CHAINCODE_LOGLEVEL=debug CORE_PEER_TLS_ENABLED=false CORE_CHAINCODE_ID_NAME=mycc1:2.2 ./simpleChaincode -peer.address 127.0.0.1:7052