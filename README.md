# Ad Mediation in Go
## How to install
Before we start clone this repo and in terminal install two necessary packages
* go get github.com/gorilla/mux
* go get github.com/go-playground/validator/v10

## How to use
The site is set up on the *Google's App Engine* ~platform and can be accessed via the link *https://gobe-test.oa.r.appspot.com/*~
not hosted any more but some screen shots are at the end of this page :).

### Navigation
#### Ads
*https://gobe-test.oa.r.appspot.com/ad* <br />
Will retunt an non ordered list of ads in database. <br />
The structure of ad is simple: 
* id
* networkId // which network ad uses ex: AdMob, Facebook
* adType // banner, interstitial, reward
* score // bigger score means bigger priorty
* link // link to the ad
#### Network
*https://gobe-test.oa.r.appspot.com/network* <br />
Will retunt a list of networks in database. <br />
The structure of network: 
* id
* name // ex: AdMob, Facebook
* platform // ex: Android, IOS 
* suppVersions // 8.0.0 - 15.5.0 the versions in between are valid
* countryList // "IT", "RS", "SL", "A", "D" the list of supported for this network. Facebook is not supported in CN.
#### Mobiles
*https://gobe-test.oa.r.appspot.com/mobile* <br />
Will retunt an non ordered list of mobiles in database. <br />
Has info about mobile:
* id
* platform // Android
* osVersion // 8.1
* appName // AppOne
* appVersion // 1.2.1
* countryCode // CN
#### Dashboard
*https://gobe-test.oa.r.appspot.com/dashboard* <br />
Will displays an ordered list of ad networks for each ad type and country.
### Filtering data
#### Get ads for specifict mobile
*https://gobe-test.oa.r.appspot.com/mobile-1* <br />
Instead of 1 you can put any mobile ID.
You will get the list of networks available for that mobile.
#### Get ads for specifict mobile and adtype
*https://gobe-test.oa.r.appspot.com/mobile-1-adtype-3* <br />
Instead of 3 you can put 1,2,3 for banner, intersittial, reward respectively.<br />
You will get the sorted list of networks best ads for that mobile and adtype.<br />
List is sorted by score. If some condition is not filled you will get just sorted list for adtype.<br />
### Adding data
For editing data we will use curl. Install form here *https://curl.se/download.html*<br />
Open cmd and run <br />
`curl https://gobe-test.oa.r.appspot.com/mobile -X POST -d "{ \"platform\":\"Android\", \"osVersion\":\"9.0\"}"`<br />
This will add new Mobile. ID will be autoincremented and if you refresth browser you will see new mobile. But if you run <br />
`curl https://gobe-test.oa.r.appspot.com/mobile -X POST -d "{ \"osVersion\":\"9.0\", \"osVersion\":\"9.0\"}"`<br />
You will get an error in cmd because field platform i *required* for mobile, only on that place.
### Updating data
Open cmd and run <br />
`curl https://gobe-test.oa.r.appspot.com/mobile/1 -X PUT -d "{ \"platform\":\"IOS\", \"osVersion\":\"15.0\"}"`<br />
This will update Mobile with id=1. Refresth browser you will see new mobile. You can change everything. <br />
* /netwokr/id
* /ad/id
* /mobile/1id
## Some real life examples
### Example 1
For mobile with id 1 and banner adtype we will show best ads.
Facebook is not suggested network here because this mobile is form China and Fb is not used in CN. (Network id=3 is Facebook for China for Adnroid)
![img1](https://user-images.githubusercontent.com/39196212/165958495-79b9dfe7-86b8-4eef-bccf-d57eb2fd1546.png)

### Example 2 - best ads
For mobile with id 1 and reward video ad will show a list wiht best scored ads.
![img2](https://user-images.githubusercontent.com/39196212/165959487-051717cd-95c2-4e07-a85e-3b7c88fe2650.png)

### Example 3
For mobile with id 2 only suggested network is facebook. Not AdMob because of version of android.
![img31](https://user-images.githubusercontent.com/39196212/165961002-e5e9419f-2da4-4dad-8450-ea3522349b74.png)
AdMob is not working on Adnoid 9.
![img32](https://user-images.githubusercontent.com/39196212/165961060-9739caac-9abf-4c05-9669-f2abd9805600.png)

### Example 4
When we got ad type but no network fits, we will suggest ad by type.
![img4](https://user-images.githubusercontent.com/39196212/165961622-61384246-164e-4cee-9093-39b86c1ae65c.png)

## Feel free to put your test!
