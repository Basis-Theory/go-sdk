# Reference
## Applications
<details><summary><code>client.Applications.List() -> *basistheory.ApplicationPaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.ApplicationsListRequest{
        ID: []*string{
            basistheory.String(
                "id",
            ),
        },
        Type: []*string{
            basistheory.String(
                "type",
            ),
        },
        Page: basistheory.Int(
            1,
        ),
        Start: basistheory.String(
            "start",
        ),
        Size: basistheory.Int(
            1,
        ),
    }
client.Applications.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.Create(request) -> *basistheory.Application</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.CreateApplicationRequest{
        Name: "name",
        Type: "type",
    }
client.Applications.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**rules:** `[]*basistheory.AccessRule` 
    
</dd>
</dl>

<dl>
<dd>

**createKey:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.Get(ID) -> *basistheory.Application</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Applications.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.Update(ID, request) -> *basistheory.Application</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.UpdateApplicationRequest{
        Name: "name",
    }
client.Applications.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**rules:** `[]*basistheory.AccessRule` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Applications.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Applications.GetByKey() -> *basistheory.Application</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Applications.GetByKey(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplicationKeys
<details><summary><code>client.ApplicationKeys.List(ID) -> []*basistheory.ApplicationKey</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.ApplicationKeysListRequest{
        KeyID: []*string{
            basistheory.String(
                "id",
            ),
        },
        Type: []*string{
            basistheory.String(
                "type",
            ),
        },
    }
client.ApplicationKeys.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**keyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplicationKeys.Create(ID) -> *basistheory.ApplicationKey</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplicationKeys.Create(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplicationKeys.Get(ID, KeyID) -> *basistheory.ApplicationKey</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplicationKeys.Get(
        context.TODO(),
        "id",
        "keyId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**keyID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplicationKeys.Delete(ID, KeyID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplicationKeys.Delete(
        context.TODO(),
        "id",
        "keyId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**keyID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplicationTemplates
<details><summary><code>client.ApplicationTemplates.List() -> []*basistheory.ApplicationTemplate</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplicationTemplates.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplicationTemplates.Get(ID) -> *basistheory.ApplicationTemplate</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplicationTemplates.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplePay
<details><summary><code>client.ApplePay.Create(request) -> *basistheory.ApplePayCreateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.ApplePayCreateRequest{}
client.ApplePay.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**expiresAt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**applePaymentData:** `*basistheory.ApplePayMethodToken` 
    
</dd>
</dl>

<dl>
<dd>

**merchantRegistrationID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Get(ID) -> *basistheory.ApplePayToken</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplePay.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Delete(ID) -> string</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplePay.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## GooglePay
<details><summary><code>client.GooglePay.Create(request) -> *basistheory.GooglePayCreateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.GooglePayCreateRequest{}
client.GooglePay.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**expiresAt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**googlePaymentData:** `*basistheory.GooglePayMethodToken` 
    
</dd>
</dl>

<dl>
<dd>

**merchantRegistrationID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GooglePay.Get(ID) -> *basistheory.GooglePayToken</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.GooglePay.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GooglePay.Delete(ID) -> string</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.GooglePay.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Documents
<details><summary><code>client.Documents.Upload(request) -> *basistheory.Document</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.DocumentsUploadRequest{
        Document: strings.NewReader(
            "",
        ),
    }
client.Documents.Upload(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Documents.Get(ID) -> *basistheory.Document</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Documents.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Documents.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Documents.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tokens
<details><summary><code>client.Tokens.Detokenize(request) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Tokens.Detokenize(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Tokenize(request) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Tokens.Tokenize(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Get(ID) -> *basistheory.Token</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tokens.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tokens.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Update(ID, request) -> *basistheory.Token</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.UpdateTokenRequest{}
client.Tokens.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `any` 
    
</dd>
</dl>

<dl>
<dd>

**privacy:** `*basistheory.UpdatePrivacy` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**searchIndexes:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**fingerprintExpression:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**mask:** `any` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**deduplicateToken:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**containers:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.Create(request) -> *basistheory.Token</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.CreateTokenRequest{}
client.Tokens.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*basistheory.CreateTokenRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.ListV2() -> *basistheory.TokenCursorPaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.TokensListV2Request{
        Type: basistheory.String(
            "type",
        ),
        Container: basistheory.String(
            "container",
        ),
        Fingerprint: basistheory.String(
            "fingerprint",
        ),
        Start: basistheory.String(
            "start",
        ),
        Size: basistheory.Int(
            1,
        ),
    }
client.Tokens.ListV2(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**type_:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**container:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**fingerprint:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokens.SearchV2(request) -> *basistheory.TokenCursorPaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.SearchTokensRequestV2{}
client.Tokens.SearchV2(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Enrichments
<details><summary><code>client.Enrichments.BankAccountVerify(request) -> *basistheory.BankVerificationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.BankVerificationRequest{
        TokenID: "token_id",
    }
client.Enrichments.BankAccountVerify(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tokenID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**countryCode:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**routingNumber:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Enrichments.CardDetails() -> *basistheory.CardDetailsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.EnrichmentsCardDetailsRequest{
        Bin: "bin",
    }
client.Enrichments.CardDetails(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bin:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Keys
<details><summary><code>client.Keys.List() -> []*basistheory.ClientEncryptionKeyMetadataResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Create(request) -> *basistheory.ClientEncryptionKeyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.ClientEncryptionKeyRequest{}
client.Keys.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**expiresAt:** `*time.Time` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Get(ID) -> *basistheory.ClientEncryptionKeyMetadataResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Logs
<details><summary><code>client.Logs.List() -> *basistheory.LogPaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.LogsListRequest{
        EntityType: basistheory.String(
            "entity_type",
        ),
        EntityID: basistheory.String(
            "entity_id",
        ),
        StartDate: basistheory.Time(
            basistheory.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        EndDate: basistheory.Time(
            basistheory.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Page: basistheory.Int(
            1,
        ),
        Start: basistheory.String(
            "start",
        ),
        Size: basistheory.Int(
            1,
        ),
    }
client.Logs.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityType:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**entityID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Logs.GetEntityTypes() -> []*basistheory.LogEntityType</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Logs.GetEntityTypes(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## NetworkTokens
<details><summary><code>client.NetworkTokens.Create(request) -> *basistheory.NetworkToken</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.CreateNetworkTokenRequest{}
client.NetworkTokens.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**data:** `*basistheory.Card` 
    
</dd>
</dl>

<dl>
<dd>

**tokenID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tokenIntentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**expirationMonth:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**expirationYear:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cardholderInfo:** `*basistheory.CardholderInfo` 
    
</dd>
</dl>

<dl>
<dd>

**merchantID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkTokens.Cryptogram(ID) -> *basistheory.NetworkTokenCryptogram</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.NetworkTokens.Cryptogram(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkTokens.Get(ID) -> *basistheory.NetworkToken</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.NetworkTokens.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkTokens.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.NetworkTokens.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkTokens.Suspend(ID) -> *basistheory.NetworkToken</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.NetworkTokens.Suspend(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkTokens.Resume(ID) -> *basistheory.NetworkToken</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.NetworkTokens.Resume(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Permissions
<details><summary><code>client.Permissions.List() -> []*basistheory.Permission</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.PermissionsListRequest{
        ApplicationType: basistheory.String(
            "application_type",
        ),
    }
client.Permissions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**applicationType:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Proxies
<details><summary><code>client.Proxies.List() -> *basistheory.ProxyPaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.ProxiesListRequest{
        ID: []*string{
            basistheory.String(
                "id",
            ),
        },
        Name: basistheory.String(
            "name",
        ),
        Page: basistheory.Int(
            1,
        ),
        Start: basistheory.String(
            "start",
        ),
        Size: basistheory.Int(
            1,
        ),
    }
client.Proxies.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxies.Create(request) -> *basistheory.Proxy</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.CreateProxyRequest{
        Name: "name",
        DestinationURL: "destination_url",
    }
client.Proxies.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**destinationURL:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**requestReactorID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**responseReactorID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requestTransform:** `*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**responseTransform:** `*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**requestTransforms:** `[]*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**responseTransforms:** `[]*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**application:** `*basistheory.Application` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**requireAuth:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**disableDetokenization:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxies.Get(ID) -> *basistheory.Proxy</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Proxies.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxies.Update(ID, request) -> *basistheory.Proxy</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.UpdateProxyRequest{
        Name: "name",
        DestinationURL: "destination_url",
    }
client.Proxies.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**destinationURL:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**requestReactorID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**responseReactorID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requestTransform:** `*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**responseTransform:** `*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**requestTransforms:** `[]*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**responseTransforms:** `[]*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**application:** `*basistheory.Application` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**requireAuth:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**disableDetokenization:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxies.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Proxies.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxies.Patch(ID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.PatchProxyRequest{}
client.Proxies.Patch(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**destinationURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requestTransform:** `*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**responseTransform:** `*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**requestTransforms:** `[]*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**responseTransforms:** `[]*basistheory.ProxyTransform` 
    
</dd>
</dl>

<dl>
<dd>

**application:** `*basistheory.Application` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**requireAuth:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**disableDetokenization:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Reactors
<details><summary><code>client.Reactors.List() -> *basistheory.ReactorPaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.ReactorsListRequest{
        ID: []*string{
            basistheory.String(
                "id",
            ),
        },
        Name: basistheory.String(
            "name",
        ),
        Page: basistheory.Int(
            1,
        ),
        Start: basistheory.String(
            "start",
        ),
        Size: basistheory.Int(
            1,
        ),
    }
client.Reactors.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactors.Create(request) -> *basistheory.Reactor</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.CreateReactorRequest{
        Name: "name",
        Code: "code",
    }
client.Reactors.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**code:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**application:** `*basistheory.Application` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**runtime:** `*basistheory.Runtime` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactors.Get(ID) -> *basistheory.Reactor</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Reactors.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactors.Update(ID, request) -> *basistheory.Reactor</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.UpdateReactorRequest{
        Name: "name",
        Code: "code",
    }
client.Reactors.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**application:** `*basistheory.Application` 
    
</dd>
</dl>

<dl>
<dd>

**code:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**runtime:** `*basistheory.Runtime` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactors.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Reactors.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactors.Patch(ID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.PatchReactorRequest{}
client.Reactors.Patch(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**application:** `*basistheory.Application` 
    
</dd>
</dl>

<dl>
<dd>

**code:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**runtime:** `*basistheory.Runtime` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactors.React(ID, request) -> *basistheory.ReactResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Reactors.React(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reactors.ReactAsync(ID, request) -> *basistheory.AsyncReactResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Reactors.ReactAsync(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Roles
<details><summary><code>client.Roles.List() -> []*basistheory.Role</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Roles.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sessions
<details><summary><code>client.Sessions.Create() -> *basistheory.CreateSessionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sessions.Create(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sessions.Authorize(request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.AuthorizeSessionRequest{
        Nonce: "nonce",
    }
client.Sessions.Authorize(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**nonce:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**rules:** `[]*basistheory.AccessRule` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TokenIntents
<details><summary><code>client.TokenIntents.Get(ID) -> *basistheory.TokenIntent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.TokenIntents.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TokenIntents.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.TokenIntents.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TokenIntents.Create(request) -> *basistheory.CreateTokenIntentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.CreateTokenIntentRequest{
        Type: "type",
    }
client.TokenIntents.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `any` 
    
</dd>
</dl>

<dl>
<dd>

**encrypted:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks
<details><summary><code>client.Webhooks.Ping() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Simple endpoint that can be utilized to verify the application is running
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.Ping(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Get(ID) -> *basistheory.Webhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the webhook
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Update(ID, request) -> *basistheory.Webhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a new webhook
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.UpdateWebhookRequest{
        Name: "webhook-update",
        URL: "http://www.example.com",
        Events: []string{
            "token:created",
        },
    }
client.Webhooks.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the webhook
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — The URL to which the webhook will send events
    
</dd>
</dl>

<dl>
<dd>

**notifyEmail:** `*string` — The email address to use for management notification events. Ie: webhook disabled
    
</dd>
</dl>

<dl>
<dd>

**events:** `[]string` — An array of event types that the webhook will listen for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a new webhook
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.List() -> *basistheory.WebhookList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the configured webhooks
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Create(request) -> *basistheory.Webhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new webhook
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.CreateWebhookRequest{
        Name: "webhook-create",
        URL: "http://www.example.com",
        Events: []string{
            "token:created",
        },
    }
client.Webhooks.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the webhook
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — The URL to which the webhook will send events
    
</dd>
</dl>

<dl>
<dd>

**notifyEmail:** `*string` — The email address to use for management notification events. Ie: webhook disabled
    
</dd>
</dl>

<dl>
<dd>

**events:** `[]string` — An array of event types that the webhook will listen for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AccountUpdater Jobs
<details><summary><code>client.AccountUpdater.Jobs.Get(ID) -> *basistheory.AccountUpdaterJob</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the account updater batch job
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AccountUpdater.Jobs.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountUpdater.Jobs.List() -> *basistheory.AccountUpdaterJobList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of account updater batch jobs
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accountupdater.JobsListRequest{
        Size: basistheory.Int(
            1,
        ),
        Start: basistheory.String(
            "start",
        ),
    }
client.AccountUpdater.Jobs.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**size:** `*int` — The maximum number of jobs to return
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` — Cursor for pagination
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AccountUpdater.Jobs.Create(request) -> *basistheory.AccountUpdaterJob</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the created account updater batch job
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accountupdater.CreateAccountUpdaterJobRequest{}
client.AccountUpdater.Jobs.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**deduplicateTokens:** `*bool` — Whether deduplication should be enabled when creating new tokens. Uses the value of the Deduplicate Tokens setting on the tenant if not set.
    
</dd>
</dl>

<dl>
<dd>

**merchantID:** `*string` — Tenant merchant identifier
    
</dd>
</dl>

<dl>
<dd>

**resultVersion:** `*accountupdater.CreateAccountUpdaterJobRequestResultVersion` — Version of the result CSV format. Version '1' returns base columns. Version '1.1' adds new_fingerprint and new_brand columns. Version '1.2' adds the new_last4 column on top of 1.1.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AccountUpdater RealTime
<details><summary><code>client.AccountUpdater.RealTime.Invoke(request) -> *basistheory.AccountUpdaterRealTimeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the update result
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accountupdater.AccountUpdaterRealTimeRequest{
        TokenID: "9a420b15-ddfe-4793-9466-48f53520e47c",
    }
client.AccountUpdater.RealTime.Invoke(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tokenID:** `string` — Card Token identifier
    
</dd>
</dl>

<dl>
<dd>

**expirationYear:** `*int` — The 4-digit expiration year of the account number. Not required if the card token already stores this value.
    
</dd>
</dl>

<dl>
<dd>

**expirationMonth:** `*int` — The 2-digit expiration month of the account number. Not required if the card token already stores this value.
    
</dd>
</dl>

<dl>
<dd>

**deduplicateToken:** `*bool` — Whether deduplication should be enabled when creating the new token. Uses the value of the Deduplicate Tokens setting on the tenant if not set.
    
</dd>
</dl>

<dl>
<dd>

**merchantID:** `*string` — Tenant merchant identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agentic Agents
<details><summary><code>client.Agentic.Agents.Create(request) -> *basistheory.Agent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agentic.CreateAgentRequest{
        Name: "name",
    }
client.Agentic.Agents.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**enrollmentIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**instanceDetails:** `*basistheory.InstanceDetails` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Agents.Get(AgentID) -> *basistheory.Agent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agentic.Agents.Get(
        context.TODO(),
        "agent_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Agents.Delete(AgentID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agentic.Agents.Delete(
        context.TODO(),
        "agent_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Agents.Update(AgentID, request) -> *basistheory.Agent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agentic.UpdateAgentRequest{}
client.Agentic.Agents.Update(
        context.TODO(),
        "agent_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**enrollmentIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**instanceDetails:** `*basistheory.InstanceDetails` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agentic Enrollments
<details><summary><code>client.Agentic.Enrollments.List() -> *basistheory.EnrollmentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all enrollments for the current tenant with cursor-based pagination.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agentic.EnrollmentsListRequest{
        Limit: basistheory.Int(
            1,
        ),
        Cursor: basistheory.String(
            "cursor",
        ),
    }
client.Agentic.Enrollments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Pagination cursor from a previous response
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Enrollments.Create(request) -> *basistheory.Enrollment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enroll a card token with a card network (Visa or Mastercard).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agentic.CreateEnrollmentRequest{
        TokenID: "token_id",
        Consumer: &basistheory.Consumer{
            Email: "email",
        },
    }
client.Agentic.Enrollments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tokenID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**consumer:** `*basistheory.Consumer` 
    
</dd>
</dl>

<dl>
<dd>

**agentID:** `*string` — Single agent ID (mutually exclusive with agent_ids)
    
</dd>
</dl>

<dl>
<dd>

**agentIDs:** `[]string` — Multiple agent IDs (mutually exclusive with agent_id)
    
</dd>
</dl>

<dl>
<dd>

**walletName:** `*string` — Display label shown to the cardholder during Mastercard managed-authentication challenges. Defaults to "Agent Wallet" when not provided.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*agentic.CreateEnrollmentRequestType` 

Enrollment type. `agentic` (default) enrolls the card for agent-driven payments and requires verification.
`autofill` enrolls the card for direct autofill credential retrieval, skips verification, and is currently
available to test tenants only.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Enrollments.Get(EnrollmentID) -> *basistheory.Enrollment</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agentic.Enrollments.Get(
        context.TODO(),
        "enrollment_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Enrollments.Delete(EnrollmentID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Soft-deletes the enrollment by marking its status as deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agentic.Enrollments.Delete(
        context.TODO(),
        "enrollment_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Enrollments.Retry(EnrollmentID) -> *basistheory.Enrollment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retry enrolling a card that previously failed. Only failed enrollments can be retried.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agentic.Enrollments.Retry(
        context.TODO(),
        "enrollment_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agentic Agents Instructions
<details><summary><code>client.Agentic.Agents.Instructions.List(AgentID) -> *basistheory.InstructionList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all purchase instructions for an agent with cursor-based pagination and optional enrollment filter.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agents.InstructionsListRequest{
        EnrollmentID: basistheory.String(
            "enrollment_id",
        ),
        Limit: basistheory.Int(
            1,
        ),
        Cursor: basistheory.String(
            "cursor",
        ),
    }
client.Agentic.Agents.Instructions.List(
        context.TODO(),
        "agent_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**enrollmentID:** `*string` — Filter instructions by enrollment ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Pagination cursor from a previous response
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Agents.Instructions.Create(AgentID, request) -> *basistheory.Instruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new payment instruction for an agent, linked to an enrollment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agents.CreateInstructionRequest{
        EnrollmentID: "enrollment_id",
        Amount: &basistheory.Amount{
            Value: "100.00",
        },
        Description: "description",
        ExpiresAt: basistheory.MustParseDateTime(
            "2024-01-15T09:30:00Z",
        ),
    }
client.Agentic.Agents.Instructions.Create(
        context.TODO(),
        "agent_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**enrollmentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*basistheory.Amount` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**assuranceData:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**recurring:** `*basistheory.Recurring` 
    
</dd>
</dl>

<dl>
<dd>

**instanceDetails:** `*basistheory.InstanceDetails` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Agents.Instructions.Get(AgentID, InstructionID) -> *basistheory.Instruction</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agentic.Agents.Instructions.Get(
        context.TODO(),
        "agent_id",
        "instruction_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**instructionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Agents.Instructions.Delete(AgentID, InstructionID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agentic.Agents.Instructions.Delete(
        context.TODO(),
        "agent_id",
        "instruction_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**instructionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Agents.Instructions.Update(AgentID, InstructionID, request) -> *basistheory.Instruction</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agents.UpdateInstructionRequest{}
client.Agentic.Agents.Instructions.Update(
        context.TODO(),
        "agent_id",
        "instruction_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**instructionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*basistheory.Amount` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*time.Time` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agentic Agents Instructions Credentials
<details><summary><code>client.Agentic.Agents.Instructions.Credentials.Create(AgentID, InstructionID, request) -> *basistheory.Credentials</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve payment credentials (card number, expiration, CVC) for a purchase instruction.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &instructions.GetCredentialsRequest{
        Merchant: &basistheory.AgenticMerchant{
            Name: "name",
            URL: "url",
            CountryCode: "country_code",
        },
    }
client.Agentic.Agents.Instructions.Credentials.Create(
        context.TODO(),
        "agent_id",
        "instruction_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**instructionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**products:** `[]*basistheory.Product` 
    
</dd>
</dl>

<dl>
<dd>

**merchant:** `*basistheory.AgenticMerchant` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*basistheory.Amount` 
    
</dd>
</dl>

<dl>
<dd>

**deliveryMethod:** `*basistheory.DeliveryMethod` 
    
</dd>
</dl>

<dl>
<dd>

**shippingAddress:** `*basistheory.ShippingAddress` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agentic Agents Instructions Verify
<details><summary><code>client.Agentic.Agents.Instructions.Verify.Start(AgentID, InstructionID, request) -> *basistheory.VerificationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Initiate cardholder verification for a purchase instruction that requires it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.StartVerificationRequest{
        DeviceContext: &basistheory.DeviceContext{
            ScreenHeight: 1,
            ScreenWidth: 1,
            UserAgentString: "user_agent_string",
            LanguageCode: "language_code",
            TimeZone: "time_zone",
            JavaScriptEnabled: true,
            ClientDeviceID: "client_device_id",
            ClientReferenceID: "client_reference_id",
            PlatformType: basistheory.DeviceContextPlatformTypeWeb,
        },
    }
client.Agentic.Agents.Instructions.Verify.Start(
        context.TODO(),
        "agent_id",
        "instruction_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**instructionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*basistheory.StartVerificationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Agents.Instructions.Verify.Passkey(AgentID, InstructionID, request) -> *basistheory.Instruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit passkey/FIDO assertion data to complete instruction verification.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &instructions.SubmitPasskeyRequest{
        AssuranceData: map[string]any{
            "key": "value",
        },
    }
client.Agentic.Agents.Instructions.Verify.Passkey(
        context.TODO(),
        "agent_id",
        "instruction_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**agentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**instructionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**assuranceData:** `map[string]any` — Visa format (identifier, dfp_session_id, fido_assertion_data) or Mastercard format (flexible object)
    
</dd>
</dl>

<dl>
<dd>

**srcCorrelationID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**flowID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agentic Enrollments Verify
<details><summary><code>client.Agentic.Enrollments.Verify.Start(EnrollmentID, request) -> *basistheory.VerificationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Initiates the cardholder verification flow for a pending enrollment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.StartVerificationRequest{
        DeviceContext: &basistheory.DeviceContext{
            ScreenHeight: 1,
            ScreenWidth: 1,
            UserAgentString: "user_agent_string",
            LanguageCode: "language_code",
            TimeZone: "time_zone",
            JavaScriptEnabled: true,
            ClientDeviceID: "client_device_id",
            ClientReferenceID: "client_reference_id",
            PlatformType: basistheory.DeviceContextPlatformTypeWeb,
        },
    }
client.Agentic.Enrollments.Verify.Start(
        context.TODO(),
        "enrollment_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*basistheory.StartVerificationRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Enrollments.Verify.Method(EnrollmentID, request) -> *basistheory.VerificationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Choose the OTP delivery method (SMS, email, etc.) for enrollment verification.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &enrollments.SelectMethodRequest{
        MethodID: "method_id",
    }
client.Agentic.Enrollments.Verify.Method(
        context.TODO(),
        "enrollment_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**methodID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Enrollments.Verify.Otp(EnrollmentID, request) -> *basistheory.VerificationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit the one-time password received by the cardholder.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &enrollments.SubmitOtpRequest{
        OtpCode: "otp_code",
    }
client.Agentic.Enrollments.Verify.Otp(
        context.TODO(),
        "enrollment_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**otpCode:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agentic.Enrollments.Verify.Complete(EnrollmentID, request) -> *basistheory.VerificationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Complete the verification flow (e.g. after passkey creation). Body is optional — Visa sends empty body, Mastercard sends assurance_data.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &enrollments.CompleteVerificationRequest{}
client.Agentic.Enrollments.Verify.Complete(
        context.TODO(),
        "enrollment_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**assuranceData:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**srcCorrelationID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplePay Merchant
<details><summary><code>client.ApplePay.Merchant.Get(ID) -> *basistheory.ApplePayMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplePay.Merchant.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Merchant.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplePay.Merchant.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Merchant.Create(request) -> *basistheory.ApplePayMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &applepay.ApplePayMerchantRegisterRequest{}
client.ApplePay.Merchant.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantIdentifier:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplePay Domain
<details><summary><code>client.ApplePay.Domain.Deregister(request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &applepay.ApplePayDomainDeregistrationRequest{
        Domain: "domain",
    }
client.ApplePay.Domain.Deregister(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domain:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Domain.Get() -> *basistheory.ApplePayDomainRegistrationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplePay.Domain.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Domain.Register(request) -> *basistheory.ApplePayDomainRegistrationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &applepay.ApplePayDomainRegistrationRequest{
        Domain: "domain",
    }
client.ApplePay.Domain.Register(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domain:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Domain.RegisterAll(request) -> *basistheory.ApplePayDomainRegistrationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &applepay.ApplePayDomainRegistrationListRequest{}
client.ApplePay.Domain.RegisterAll(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domains:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplePay Session
<details><summary><code>client.ApplePay.Session.Create(request) -> string</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &applepay.ApplePaySessionRequest{}
client.ApplePay.Session.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**validationURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**domain:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantRegistrationID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplePay Merchant Certificates
<details><summary><code>client.ApplePay.Merchant.Certificates.Get(MerchantID, ID) -> *basistheory.ApplePayMerchantCertificates</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplePay.Merchant.Certificates.Get(
        context.TODO(),
        "merchantId",
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Merchant.Certificates.Delete(MerchantID, ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ApplePay.Merchant.Certificates.Delete(
        context.TODO(),
        "merchantId",
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ApplePay.Merchant.Certificates.Create(MerchantID, request) -> *basistheory.ApplePayMerchantCertificates</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchant.ApplePayMerchantCertificatesRegisterRequest{
        PaymentProcessorCertificateData: "payment_processor_certificate_data",
        PaymentProcessorCertificatePassword: "payment_processor_certificate_password",
    }
client.ApplePay.Merchant.Certificates.Create(
        context.TODO(),
        "merchantId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantCertificateData:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantCertificatePassword:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**paymentProcessorCertificateData:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**paymentProcessorCertificatePassword:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**domain:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Documents Data
<details><summary><code>client.Documents.Data.Get(DocumentID) -> string</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Documents.Data.Get(
        context.TODO(),
        "documentId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**documentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## GooglePay Merchant
<details><summary><code>client.GooglePay.Merchant.Get(ID) -> *basistheory.GooglePayMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.GooglePay.Merchant.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GooglePay.Merchant.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.GooglePay.Merchant.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GooglePay.Merchant.Create(request) -> *basistheory.GooglePayMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &googlepay.GooglePayMerchantRegisterRequest{}
client.GooglePay.Merchant.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantIdentifier:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## GooglePay Merchant Certificates
<details><summary><code>client.GooglePay.Merchant.Certificates.Get(MerchantID, ID) -> *basistheory.GooglePayMerchantCertificates</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.GooglePay.Merchant.Certificates.Get(
        context.TODO(),
        "merchantId",
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GooglePay.Merchant.Certificates.Delete(MerchantID, ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.GooglePay.Merchant.Certificates.Delete(
        context.TODO(),
        "merchantId",
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GooglePay.Merchant.Certificates.Create(MerchantID, request) -> *basistheory.GooglePayMerchantCertificates</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &merchant.GooglePayMerchantCertificatesRegisterRequest{}
client.GooglePay.Merchant.Certificates.Create(
        context.TODO(),
        "merchantId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantCertificateData:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantCertificatePassword:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## NetworkTokens Account
<details><summary><code>client.NetworkTokens.Account.Get(ID) -> *basistheory.NetworkTokenAccount</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.NetworkTokens.Account.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Reactors Results
<details><summary><code>client.Reactors.Results.Get(ID, RequestID) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Reactors.Results.Get(
        context.TODO(),
        "id",
        "requestId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants SecurityContact
<details><summary><code>client.Tenants.SecurityContact.Get() -> *basistheory.SecurityContactEmailResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.SecurityContact.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.SecurityContact.Update(request) -> *basistheory.SecurityContactEmailResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.SecurityContactEmailRequest{
        Email: "email",
    }
client.Tenants.SecurityContact.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants Connections
<details><summary><code>client.Tenants.Connections.Create(request) -> *basistheory.CreateTenantConnectionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.CreateTenantConnectionRequest{
        Strategy: "strategy",
        Options: &basistheory.TenantConnectionOptions{},
    }
client.Tenants.Connections.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**strategy:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**options:** `*basistheory.TenantConnectionOptions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Connections.Delete() -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Connections.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants Invitations
<details><summary><code>client.Tenants.Invitations.List() -> *basistheory.TenantInvitationResponsePaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.InvitationsListRequest{
        Status: basistheory.TenantInvitationStatusPending.Ptr(),
        Page: basistheory.Int(
            1,
        ),
        Start: basistheory.String(
            "start",
        ),
        Size: basistheory.Int(
            1,
        ),
    }
client.Tenants.Invitations.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**status:** `*basistheory.TenantInvitationStatus` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Invitations.Create(request) -> *basistheory.TenantInvitationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.CreateTenantInvitationRequest{
        Email: "email",
    }
client.Tenants.Invitations.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Invitations.Resend(InvitationID) -> *basistheory.TenantInvitationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Invitations.Resend(
        context.TODO(),
        "invitationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invitationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Invitations.Get(InvitationID) -> *basistheory.TenantInvitationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Invitations.Get(
        context.TODO(),
        "invitationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invitationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Invitations.Delete(InvitationID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Invitations.Delete(
        context.TODO(),
        "invitationId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**invitationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants Members
<details><summary><code>client.Tenants.Members.List() -> *basistheory.TenantMemberResponsePaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.MembersListRequest{
        UserID: []*string{
            basistheory.String(
                "user_id",
            ),
        },
        Page: basistheory.Int(
            1,
        ),
        Start: basistheory.String(
            "start",
        ),
        Size: basistheory.Int(
            1,
        ),
    }
client.Tenants.Members.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Members.Update(MemberID, request) -> *basistheory.TenantMemberResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.UpdateTenantMemberRequest{
        Role: "role",
    }
client.Tenants.Members.Update(
        context.TODO(),
        "memberId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Members.Delete(MemberID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Members.Delete(
        context.TODO(),
        "memberId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants Merchants
<details><summary><code>client.Tenants.Merchants.List(TenantID) -> *basistheory.TenantMerchantPaginatedList</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.MerchantsListRequest{
        Page: basistheory.Int(
            1,
        ),
        Start: basistheory.String(
            "start",
        ),
        Size: basistheory.Int(
            1,
        ),
    }
client.Tenants.Merchants.List(
        context.TODO(),
        "tenantId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Merchants.Create(TenantID, request) -> *basistheory.TenantMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.TenantMerchantRequest{
        Name: "name",
        Details: &basistheory.MerchantDetails{},
    }
client.Tenants.Merchants.Create(
        context.TODO(),
        "tenantId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*basistheory.TenantMerchantRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Merchants.Get(TenantID, MerchantID) -> *basistheory.TenantMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Merchants.Get(
        context.TODO(),
        "tenantId",
        "merchantId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Merchants.Delete(TenantID, MerchantID) -> *basistheory.TenantMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Merchants.Delete(
        context.TODO(),
        "tenantId",
        "merchantId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Merchants.Update(TenantID, MerchantID, request) -> *basistheory.TenantMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.TenantMerchantRequest{
        Name: "name",
        Details: &basistheory.MerchantDetails{},
    }
client.Tenants.Merchants.Update(
        context.TODO(),
        "tenantId",
        "merchantId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*basistheory.TenantMerchantRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Merchants.RequestOnboarding(TenantID, MerchantID, request) -> *basistheory.TenantMerchant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.ServiceOnboardingRequest{}
client.Tenants.Merchants.RequestOnboarding(
        context.TODO(),
        "tenantId",
        "merchantId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tenantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**merchantID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**accountUpdater:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**networkToken:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**agenticCommerce:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**cardNetworkInfo:** `*basistheory.CardNetworkInfo` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants Owner
<details><summary><code>client.Tenants.Owner.Get() -> *basistheory.TenantMemberResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Owner.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Owner.Transfer(request) -> *basistheory.TenantMemberResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.TransferTenantOwnerRequest{
        MemberID: "member_id",
    }
client.Tenants.Owner.Transfer(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants Self
<details><summary><code>client.Tenants.Self.GetUsageReports() -> *basistheory.TenantUsageReport</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Self.GetUsageReports(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Self.Get() -> *basistheory.Tenant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Self.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Self.Update(request) -> *basistheory.Tenant</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tenants.UpdateTenantRequest{
        Name: "name",
    }
client.Tenants.Self.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**settings:** `map[string]*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Self.Delete() -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.Self.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Threeds Sessions
<details><summary><code>client.Threeds.Sessions.Create(request) -> *basistheory.CreateThreeDsSessionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &threeds.CreateThreeDsSessionRequest{}
client.Threeds.Sessions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pan:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tokenID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tokenIntentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**device:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**webChallengeMode:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**deviceInfo:** `*basistheory.ThreeDsDeviceInfo` 
    
</dd>
</dl>

<dl>
<dd>

**authenticationRequest:** `*basistheory.AuthenticateThreeDsSessionRequest` 
    
</dd>
</dl>

<dl>
<dd>

**callbackURLs:** `*basistheory.ThreeDsCallbackURLs` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Threeds.Sessions.Authenticate(SessionID, request) -> *basistheory.ThreeDsAuthentication</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &basistheory.AuthenticateThreeDsSessionRequest{
        AuthenticationCategory: "authentication_category",
        AuthenticationType: "authentication_type",
    }
client.Threeds.Sessions.Authenticate(
        context.TODO(),
        "sessionId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sessionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*basistheory.AuthenticateThreeDsSessionRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Threeds.Sessions.GetChallengeResult(SessionID) -> *basistheory.ThreeDsAuthentication</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Threeds.Sessions.GetChallengeResult(
        context.TODO(),
        "sessionId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sessionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Threeds.Sessions.Get(ID) -> *basistheory.ThreeDsSession</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Threeds.Sessions.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks Events
<details><summary><code>client.Webhooks.Events.List() -> basistheory.EventTypes</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return a list of available event types
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.Events.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

