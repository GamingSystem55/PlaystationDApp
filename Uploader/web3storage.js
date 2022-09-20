//How to store data using Web3.Storage
//In this how-to guide, you'll learn how to store data programmatically for your development projects using the Web3.Storage client libraries in JavaScript and Go. This includes making your data available on the decentralized IPFS network with persistent long-term storage provided by Filecoin — all for free.

//If you just want to quickly store a few files using Web3.Storage rather than include upload functionality in an app or service you're building, you may want to skip this guide for now and simply use the Files page on the Web3.Storage site.

//For developers, Web3.Storage provides a simple interface for storing data using syntax inspired by familiar web APIs such as fetch and File. This guide focuses on JavaScript client library and Go client library, which are the simplest way to use Web3.Storage programmatically.

//If you're using another language, see the HTTP API reference for details on working with the underlying HTTP API.

//Uploading data to Web3.Storage using a client library requires a free API token, which in turn requires a Web3.Storage account. If you already have an account and a token, read on. If not, have a look at the quickstart guide to get up and running in just a few minutes.

//The Web3.Storage client's put method accepts an array of File objects.

//There are a few different ways of creating File objects available, depending on your platform.


//Installing the client
//JavaScriptGo
//In your JavaScript project, add the web3.storage package to your dependencies:

//npm install web3.storage

import { Web3Storage } from 'web3.storage'

function getAccessToken () 
{
  // If you're just testing, you can paste in a token
  // and uncomment the following line:
  // return 'paste-your-token-here'

  // In a real app, it's better to read an access token from an
  // environement variable or other configuration that's kept outside of
  // your code base. For this to work, you need to set the
  // WEB3STORAGE_TOKEN environment variable before you run your code.
  return process.env.WEB3STORAGE_TOKEN
}

function makeStorageClient ()
{
  return new Web3Storage({ token: getAccessToken() })
}

function getFiles ()
{
  const fileInput = document.querySelector('input[type="file"]')
  return fileInput.files
}

function makeFileObjects () {
  // You can create File objects from a Blob of binary data
  // see: https://developer.mozilla.org/en-US/docs/Web/API/Blob
  // Here we're just storing a JSON object, but you can store images,
  // audio, or whatever you want!
  const obj = { hello: 'world' }
  const blob = new Blob([JSON.stringify(obj)], { type: 'application/json' })

  const files = [
    new File(['contents-of-file-1'], 'plain-utf8.txt'),
    new File([blob], 'hello.json')
  ]
  return files
}
async function storeFiles (files) {
  const client = makeStorageClient()
  const cid = await client.put(files)
  console.log('stored files with cid:', cid)
  return cid
}

async function storeWithProgress (files) {
  // show the root cid as soon as it's ready
  const onRootCidReady = cid => {
    console.log('uploading files with cid:', cid)
  }

  // when each chunk is stored, update the percentage complete and display
  const totalSize = files.map(f => f.size).reduce((a, b) => a + b, 0)
  let uploaded = 0

  const onStoredChunk = size => {
    uploaded += size
    const pct = 100 * (uploaded / totalSize)
    console.log(`Uploading... ${pct.toFixed(2)}% complete`)
  }

  // makeStorageClient returns an authorized Web3.Storage client instance
  const client = makeStorageClient()

  // client.put will invoke our callbacks during the upload
  // and return the root cid when the upload completes
  return client.put(files, { onRootCidReady, onStoredChunk })
}

import { CarReader } from '@ipld/car';

// assume loadCarData returns the contents of a CAR file as a Uint8Array
const carBytes = await loadCarData();
const reader = await CarReader.fromBytes(carBytes);

const client = makeStorageClient();
const cid = await client.putCar(reader);
console.log('Uploaded CAR file to Web3.Storage! CID:', cid);
