
<h1 id="imger">IMGER v1.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Imger it is an HTTP service for image processing based on filters and profiles. Supported filters: overlay, rotate, blur, contrast, brightness, crop, gamma.

Base URLs:

* <a href="/api/v1">/api/v1</a>

Email: <a href="mailto:sumandas.workplace@gmail.com">Support</a> 
 License: MIT

<h1 id="imger-profiles">profiles</h1>

Profile is a configured set of filters that can be applied when processing images.

## get__profiles

> Code samples

```http
GET /api/v1/profiles HTTP/1.1

Accept: application/json

```

`GET /profiles`

Return a list of profiles

<h3 id="get__profiles-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|limit|query|integer(int32)|false|Number of profiles to return (max 10, default 5)|
|skip|query|integer(int32)|false|Number of profiles to skip|

> Example responses

> 200 Response

```json
[
  {
    "id": "string",
    "created": "2019-08-24T14:15:22Z",
    "updated": "2019-08-24T14:15:22Z",
    "filters": [
      {
        "id": "string",
        "parameters": [
          {
            "property1": {},
            "property2": {}
          }
        ]
      }
    ]
  }
]
```

<h3 id="get__profiles-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|List of profiles|Inline|

<h3 id="get__profiles-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Profile](#schemaprofile)]|false|none|none|
|» id|string|false|none|none|
|» created|string(date-time)|false|none|none|
|» updated|string(date-time)|false|none|none|
|» filters|[[Filter](#schemafilter)]|false|none|none|
|»» id|string|false|none|none|
|»» parameters|[object]|false|none|none|
|»»» **additionalProperties**|object|false|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## post__profiles

> Code samples

```http
POST /api/v1/profiles HTTP/1.1

Content-Type: application/json
Accept: application/json

```

`POST /profiles`

Creates a new profile

> Body parameter

```json
{
  "id": "string",
  "filters": [
    {
      "id": "string",
      "parameters": [
        {
          "property1": {},
          "property2": {}
        }
      ]
    }
  ]
}
```

<h3 id="post__profiles-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[CreateProfile](#schemacreateprofile)|true|Profile properties|

> Example responses

> 201 Response

```json
{
  "id": "string",
  "created": "2019-08-24T14:15:22Z",
  "updated": "2019-08-24T14:15:22Z",
  "filters": [
    {
      "id": "string",
      "parameters": [
        {
          "property1": {},
          "property2": {}
        }
      ]
    }
  ]
}
```

<h3 id="post__profiles-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Returns the created profile|[Profile](#schemaprofile)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Body is not a valid json|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Profile not found|None|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|Validation error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## get__profiles_{id}

> Code samples

```http
GET /api/v1/profiles/{id} HTTP/1.1

Accept: application/json

```

`GET /profiles/{id}`

Returns a profile with the given ID

<h3 id="get__profiles_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Profile ID|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "created": "2019-08-24T14:15:22Z",
  "updated": "2019-08-24T14:15:22Z",
  "filters": [
    {
      "id": "string",
      "parameters": [
        {
          "property1": {},
          "property2": {}
        }
      ]
    }
  ]
}
```

<h3 id="get__profiles_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Return a profile|[Profile](#schemaprofile)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Profile not found|None|

<aside class="success">
This operation does not require authentication
</aside>

## put__profiles_{id}

> Code samples

```http
PUT /api/v1/profiles/{id} HTTP/1.1

Content-Type: application/json
Accept: application/json

```

`PUT /profiles/{id}`

Update profile with the given ID

> Body parameter

```json
{
  "filters": [
    {
      "id": "string",
      "parameters": [
        {
          "property1": {},
          "property2": {}
        }
      ]
    }
  ]
}
```

<h3 id="put__profiles_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Profile ID|
|body|body|[UpdateProfile](#schemaupdateprofile)|true|Profile properties to update|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "created": "2019-08-24T14:15:22Z",
  "updated": "2019-08-24T14:15:22Z",
  "filters": [
    {
      "id": "string",
      "parameters": [
        {
          "property1": {},
          "property2": {}
        }
      ]
    }
  ]
}
```

<h3 id="put__profiles_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Returns the updated profile|[Profile](#schemaprofile)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Body is not a valid json|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Profile not found|None|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|Validation error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## delete__profiles_{id}

> Code samples

```http
DELETE /api/v1/profiles/{id} HTTP/1.1

```

`DELETE /profiles/{id}`

Delete profile with the given ID

<h3 id="delete__profiles_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Profile ID|

<h3 id="delete__profiles_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|Deleted profile successfully|None|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Profile not found|None|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="imger-effects">effects</h1>

Effects are used to transform images.

## get__effects

> Code samples

```http
GET /api/v1/effects HTTP/1.1

Accept: application/json

```

`GET /effects`

Returns all available effects

> Example responses

> 200 Response

```json
[
  {
    "id": "string",
    "description": "string",
    "parameters": {
      "description": "string",
      "required": true,
      "type": "string",
      "example": "string",
      "default": "string",
      "values": "string"
    }
  }
]
```

<h3 id="get__effects-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A list of effects|Inline|

<h3 id="get__effects-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Effect](#schemaeffect)]|false|none|none|
|» id|string|false|none|none|
|» description|string|false|none|none|
|» parameters|object|false|none|none|
|»» description|string|false|none|none|
|»» required|boolean|false|none|none|
|»» type|string|false|none|none|
|»» example|string|false|none|none|
|»» default|string|false|none|none|
|»» values|string|false|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## get__effects_{id}

> Code samples

```http
GET /api/v1/effects/{id} HTTP/1.1

Accept: application/json

```

`GET /effects/{id}`

Returns an effect with the given ID

<h3 id="get__effects_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|Effect ID|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "description": "string",
  "parameters": {
    "description": "string",
    "required": true,
    "type": "string",
    "example": "string",
    "default": "string",
    "values": "string"
  }
}
```

<h3 id="get__effects_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Return an effect|[Effect](#schemaeffect)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Effect not found|None|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="imger-images">images</h1>

Processes images based on filters and profiles.

## get__images

> Code samples

```http
GET /api/v1/images?imgSrc=string HTTP/1.1

