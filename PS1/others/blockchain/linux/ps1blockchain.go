//How to build a blockchain from scratch with Go
//December 28, 2021  5 min read 


///Blockchains are the underlying technology for many decentralized applications and cryptocurrencies. They are considered one of this generation’s most significant discoveries. Despite being in its early stages, blockchain is already applicable in many industries and is creating new job roles and opportunities for developers, artists, gamers, content writers, and many more.

//This tutorial aims to teach you how blockchains work by guiding you through building one from scratch with Go. If you have heard of blockchains, but are still confused about how they work, this article is for you.

//Why build a blockchain with Go?
///Go provides many unique features and functionalities that make it a good fit for building a blockchain. For example, Go allows you to create highly efficient and performant applications with little effort.

//Go is also excellent for building applications that require parallelism and concurrency (like blockchains) with its ability to spawn and manage thousands of Goroutines. Go implements automatic garbage collection and stack management with its runtime system.

//Finally, it compiles applications to machine code and single binaries, supporting multiple OSs and processor architectures, and deploys easily on server infrastructure.

//Prerequisites
//To follow and understand this tutorial, you will need the following:

//orking knowledge of Go
//Go v1.x installed on your machine
//A Go development environment (e.g., text editor, IDE)
//What is a blockchain?
//A blockchain is a digital record of transactions distributed and shared among the nodes of a computer network. Each transaction in the blockchain is called a block, and links to another with cryptography techniques.

//Blockchains are helpful when you are trying to build a decentralized system that ensures the security and integrity of data while maintaining trust between every system user.

//What is a block?
//We mentioned blocks earlier, and you might be wondering what they are. Put simply, a block is a group of data, and multiple blocks come together to form a blockchain.

//Every block in a blockchain possesses the following properties:

//Data to record on the blockchain, e.g., transaction data
//A block hash, the ID of the block generated using cryptography techniques
//The previous block’s hash, which is the cryptographic hash of the last block in the blockchain. It is recorded in every block to link it to the chain and improve its security
//A timestamp of when the block was created and added to the blockchain
//Proof of Work (PoW), which is the amount of effort taken to derive the current block’s hash. We will explain this in depth later in the tutorial
//Getting started
//Let us start by creating a new Go project and importing all the necessary packages to build our blockchain. Create a file named blockchain.go and save the following code in it:

package main

import (
        "crypto/sha256"
        "encoding/json"
        "fmt"
        "strconv"
        "strings"
        "time"
)
//Next, we will create a custom block type to hold our blockchain’s data. Add the following code to the blockchain.go file:

type Block struct {
        data         map[string]interface{}
        hash         string
        previousHash string
        timestamp    time.Time
        pow          int
}
//Then, we will create a custom Blockchain type that contains our blocks. Add the following code to the blockchain.go file:

type Blockchain struct {
        genesisBlock Block
        chain        []Block
        difficulty   int
}

//The genesisBlock property represents the first block added to the blockchain. In contrast, the difficulty property defines the minimum effort miners have to undertake to mine a block and include it in the blockchain.

//Calculating the hash of a block
//Like we discussed earlier, the hash of a block is its identifier generated using cryptography. We will derive the block hash for our blockchain by hashing the previous block hash, current block data, timestamp, and PoW using the SHA256 algorithm.

//Let’s create a method for our Block type that implements this functionality:

func (b Block) calculateHash() string {
        data, _ := json.Marshal(b.data)
        blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
        blockHash := sha256.Sum256([]byte(blockData))
        return fmt.Sprintf("%x", blockHash)
}
//In the code above, we did the following:

//Converted the block’s data to JSON format
//Concatenated the block’s previous hash, data, timestamp, and proof of work (PoW)
//Hashed the earlier concatenation with the SHA256 algorithm
//Returned the hashing result in base 16, with lowercase letters for A-F
//Mining new blocks
//Mining a new block involves generating a block hash that starts with a desired number of zeros (the desired number is the mining difficulty). This means if the difficulty of the blockchain is three, you have to generate a block hash that starts with "000" e.g., "0009a1bfb506…".

//Because we derive a block’s hash from its content, we need to keep changing the PoW value of the current block until we get a hash that satisfies our mining condition (starting zeros > difficulty).

//To implement this, we will create a mine() method for our Block type that will keep incrementing the PoW value and calculating the block hash until we get a valid hash.

//Add the following code to the blockchain.go file:

func (b *Block) mine(difficulty int) {
        for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
                b.pow++
                b.hash = b.calculateHash()
        }
}

//Creating the genesis block
//Next, we will write a function that creates a genesis block for our blockchain and returns a new instance of the Blockchain type.

//Add the following code to the blockchain.go file:

func CreateBlockchain(difficulty int) Blockchain {
        genesisBlock := Block{
                hash:      "0",
                timestamp: time.Now(),
        }
        return Blockchain{
                genesisBlock,
                []Block{genesisBlock},
                difficulty,
        }
}
//Here, we set the hash of our genesis block to "0". Because it is the first block in the blockchain, there is no value for the previous hash, and the data property is empty.

//Then, we created a new instance of the Blockchain type and stored the genesis block along with the blockchain’s difficulty.

//Adding new blocks to the blockchain
//Now that we have implemented functionalities for our blocks to calculate their hash and mine themselves, let’s see how to include new blocks into a blockchain.

//Add the following code to the blockchain.go file:

func (b *Blockchain) addBlock(from, to string, amount float64) {
        blockData := map[string]interface{}{
                "from":   from,
                "to":     to,
                "amount": amount,
        }
        lastBlock := b.chain[len(b.chain)-1]
        newBlock := Block{
                data:         blockData,
                previousHash: lastBlock.hash,
                timestamp:    time.Now(),
        }
        newBlock.mine(b.difficulty)
        b.chain = append(b.chain, newBlock)
}
//Here, we created an addBlock method to the Blockchain type that does the following:

//Collects the details of a transaction (sender, receiver, and transfer amount)
//Creates a new block with the transaction details
//Mines the new block with the previous block hash, current block data, and generated PoW
//Adds the newly created block to the blockchain
//Checking the validity of the blockchain
//We have successfully created a blockchain that can record transactions, and we need a functionality that checks if the blockchain is valid so we know that no transactions have been tampered with.

//Add the following code to the blockchain.go file:

func (b Blockchain) isValid() bool {
        for i := range b.chain[1:] {
                previousBlock := b.chain[i]
                currentBlock := b.chain[i+1]
                if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
                        return false
                }
        }
        return true
}
//Here, we recalculated the hash of every block, compared them with the stored hash values of the other blocks, and checked if the previous hash value of any other block is equal to the hash value of the block before it. If any of the checks fail, the blockchain has been tampered with.

//Using the blockchain to make transactions
//We now have a fully functional blockchain! Let’s create a main() function to show its usage.

//Add the following code to the blockchain.go file:

func main() {
        // create a new blockchain instance with a mining difficulty of 2
        blockchain := CreateBlockchain(2)

        // record transactions on the blockchain for Alice, Bob, and John
        blockchain.addBlock("Alice", "Bob", 5)
        blockchain.addBlock("John", "Bob", 2)

        // check if the blockchain is valid; expecting true
        fmt.Println(blockchain.isValid())
}
//Conclusion
//In this tutorial, you learned how blockchains work under the hood, including what blocks are and what they contain, and how to calculate a block hash, implement a consensus algorithm for mining blocks, record transactions on the blockchain, and validate the authenticity of a blockchain.

//The source code of the Go blockchain is available as a GitHub Gist. I can’t wait to see the amazing things you build with Go, as it is a good fit for applications where efficiency and performance are a top priority.

//Happy coding!