---
rule:
  aep: 231
  name: [core, '0231', http-method]
  summary: Batch Get methods must use the GET HTTP verb.
permalink: /231/http-method
redirect_from:
  - /0231/http-method
---

# Get methods: GET HTTP verb

This rule enforces that all `BatchGet` RPCs use the `GET` HTTP verb, as
mandated in [AEP-231][].

## Details

This rule looks at any method beginning with `BatchGet`, and
complains if the HTTP verb is anything other than `GET`. It _does_ check
additional bindings if they are present.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc BatchGetBooks(BatchGetBooksRequest) returns (BatchGetBooksResponse) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*}/books:batchGet"  // Should be `get:`.
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc BatchGetBooks(BatchGetBooksRequest) returns (BatchGetBooksResponse) {
  option (google.api.http) = {
    get: "/v1/{name=publishers/*}/books:batchGet"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0231::http-method=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc BatchGetBooks(BatchGetBooksRequest) returns (BatchGetBooksResponse) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*}/books:batchGet"  // Should be `get:`.
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-231]: https://aep.dev/231
[aep.dev/not-precedent]: https://aep.dev/not-precedent
