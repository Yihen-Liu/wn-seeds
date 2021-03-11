//加载nodejs的模块，模块名叫ethereumjs-wallet
var Wallet = require('ethereumjs-wallet');
//填入自己的密钥
var key = Buffer.from('99a28df28283594a7180261c47e5e65e4bbcecf8b4b5ca4b3e9aae06a772f45e', 'hex');
//console.log(Wallet)
var wallet = Wallet.default.fromPrivateKey(key)
//console.log(wallet)
//填入自己设置的密码
var keystore = wallet.toV3String('12345').then(res=>{
    console.log(res)
});