Accept: image/png

```

`GET /images`

Process image applying the given filters

<h3 id="get__images-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|imgSrc|query|string|true|Image source url|
|profile|query|string|false|Profile to apply|
|filters|query|string|false|Json with filters|

> Example responses

> 400 Response

<h3 id="get__images-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Filters are not a valid json|[Error](#schemaerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Image not found|None|
|422|[Unprocessable Entity](https://tools.ietf.org/html/rfc2518#section-10.3)|Filters are not valid|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Filter">Filter</h2>
<!-- backwards compatibility -->
<a id="schemafilter"></a>
<a id="schema_Filter"></a>
<a id="tocSfilter"></a>
<a id="tocsfilter"></a>

```json
{
  "id": "string",
  "parameters": [
    {
      "property1": {},
      "property2": {}
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|parameters|[object]|false|none|none|
|» **additionalProperties**|object|false|none|none|

<h2 id="tocS_Profile">Profile</h2>
<!-- backwards compatibility -->
<a id="schemaprofile"></a>
<a id="schema_Profile"></a>
<a id="tocSprofile"></a>
<a id="tocsprofile"></a>

```json
{
  "id": "string",
  "created": "2019-08-24T14:15:22Z",
  "updated": "2019-08-24T14:15:22Z",
  "filters": [
    {
      "id": "string",
      "parameters": [
        {
          "property1": {},
          "property2": {}
        }
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|created|string(date-time)|false|none|none|
|updated|string(date-time)|false|none|none|
|filters|[[Filter](#schemafilter)]|false|none|none|

<h2 id="tocS_CreateProfile">CreateProfile</h2>
<!-- backwards compatibility -->
<a id="schemacreateprofile"></a>
<a id="schema_CreateProfile"></a>
<a id="tocScreateprofile"></a>
<a id="tocscreateprofile"></a>

```json
{
  "id": "string",
  "filters": [
    {
      "id": "string",
      "parameters": [
        {
          "property1": {},
          "property2": {}
        }
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|filters|[[Filter](#schemafilter)]|false|none|none|

<h2 id="tocS_UpdateProfile">UpdateProfile</h2>
<!-- backwards compatibility -->
<a id="schemaupdateprofile"></a>
<a id="schema_UpdateProfile"></a>
<a id="tocSupdateprofile"></a>
<a id="tocsupdateprofile"></a>

```json
{
  "filters": [
    {
      "id": "string",
      "parameters": [
        {
          "property1": {},
          "property2": {}
        }
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|filters|[[Filter](#schemafilter)]|false|none|none|

<h2 id="tocS_Effect">Effect</h2>
<!-- backwards compatibility -->
<a id="schemaeffect"></a>
<a id="schema_Effect"></a>
<a id="tocSeffect"></a>
<a id="tocseffect"></a>

```json
{
  "id": "string",
  "description": "string",
  "parameters": {
    "description": "string",
    "required": true,
    "type": "string",
    "example": "string",
    "default": "string",
    "values": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|description|string|false|none|none|
|parameters|object|false|none|none|
|» description|string|false|none|none|
|» required|boolean|false|none|none|
|» type|string|false|none|none|
|» example|string|false|none|none|
|» default|string|false|none|none|
|» values|string|false|none|none|

<h2 id="tocS_Error">Error</h2>
<!-- backwards compatibility -->
<a id="schemaerror"></a>
<a id="schema_Error"></a>
<a id="tocSerror"></a>
<a id="tocserror"></a>

```json
{
  "error_type": "string",
  "message": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|error_type|string|false|none|none|
|message|string|false|none|none|

