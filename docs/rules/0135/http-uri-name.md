---
rule:
  aep: 135
  name: [core, '0135', http-uri-name]
  summary: Delete methods must map the name field to the URI.
permalink: /135/http-uri-name
redirect_from:
  - /0135/http-uri-name
---

# Delete methods: HTTP URI name field

This rule enforces that all `Delete` RPCs map the `name` field to the HTTP URI,
as mandated in [AEP-135][].

## Details

This rule looks at any message matching beginning with `Delete`, and complains
if the `name` variable is not included in the URI. It _does_ check additional
bindings if they are present.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty) {
  option (google.api.http) = {
    delete: "/v1/publishers/*/books/*"  // The `name` field should be extracted.
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty) {
  option (google.api.http) = {
    delete: "/v1/{name=publishers/*/books/*}"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0135::http-uri-name=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty) {
  option (google.api.http) = {
    delete: "/v1/publishers/*/books/*"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-135]: https://aep.dev/135
[aep.dev/not-precedent]: https://aep.dev/not-precedent
