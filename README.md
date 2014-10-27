VAULT  [![Build Status](https://travis-ci.org/franela/vault.svg?branch=master)](https://travis-ci.org/franela/vault)
=====


#What is Vault?

Vault is a __secure__ cross-platform KV store wrapped around GPG which makes team collaboraion extremely easy. 
You can define as many Vaults as needed and use them rightaway without changing any configuration at all.


#Why Vault?

1. Because security is something really important for us and we take it seriously
2. It really simplifies the friction to work with GPG between multiple teams and recipients
3. We needed something which was easy to use and platform agonsotic
4. It was fun to do it and to learn during the process


#What can I do with Vault?

Using Vault is super easy. Below is the list of all the possible operations:

`vault init` - Creates a new Vault in the directory you're located  
`vault add` - Adds one or more recipients to your Vault. Vault will automatically re-encrypt all your files for the new recipient  
`vault set` - Stores something (text or file) into your Vault and ecrypts it for all of your vault recipients  
`vault get` - Retrieves an encrypted file from your Vault and decrypts it.  
`vault remove` - Removes specified recipients from the Vault. It will automatally keep te integrity of your Vault by upgrading encrypted recipients  


#This sounds cool, where do I get it?


You can find Vault for your favorite distrution below:  






Or you can download the source and compile it yourself here.



#Contributing

We're always working to improve vault, but if you find a bug or you just want to collaborate you can send us your PR (with tests please) and we'll be more than glad to review it.

