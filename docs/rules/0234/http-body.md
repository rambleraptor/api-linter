---
rule:
  aep: 234
  name: [core, '0234', http-body]
  summary: Batch Update methods should use `*` as the HTTP body.
permalink: /234/http-body
redirect_from:
  - /0234/http-body
---

# Batch Update methods: HTTP body

This rule enforces that all `BatchUpdate` RPCs use `*` as the HTTP `body`, as
mandated in [AEP-234][].

## Details

This rule looks at any RPC methods beginning with `BatchUpdate`, and
complains if the HTTP `body` field is anything other than `*`.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc BatchUpdateBooks(BatchUpdateBooksRequest) returns (BatchUpdateBooksResponse) {
  option (google.api.http) = {
    post: "/v1/{parent=publishers/*}/books:batchUpdate"
    body: "book"  // The http body should be "*".
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc BatchUpdateBooks(BatchUpdateBooksRequest) returns (BatchUpdateBooksResponse) {
  option (google.api.http) = {
    post: "/v1/{parent=publishers/*}/books:batchUpdate"
    body: "*"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0234::http-body=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc BatchUpdateBooks(BatchUpdateBooksRequest) returns (BatchUpdateBooksResponse) {
  option (google.api.http) = {
    post: "/v1/{parent=publishers/*}/books:batchUpdate"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-234]: https://aep.dev/234
[aep.dev/not-precedent]: https://aep.dev/not-precedent
