make orderer peer configtxgen
source myScripts/./export_script.sh
configtxgen -profile SampleDevModeSolo -channelID syschannel -outputBlock genesisblock -configPath $FABRIC_CFG_PATH -outputBlock "$(pwd)/sampleconfig/genesisblock"
bash myScripts/./orderer_start_script.sh
