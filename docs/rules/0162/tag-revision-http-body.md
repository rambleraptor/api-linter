---
rule:
  aep: 162
  name: [core, '0162', tag-revision-http-body]
  summary: Tag Revision methods should use `*` as the HTTP body.
permalink: /162/tag-revision-http-body
redirect_from:
  - /0162/tag-revision-http-body
---

# Tag Revision methods: HTTP body

This rule enforces that all Tag Revision RPCs use `*` as the HTTP `body`, as mandated in
[AEP-162][].

## Details

This rule looks at any method matching `Tag*Revision`, and complains
if the HTTP `body` field is anything other than `*`.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc TagBookRevision(TagBookRevisionRequest) returns (Book) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/books/*}:tagRevision"
    // body: "*" should be set.
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc TagBookRevision(TagBookRevisionRequest) returns (Book) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/books/*}:tagRevision"
    body: "*"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0162::tag-revision-http-body=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc TagBookRevision(TagBookRevisionRequest) returns (Book) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/books/*}:tagRevision"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-162]: https://aep.dev/162
[aep.dev/not-precedent]: https://aep.dev/not-precedent
