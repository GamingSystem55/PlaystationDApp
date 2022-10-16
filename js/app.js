
const ipfsClient = require('ipfs-http-client');
const express = require('express');
const bodyParser = require('body-parser');
const fileUpload = require('express-fileupload');

const ipfs = new ipfsClient({host:'localhost',port})