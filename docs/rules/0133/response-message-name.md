---
rule:
  aep: 133
  name: [core, '0133', response-message-name]
  summary: Create methods must return the resource.
permalink: /133/response-message-name
redirect_from:
  - /0133/response-message-name
---

# Create methods: Resource response message

This rule enforces that all `Create` RPCs have a response message of the
resource, as mandated in [AEP-133][].

## Details

This rule looks at any message matching beginning with `Create`, and complains
if the name of the corresponding output message does not match the name of the
RPC with the prefix `Create` removed.

It also permits a response of `google.longrunning.Operation`; in this case, it
checks the `response_type` in the `google.longrunning.operation_info`
annotation and ensures that _it_ corresponds to the name of the RPC with the
prefix `Create` removed.

## Examples

### Standard

**Incorrect** code for this rule:

```proto
// Incorrect.
// Should be `Book`.
rpc CreateBook(CreateBookRequest) returns (CreateBookResponse) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*}/books"
    body: "book"
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc CreateBook(CreateBookRequest) returns (Book) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*}/books"
    body: "book"
  };
}
```

### Long-running operation

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc CreateBook(CreateBookRequest) returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v1/{book.name=publishers/*}/books"
    body: "book"
  };
  option (google.longrunning.operation_info) = {
    response_type: "CreateBookResponse"  // Should be "Book".
    metadata_type: "CreateBookMetadata"
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc CreateBook(CreateBookRequest) returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v1/{book.name=publishers/*}/books"
    body: "book"
  };
  option (google.longrunning.operation_info) = {
    response_type: "Book"
    metadata_type: "CreateBookMetadata"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0133::response-message-name=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc CreateBook(CreateBookRequest) returns (CreateBookResponse) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*}/books"
    body: "book"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-133]: https://aep.dev/133
[aep.dev/not-precedent]: https://aep.dev/not-precedent
