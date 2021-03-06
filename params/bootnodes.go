// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://9524b5aaf0871ea17ae2696aabb4c4056665a86c10e76b9fb8a47134615d0b9702307d03cfabfaf630727a16198b3039b49ca4661dbe559afba7e64b8b2becdf@13.209.165.97:30303",
	"enode://8c190dd5d91f5d8bb365941eb539102be524bf15ec3f7bd84e6dba925bcc095977b1f2d761c584e7d9d4c39ed0b6fe987fe5f4ef67849de8a56f7a37a2ede987@52.78.48.210:30303",
	"enode://ba6ccea6a4c8e7ad0f3b6931438daaf9a462c19399084f0bc94ceb72b06f1756f058e243b569800fb096b7a40ce11164a7e3b600d43181b7439a0a41ce425cf6@52.79.230.227:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GangnamBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Gangnam test network.
var GangnamBootnodes = []string{
	"enode://b89f336bfe23e24f13590f4d0e1249281d4ffb1a92a5337151a6592db841980cd6756b29f22fbf1990ac8a09ed4662f46a9745d820c6524b6365099d16772603@5.45.85.1:30303",
	"enode://0a3fe912cfc1de4f23cf7f30f9bda221c8b6ee3214e496bd2895739642d7b0e7b8440896cb789ac239ea67f52de18e1b78940c7808f5ace3f25a6c4edf30405c@46.228.240.23:30303",
	"enode://a45d27f2d627eeb71deec6cfecd8e0d87873e5529d3165f984c420f862ecb1be4f2dd483cb82a74822dd40206e73cb7bbbc27d74c38860105f516a5fa293265c@35.240.35.207:30303",
	"enode://e61e62f45eb1fc99410cd717143321351e8ab8b4046c0b853c3f1c307c15b454e4e39c32bd08335ac48f74970f118a6c9f30d420d936f36bc5a449ebe45076a1@74.108.57.80:30303",
	"enode://d470c8590de9fa487d1a4df777e5640cdb34c11ab41a1f1020a0f49243a101206d94faf04fa6ca28a11c3f4b042348c483f486cd5ae6210d82f5864d78b39435@35.240.35.207:30303",
	"enode://4d7719eaa978e3d1ebc953fcfac83adf18f2642782de383d4a4c61f6f2a4d1492ae68c456d9bc7f41089f68cada1a586534c91941232760ef56b2ce33a3094d3@35.204.138.0:30303",
	"enode://6a6832675ea227d552ccd860d540fbff8b86dacc66add979c15711099f5fa0edb46de5a393b9e4a95121b01ab03c16b71ca42185aeef2dd624a4048015a36a32@13.209.112.103:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
