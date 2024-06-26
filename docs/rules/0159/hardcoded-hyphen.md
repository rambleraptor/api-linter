---
rule:
  aep: 159
  name: [core, '0159', hardcoded-hyphen]
  summary: Request URIs must not hard-code a `-` segment.
permalink: /159/hardcoded-hyphen
redirect_from:
  - /0159/hardcoded-hyphen
---

# Hardcoded hyphen segments

This rule enforces that URIs do not "hard-code" a segment of `-` in their URIs,
as mandated in [AEP-159][].

## Details

This rule looks at every method and complains if it sees a segment in the URI
that is just the `-` character and nothing else.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc ListBooks(ListBooksRequest) returns (ListBooksResponse) {
  option (google.api.http) = {
    get: "/v1/{parent=publishers/-}/books"  // Should use `*`, not `-`.
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc ListBooks(ListBooksRequest) returns (ListBooksResponse) {
  option (google.api.http) = {
    get: "/v1/{parent=publishers/*}/books"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0159::hardcoded-hyphen=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc ListBooks(ListBooksRequest) returns (ListBooksResponse) {
  option (google.api.http) = {
    get: "/v1/{parent=publishers/-}/books"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-159]: https://aep.dev/159
[aep.dev/not-precedent]: https://aep.dev/not-precedent
