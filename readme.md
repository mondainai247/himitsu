# Himitsu
**himitsu** is a blockchain built using Cosmos SDK and Tendermint and 
created with [Ignite CLI](https://ignite.com/cli)


### Install
To install the latest version

```
curl https://get.ignite.com/mondainai247/himitsu | sudo bash
```

Himitsu Network: A Scalable, Privacy-Preserving, and Secure Blockchain 
Network

Whitepaper Abstract:
Himitsu is made up of two key elements:

1)  The Himitsu Blockchain Network, a proof-of-stake blockchain built on 
the Cosmos SDK framework, aiming to provide scalability, privacy, and 
security in decentralized systems. Leveraging the widely adopted Cosmos 
SDK, which is used by major cryptocurrency networks such as Binance Coin, 
Cosmos Atom, and Crypto.com's native token CRO. Additionally, Himitsu Network facilitates the 
transfer of Himitsu vouchers and Content Identification Locators (CIDs) 
such as URLs, Google Drive links, BitTorrent hashes, and IPFS CIDs. 

2) The Himitsu Native App
Develoepd using Python as a crude example of the Himitsu Networks functionality. The App Contains:
- A QR Code generator to add verfied friends(tomodachi's).
- A Block Explorer
- A Wallet
- IPFS / Himitsu Network Messaging Application with Full peer-to-peer encryption. 

How it works?
Alice writes Bob a love letter using the Himitsu Native App. When she presses "send" the message is 
encrypted using a shared passcode only she and Bob have. The encrypted file is uploaded to IPFS(IPFS Daemon required). 
IPFS generates a unique CID Hash which is then written to the Himitsu Blockchain Network. Bob receives a notification 
of a new transaction from Alice. Using the CID and their shared password, Bob downloads and decrypts the love letter 
from Alice. Since Alice and Bob, must be verfied friends before messaging, this emiliminates spam. 

Current Versions(Desktop Only - Mobile Comings soon). 
Windows, Mac, Linux


## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
