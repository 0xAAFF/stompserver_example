# stompdemo


测试Stomp服务的

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).


### 说明

#### A
/network/stompwork.js:11   

```javascript
    var url = "ws://" + window.location.host + "/stomp"//+ "127.0.0.1:3459/stomp" //
```

when npm run serve: 
```javascript
    url = "ws://" + "127.0.0.1:3459/stomp"
```

when npm run build:
```javascript
    url = "ws://" + window.location.host + "/stomp"
```

#### B
/src/components/HelloWorld.vue:34
```javascript
    if (query != "")

    // http://localhost:8080/
    // http://localhost:8080/i
    // http://localhost:8080/someone=test
    // http://localhost:8080/someone=messagesender
```


so open a broser:

if query not '',the page will send message to stomp server.


1 open tow+ web page,
2 use diff uri,
3 F12 open develop tool - console
