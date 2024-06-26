---
rule:
  aep: 135
  name: [core, '0135', request-name-required]
  summary: Delete RPCs must have a `name` field in the request.
permalink: /135/request-name-required
redirect_from:
  - /0135/request-name-required
---

# Delete methods: Name field

This rule enforces that all `Delete` standard methods have a `string name`
field in the request message, as mandated in [AEP-135][].

## Details

This rule looks at any message matching `Delete*Request` and complains if
the `name` field is missing.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message DeleteBookRequest {
  // Field name should be `name`.
  string book = 1;
}
```

**Correct** code for this rule:

```proto
// Correct.
message DeleteBookRequest {
  string name = 1;
}
```

## Disabling

If you need to violate this rule, use a leading comment above the message.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0135::request-name-required=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
message DeleteBookRequest {
  string book = 1;
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-135]: https://aep.dev/135
[aep.dev/not-precedent]: https://aep.dev/not-precedent
