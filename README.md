# Go-Projects
I publish various Go projects.


# passwordHash.go
Converts a written text into a hash statement for you.



# average.go 
A simple average calculation project



# blockchain-1.go  & blockchain-2.go  & blockchain-3.go
This is a Go program that implements a simple blockchain using the Bitcoin protocol.  The program defines two structs: BitcoinBlock and Blockchain. A BitcoinBlock represents a block in the blockchain, and contains a previous block hash, a list of transactions, the block data (which is the concatenated transaction list and the previous block hash), and the block hash (which is the SHA-256 hash of the block data). The Blockchain struct contains a slice of BitcoinBlock pointers.  The program defines several methods on the Blockchain struct. The generate_genesis_block method generates the first block in the blockchain (the genesis block). The create_block_from_transaction method creates a new block in the blockchain from a list of transactions. The display_chain method displays the data and hash of each block in the blockchain. The last_block method returns the last block in the blockchain.  The main function creates a new blockchain, adds several blocks to the blockchain from a list of transactions, and displays the blockchain. The transactions are simple strings that represent transactions between different users, with a specified amount of BTC being sent from one user to another. **Overall, this is a basic implementation of a blockchain using Go and the Bitcoin protocol, but it is not a complete or secure implementation. The program does not include features such as mining, proof-of-work, or a consensus mechanism, which are necessary for a fully functional and secure blockchain**. And in blockchain-3, it takes a name and a recipient name from the user and then combines it with the phrase "sends 10 BTC to" to create a list of transactions.
