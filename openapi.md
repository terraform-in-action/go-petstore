---
title: Swagger Petstore
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - javascript--nodejs: Node.JS
  - ruby: Ruby
  - python: Python
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<h1 id="swagger-petstore">Swagger Petstore v1.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

* <a href="https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api">https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api</a>

 License: MIT

<h1 id="swagger-petstore-pets">pets</h1>

## listPets

<a id="opIdlistPets"></a>

> Code samples

```shell
# You can also use wget
curl -X GET https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets \
  -H 'Accept: application/json'

```

```http
GET https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets HTTP/1.1
Host: g6pny6dke9.execute-api.us-west-2.amazonaws.com
Accept: application/json

```

```javascript
var headers = {
  'Accept':'application/json'

};

$.ajax({
  url: 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets',
  method: 'get',

  headers: headers,
  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const fetch = require('node-fetch');

const headers = {
  'Accept':'application/json'

};

fetch('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /pets`

*List all pets*

<h3 id="listpets-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|limit|query|integer(int32)|false|How many items to return at one time (default 100)|

> Example responses

> 200 Response

```json
{
  "items": [
    {
      "id": "string",
      "name": "string",
      "species": "string",
      "age": 0
    }
  ]
}
```

<h3 id="listpets-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|An array of pets|[PetList](#schemapetlist)|

<aside class="success">
This operation does not require authentication
</aside>

## createPet

<a id="opIdcreatePet"></a>

> Code samples

```shell
# You can also use wget
curl -X POST https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets HTTP/1.1
Host: g6pny6dke9.execute-api.us-west-2.amazonaws.com
Content-Type: application/json
Accept: application/json

```

```javascript
var headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

$.ajax({
  url: 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets',
  method: 'post',

  headers: headers,
  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const fetch = require('node-fetch');
const inputBody = '{
  "name": "string",
  "species": "string",
  "age": 0
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

fetch('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /pets`

*Create a Pet*

> Body parameter

```json
{
  "name": "string",
  "species": "string",
  "age": 0
}
```

<h3 id="createpet-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[CreatePetRequest](#schemacreatepetrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "name": "string",
  "species": "string",
  "age": 0
}
```

<h3 id="createpet-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Expected response to a valid request|[Pet](#schemapet)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|None|

<aside class="success">
This operation does not require authentication
</aside>

## showPetById

<a id="opIdshowPetById"></a>

> Code samples

```shell
# You can also use wget
curl -X GET https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId} \
  -H 'Accept: application/json'

```

```http
GET https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId} HTTP/1.1
Host: g6pny6dke9.execute-api.us-west-2.amazonaws.com
Accept: application/json

```

```javascript
var headers = {
  'Accept':'application/json'

};

$.ajax({
  url: 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
  method: 'get',

  headers: headers,
  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const fetch = require('node-fetch');

const headers = {
  'Accept':'application/json'

};

fetch('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /pets/{petId}`

*Info for a specific pet*

<h3 id="showpetbyid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|petId|path|string|true|The id of the pet to retrieve|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "name": "string",
  "species": "string",
  "age": 0
}
```

<h3 id="showpetbyid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Expected response to a valid request|[Pet](#schemapet)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not found|None|

<aside class="success">
This operation does not require authentication
</aside>

## deletePetById

<a id="opIddeletePetById"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}

```

```http
DELETE https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId} HTTP/1.1
Host: g6pny6dke9.execute-api.us-west-2.amazonaws.com

```

```javascript

$.ajax({
  url: 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
  method: 'delete',

  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const fetch = require('node-fetch');

fetch('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
{
  method: 'DELETE'

})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

result = RestClient.delete 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
  params: {
  }

p JSON.parse(result)

```

```python
import requests

r = requests.delete('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}', params={

)

print r.json()

```

```java
URL obj = new URL("https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /pets/{petId}`

*Delete a Pet*

<h3 id="deletepetbyid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|petId|path|string|true|The id of the pet to delete|

<h3 id="deletepetbyid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|ok|None|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not found|None|

<aside class="success">
This operation does not require authentication
</aside>

## patchPetById

<a id="opIdpatchPetById"></a>

> Code samples

```shell
# You can also use wget
curl -X PATCH https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId} \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
PATCH https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId} HTTP/1.1
Host: g6pny6dke9.execute-api.us-west-2.amazonaws.com
Content-Type: application/json
Accept: application/json

```

```javascript
var headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

$.ajax({
  url: 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
  method: 'patch',

  headers: headers,
  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const fetch = require('node-fetch');
const inputBody = '{
  "name": "string",
  "species": "string",
  "age": 0
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

fetch('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
{
  method: 'PATCH',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.patch 'https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.patch('https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PATCH");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PATCH", "https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/v1/api/pets/{petId}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PATCH /pets/{petId}`

*Update a pet*

> Body parameter

```json
{
  "name": "string",
  "species": "string",
  "age": 0
}
```

<h3 id="patchpetbyid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|petId|path|string|true|The id of the pet to update|
|body|body|[UpdatePetRequest](#schemaupdatepetrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "name": "string",
  "species": "string",
  "age": 0
}
```

<h3 id="patchpetbyid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Ok|[Pet](#schemapet)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|None|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not found|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|None|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocSpet">Pet</h2>

<a id="schemapet"></a>

```json
{
  "id": "string",
  "name": "string",
  "species": "string",
  "age": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|true|none|none|
|name|string|true|none|none|
|species|string|true|none|none|
|age|integer|true|none|none|

<h2 id="tocSpetlist">PetList</h2>

<a id="schemapetlist"></a>

```json
{
  "items": [
    {
      "id": "string",
      "name": "string",
      "species": "string",
      "age": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|items|[[Pet](#schemapet)]|true|none|none|

<h2 id="tocScreatepetrequest">CreatePetRequest</h2>

<a id="schemacreatepetrequest"></a>

```json
{
  "name": "string",
  "species": "string",
  "age": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|true|none|none|
|species|string|true|none|none|
|age|integer|true|none|none|

<h2 id="tocSupdatepetrequest">UpdatePetRequest</h2>

<a id="schemaupdatepetrequest"></a>

```json
{
  "name": "string",
  "species": "string",
  "age": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|true|none|none|
|species|string|true|none|none|
|age|integer|true|none|none|

